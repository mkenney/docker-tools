
##############################################################################
##############################################################################
##
##  Manage recipes
##
##  Recipes are structured, delimited lists of values that define a tool
##
##  0 - [recipe_name]
##  1 - [tool_name]
##  2 - [tool_prefix]
##  3 - [tool_template]
##  4 - [docker_image]
##  5 - [image_tag]
##  6 - [entrypoint]
##  7 - [cmd]
##  8 - [volumes]
##  9 - [docker options]
## 10 - [note]
##
##############################################################################
##############################################################################

declare __recipe_delimiter__="|"

#
# Validate a recipe
#
# recipe name|tool name|tool prefix|tool template|docker image|image tag|entrypoint|cmd|volumes|docker options|note
#
# @param a tab-delimited recipe string
# @return 0 or 1
#
function __recipe_validate {
    if [ "" == "$1" ]; then
        echo "A recipe is required"
    fi

    local ret_val=1
    local -a recipe=($(echo "$1" | tr "$__recipe_delimiter__" "\n"))
    local length=${#recipe[@]}

    local name=${recipe[0]}
    local tool=${recipe[1]}
    local prefix=${recipe[2]}
    local template=${recipe[3]}
    local image=${recipe[4]}
    local tag=${recipe[5]}
    local entrypoint=${recipe[6]}
    local cmd=${recipe[7]}
    local volumes=${recipe[8]}
    local docker_opts=${recipe[9]}
    local note=${recipe[10]}

    if [ ""  = "$name" ];  then ret_val="Invalid recipe name"; fi
    if [ ""  = "$tool" ];  then ret_val="Invalid recipe tool"; fi
    if [ ""  = "$image" ]; then ret_val="Invalid recipe image"; fi
    if [ 10 -ne $length ]; then ret_val="Invalid recipe length. $length of 10 values provided"; fi

    printf "$ret_val"
}

#
# Delete a user recipe by name
#
# recipe name|tool name|tool prefix|tool template|docker image|image tag|entrypoint|cmd|volumes|docker options|note
#
# @param $recipe_name Required
#
function __recipe_delete {
    if [ "" == "$1" ]; then
        echo "A recipe name must be provided"
        exit 1
    fi

    local recipe_name=$1
    local recipe_deleted=0
    local -a parts
    local a=0

    while read line; do
        if [ "" != "${line:0:1}" ] && [ "#" != "${line:0:1}" ]; then
            parts=($(echo "$line" | tr "$__recipe_delimiter__" "\n"))
            if [ "$recipe_name" == "${parts[0]}" ]; then
                line=
                recipe_deleted=1
            fi
        fi
        if [ "" != "$line" ]; then
            ret_val[$a]=$line
            a=$((a + 1))
        fi
    done < $DOCKER_TOOLS_RECIPES

    if [ 0 -eq $recipe_deleted ]; then
        echo "Recipe '$recipe_name' not found"
        exit 2
    fi
}

#
# Get a saved recipe by name
# If source is spcified, only the specified recipe source is searched,
# otherwise the user recipes are searched first, followed by the registered
# recipes.
#
# recipe name|tool name|tool prefix|tool template|docker image|image tag|entrypoint|cmd|volumes|docker options|note
#
# @param $recipe_name Required
# @option --source=[recipes|registry] Optional
# @return A stored recipe or ""
#
function __recipe_get {
    if [ "" == "$1" ]; then
        echo "A recipe name must be provided"
        exit 1
    fi

    local recipe_name=$1
    local recipe_sources="$(__get_opt source $@)"
    local recipe_source
    local ret_val
    local ret_source

    IFS=$"$__recipe_delimiter__"
    local -a parts

    case $recipe_sources in
        recipes)
            recipe_sources=$(printf "$DOCKER_TOOLS_RECIPES")
            ;;
        registry)
            recipe_sources=$(printf "$DOCKER_TOOLS_REGISTRY")
            ;;
    esac
    if [ "" == "$recipe_sources" ]; then recipe_sources=$(printf "${DOCKER_TOOLS_RECIPES}${__recipe_delimiter__}${DOCKER_TOOLS_REGISTRY}"); fi

    for recipe_source in $recipe_sources; do
        while read line; do
            if [ "" != "${line:0:1}" ] && [ "#" != "${line:0:1}" ]; then
                parts=($line)
                if [ "$recipe_name" == "${parts[0]}" ]; then
                    ret_val=$line
                    ret_source=$(basename $recipe_source)
                    break
                fi
            fi
        done < $recipe_source
    done

    if [ "" != "$ret_val" ]; then
        printf "${ret_val}${__recipe_delimiter__}${ret_source}"
    fi
}

#
# Save a recipe
# Add a new or update an existing user recipe
#
# recipe name|tool name|tool prefix|tool template|docker image|image tag|entrypoint|cmd|volumes|docker options|note
#
# @param Required, recipe string
#
function __recipe_save {
    if [ "" == "$1" ]; then
        echo "A recipe must be provided"
        exit 1
    fi

    local is_valid="$(__recipe_validate $1)"
    if [ 1 -ne $is_valid ]; then
        echo "Invalid recipe '$1'"
        exit 2
    fi

    local -a ret_val
    local recipe=$1
    local -a recipe_parts=($(echo "$recipe" | tr "$__recipe_delimiter__" "\n"))
    local recipe_saved=0
    local -a parts


    local a=0
    while read line; do
        if [ "" != "${line:0:1}" ] && [ "#" != "${line:0:1}" ]; then
            parts=($(echo "$line" | tr "$__recipe_delimiter__" "\n"))
            if [ "${recipe_parts[0]}" == "${parts[0]}" ]; then
                line=$recipe
                recipe_saved=1
            fi
            name="${parts[0]}"

        fi
        ret_val[$a]=$line
        a=$((a + 1))
    done < $DOCKER_TOOLS_RECIPES
    if [ 0 -eq $recipe_saved ]; then
        ret_val[$a]=$recipe
        a=$((a + 1))
    fi

    # where's the save part... ?
}

#
# Expects `install` command arguments
#
# recipe name|tool name|tool prefix|tool template|docker image|image tag|entrypoint|cmd|volumes|docker options|note
#
# @params `docker-tools install` command arguments
#
function __args_to_recipe {

    # script argument
    local script=$(__get_arg 1)
    if [ "" == "$script" ]; then
        echo "Invalid script name '$script'"
        exit 1
    fi

    # all other options
    local name="$(__get_arg 1 $@)"
    local status="$(__get_opt status $@)"
    local metadata="$(__get_opt metadata $@)"
    local image="$(__get_opt image $@)"
    local tag="$(__get_opt tag $@)"
    local prefix="$(__get_opt prefix $@)"
    local template="$(__get_opt template $@)"
    local entrypoint="$(__get_opt entrypoint $@)"
    local cmd="$(__get_opt cmd $@)"
    local volumes="$(__get_opt volumes $@)"
    local docker_opts="$(__get_opt docker-opts $@)"
    local note="$(__get_opt note $@)"

    name=${name/$__recipe_delimiter__/_}
    status=${status/$__recipe_delimiter__/_}
    metadata=${metadata/$__recipe_delimiter__/_}
    image=${image/$__recipe_delimiter__/_}
    tag=${tag/$__recipe_delimiter__/_}
    prefix=${prefix/$__recipe_delimiter__/_}
    template=${template/$__recipe_delimiter__/_}
    entrypoint=${entrypoint/$__recipe_delimiter__/_}
    cmd=${cmd/$__recipe_delimiter__/_}
    volumes=${volumes/$__recipe_delimiter__/_}
    docker_opts=${docker_opts/$__recipe_delimiter__/_}
    note=${note/$__recipe_delimiter__/_}

    local recipe="${name}${__recipe_delimiter__}${status}${__recipe_delimiter__}${metadata}${__recipe_delimiter__}${script}${__recipe_delimiter__}${image}${__recipe_delimiter__}${tag}${__recipe_delimiter__}${prefix}${__recipe_delimiter__}${template}${__recipe_delimiter__}${entrypoint}${__recipe_delimiter__}${cmd}${__recipe_delimiter__}${volumes}${__recipe_delimiter__}${docker_opts}${__recipe_delimiter__}${note}"
    local is_valid="$(__recipe_validate $recipe)"
    if [ 1 -ne $is_valid ]; then
        echo "Invalid recipe '$recipe'"
        exit 2
    fi

    # Recipe string
    printf "$recipe"
}


#
# Convert a recipe to a list of `docker-tools install` compatible arguments
#
# @param Required, recipe string
# @return space-delmited list of `docker-tools install` arguments
#
function __recipe_to_args {

    if [ "" == "$1" ]; then
        echo "A recipe must be provided"
        exit 1
    fi
    local -a recipe=($(echo "$1" | tr "$__recipe_delimiter__" "\n"))

    local name=${recipe[0]}
    local tool=${recipe[1]}
    local prefix=${recipe[2]}
    local template=${recipe[3]}
    local image=${recipe[4]}
    local tag=${recipe[5]}
    local entrypoint=${recipe[6]}
    local cmd=${recipe[7]}
    local volumes=${recipe[7]}
    local docker_opts=${recipe[9]}
    local note=${recipe[10]}


    if [ "" == "$script" ] || [[ $script == *" "* ]]; then
        echo "Invalid script name '$script'"
        exit 1
    fi
    if [ "" == "$image" ] || [[ $image == *" "* ]]; then
        echo "Invalid image name '$image'"
        exit 1
    fi


    local options
    if [ "" != "$name" ];        then options="$options --name='${name/\'/\\\'}'";               fi
    if [ "" != "$tool" ];        then options="$options --tool='${tool/\'/\\\'}'";               fi
    if [ "" != "$prefix" ];      then options="$options --prefix='${prefix/\'/\\\'}'";           fi
    if [ "" != "$template" ];    then options="$options --template='${template/\'/\\\'}'";       fi
    if [ "" != "$image" ];       then options="$options --image='${image/\'/\\\'}'";             fi
    if [ "" != "$tag" ];         then options="$options --tag='${tag/\'/\\\'}'";                 fi
    if [ "" != "$entrypoint" ];  then options="$options --entrypoint='${entrypoint/\'/\\\'}'";   fi
    if [ "" != "$cmd" ];         then options="$options --cmd='${cmd/\'/\\\'}'";                 fi
    if [ "" != "$volumes" ];     then options="$options --volumes='${volumes/\'/\\\'}'";         fi
    if [ "" != "$docker_opts" ]; then options="$options --docker_opts='${docker_opts/\'/\\\'}'"; fi
    if [ "" != "$note" ];        then options="$options --note='${note/\'/\\\'}'";               fi

    printf "${script}${options}"
}

#
# Generate recipe documentation
#
# @param Required, recipe name
# @option source Optional, a recipe source file in $DOCKER_TOOLS_CONFIG_DIR
# @return Human-readable construct describing the recipe
#
function __recipe_describe {
    IFS=$"$__recipe_delimiter__"

    local ret_val

    if [ "" == "$1" ]; then
        echo "A recipe must be specified"
        exit 1
    fi

    local recipe_name="$1"

    #
    # Load the recipe data
    #
    local -a recipe_parts=($(__recipe_get $@))
    local tool_name=${recipe_parts[1]}
    local tool_prefix=${recipe_parts[2]}
    local tool_template=${recipe_parts[3]}
    local docker_image=${recipe_parts[4]}
    local image_tag=${recipe_parts[5]}
    local entrypoint=${recipe_parts[6]}
    local cmd=${recipe_parts[7]}
    local volumes=${recipe_parts[8]}
    local recipe_source=${recipe_parts[9]}
    local recipe_note=${recipe_parts[10]}

    if [ "" == "$tool_prefix" ]; then tool_prefix=$DOCKER_TOOLS_PREFIX; fi
    if [ "" == "$image_tag" ];   then image_tag="latest";               fi


    # expand the tool path
    local tool_path="$(eval "echo $tool_prefix/$tool_name")"

    # installation status
    local tool_installed=0
    local tool_managed=0
    local tool_updated=0

    if [ -f "$tool_path" ]; then
        tool_installed=1
        if grep -q '__TOOLS_VERSION__=' "$tool_path" && grep -q "__RECIPE_NAME__=${recipe_name/\"/}" "$tool_path"; then
            tool_managed=1
            if grep -q "^__TOOLS_VERSION__=${DOCKER_TOOLS_VERSION/\"/}$" "$tool_path"; then
                tool_updated=1
            fi
        fi
    fi

    # Template vars
    local color_installed=$(_s green bt)
    local color_outdated=$(_s yellow bt)
    local color_unmanaged=$(_s brown)
    local color_not_installed=$(_s b)

    local _status_installed_="$(_s green bt)installed$(_s r)"
    local _status_outofdate_="$(_s yellow bt)outdated$(_s r)"
    local _status_unmanaged_="$(_s brown)unmanaged$(_s r)"
    local _status_not_installed_="$(_s b)not installed$(_s r)"

    local _icon_installed_="→"
    local _icon_outofdate_="→"
    local _icon_unmanaged_="⤳"
    local _icon_not_installed_="─"

    local _recipe_name_="$(_s u)${recipe_name}$(_s r)"
    local _tool_path_="$(_s cyan lt)${tool_path}$(_s r)"
    local _image_=$docker_image:$image_tag
    local _managed_status_="‣"
    local _recipe_status_=$_status_not_installed_

    local _recipe_status_icon_="‣"
    if [ "registry" == "$recipe_source" ]; then
        _recipe_status_icon_="*"
    fi


#✔ ☑ ☆ ★ ✧ ¤ * ｡ﾟ. ☆  ☺  ☻ ☸    ‣




    if [ 1 -eq $tool_installed ]; then
        _recipe_status_=$_status_unmanaged_
        _recipe_status_icon_="$(_s red bt)${_recipe_status_icon_}$(_s r)"
        _tool_path_="$(_s grey)${tool_path}$(_s r)"

        if [ 1 -eq $tool_managed ]; then
            _recipe_status_=$_status_installed_
            _recipe_status_icon_="$(_s green bt)${_recipe_status_icon_}$(_s r)"
            _tool_path_="$(_s white bt bold)${tool_path}$(_s r)"

            if [ 1 -ne $tool_updated ]; then
                _recipe_status_=$_status_outofdate_
                _recipe_status_icon_="$(_s yellow bt)${_recipe_status_icon_}$(_s r)"
                _tool_path_="$(_s i)${tool_path}$(_s r)"
            fi
        fi
    fi


    ret_val="

${_recipe_status_icon_} ${_recipe_name_} -- ${_tool_path_}
    ${_recipe_status_}

    image      - ${_image_}"

    ret_val="$ret_val\n"
    if [ "" != "$entrypoint" ]; then
        ret_val="$ret_val    entrypoint - $entrypoint\n"
    fi
    if [ "" != "$cmd" ]; then
        ret_val="$ret_val    command    - $cmd\n"
    fi
    if [ "" != "$volumes" ]; then
        ret_val="$ret_val    volumes    - ${volumes/;/\n                 - }\n"
    fi

#hash foo 2>/dev/null

    printf "$ret_val"
}


