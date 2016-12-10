
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
## 11 - [recipe_source] - populated by __recipe_get
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
# @param recipe_source
#
declare _recipe_source_status_style_=$(_s indigo)
declare _recipe_source_status_="●"
function __recipe_get_source_status {
    local recipe_source=$1
    if [ "recipes" == "$recipe_source" ]; then
        _recipe_source_status_style_=$(_s blue bt)
    fi

    printf "${_recipe_source_status_style_}|${_recipe_source_status_}"
}

#
# @param recipe_name
# @param tool_path
#
declare _rstatus_na_no_conflict_style_=$(_s grey b)
declare _rstatus_na_no_conflict_status_="not installed"
declare _rstatus_na_unmanaged_style_=$(_s orange b)
declare _rstatus_na_unmanaged_status_="unmanaged file installed"
declare _rstatus_na_managed_style_=$(_s orange b)
declare _rstatus_na_managed_status_="$installed_recipe_name installed"
declare _rstatus_outdated_style_=$(_s yellow bt b)
declare _rstatus_outdated_status_="outdated"
declare _rstatus_updated_style_=$(_s green bt b)
declare _rstatus_updated_status_="installed"
declare _recipe_status_style_
declare _recipe_status_
function __recipe_get_status {
    local recipe_name=$1
    local tool_path=$2

    _recipe_status_style_=$_rstatus_na_no_conflict_style_
    _recipe_status_=$_rstatus_na_no_conflict_status_

    if [ -f "$tool_path" ]; then
        _recipe_status_style_=$_rstatus_na_unmanaged_style_
        _recipe_status_=$_rstatus_na_unmanaged_status_
        if grep -q '__TOOLS_VERSION__=' "$tool_path"; then
            _recipe_status_style_=$_rstatus_na_managed_style_
            _recipe_status_=$_rstatus_na_managed_status_
            if grep -q "__RECIPE_NAME__=${recipe_name/\"/}" "$tool_path"; then
                _recipe_status_style_=$_rstatus_outdated_style_
                _recipe_status_=$_rstatus_outdated_status_
                if grep -q "^__TOOLS_VERSION__=${DOCKER_TOOLS_VERSION/\"/}$" "$tool_path"; then
                    _recipe_status_style_=$_rstatus_updated_style_
                    _recipe_status_=$_rstatus_updated_status_
                fi
            fi
        fi
    fi

    printf "${_recipe_status_style_}|${_recipe_status_}"
}

#
# @param recipe name
# @param tool path
# @param default style
# @param default status
#

declare _rtstatus_na_unmanaged_style_=$(_s red)
declare _rtstatus_na_unmanaged_status_="unmanaged"

declare _rtstatus_na_managed_style_=$(_s green dk)

declare _rtstatus_na_1off_style_=$(_s orange)
declare _rtstatus_na_1off_status_="1-off"

declare _recipe_tool_status_style_
declare _recipe_tool_status_

function __recipe_get_tool_status {

    local recipe_name=$1
    local tool_path=$2
    local default_style=$3
    local default_status=$4

    # this recipe
    _recipe_tool_status_style_=$default_style
    _recipe_tool_status_=$default_status

    # installed, may not be managed
    if [ -f "$tool_path" ]; then
        _recipe_tool_status_style_=${_rtstatus_na_unmanaged_style_}
        _recipe_tool_status_=${_rtstatus_na_unmanaged_status_}

        # installed by docker-tools
        if grep -q '__TOOLS_VERSION__=' "$tool_path"; then
            # not this recipe
            if ! grep -q "__RECIPE_NAME__=${recipe_name/\"/}" "$tool_path"; then

                # installed / other recipe
                _recipe_tool_status_style_=${_rtstatus_na_managed_style_}
                _recipe_tool_status_="$(sed -n -e 's/^declare __RECIPE_NAME__=//p' "$tool_path")"

                if [ "" == "$_recipe_tool_status_" ]; then
                    # installed / no recipe
                    _recipe_tool_status_style_=${_rtstatus_na_1off_style_}
                    _recipe_tool_status_=${_rtstatus_na_1off_status_}
                fi
            fi
        fi
    fi

    printf "${_recipe_tool_status_style_}|${_recipe_tool_status_}"
}


#
# @param recipe name
# @param tool path
# @param default style
# @param default status
#
declare _recipe_path_status_style_
declare _recipe_path_status_
function __recipe_get_path_status {

    local recipe_name=$1
    local tool_path=$2
    local path_path=`which $(basename $tool_path)`
    local default_style=$3
    local default_status=$4

    # not installed
    _recipe_path_status_style_=$(_s grey b)
    _recipe_path_status_="not found"

    if [ -f "$path_path" ]; then
        # installed / unmanaged
        _recipe_path_status_style_=$(_s red bt b)
        _recipe_path_status_="unmanaged"

        if grep -q '__TOOLS_VERSION__=' "$path_path"; then
            # installed by docker-tools
            if grep -q "__RECIPE_NAME__=${recipe_name/\"/}" "$path_path"; then
                # installed / this recipe
                _recipe_path_status_style_=$default_style
                _recipe_path_status_=$default_status

            else
                # installed / other recipe
                _recipe_path_status_style_=$(_s green dk b)
                _recipe_path_status_="$(sed -n -e 's/^declare __RECIPE_NAME__=//p' "$path_path")"

                if [ "" == "$_recipe_path_status_" ]; then
                    # installed / no recipe
                    _recipe_path_status_style_=$(_s orange)
                    _recipe_path_status_="1-off"
                fi
            fi
        fi
    fi

    printf "${_recipe_path_status_style_}|${_recipe_path_status_}"
}

#
# Generate recipe documentation
#
# TODO The styling is horribly slow, replace with a docker service or something
#
# @param Required, recipe name
# @option source Optional, a recipe source file in $DOCKER_TOOLS_CONFIG_DIR
# @return Human-readable construct describing the recipe
#
#
# Notes
#   $recipe_source
#   Sourcefile of this recipe
#       - registry                   :                                                  :
#       - recipes                    :                                                  :
#
#   $recipe_status
#   Status of this recipe
#       - not installed, no conflict        : local _recipe_status_style_=$(_s grey b)         ; local _recipe_status_="not installed";
#       - not installed, conflict           : local _recipe_status_style_=$(_s orange b)       ; local _recipe_status_="not installed";
#       - not installed, conflict recipe    : local _recipe_status_style_=$(_s orange b)       ; local _recipe_status_="not installed";
#       - not installed, conflict unmanaged : local _recipe_status_style_=$(_s orange b)       ; local _recipe_status_="not installed";
#       - installed, out-of-date            : local _recipe_status_style_=$(_s yellow bt b)    ; local _recipe_status_="outdated";
#       - installed, up-to-date             : local _recipe_status_style_=$(_s green bt b)     ; local _recipe_status_="installed";
#
#   $tool_status
#   Status of the file at this recipes tool location
#       - 1-off unsaved tool         : local _tool_status_style_=$(_s orange)           ; local _tool_status="1-off"
#       - other recipe installed     : local _tool_status_style_=$(_s green dk)         ; local _tool_status="recipe: $tool_recipe_name"
#       - installed                  : local _tool_status_style_=$_recipe_status_style_ ; local _tool_status_=$_recipe_status_
#       - not a docker-tools tool    : local _tool_status_style_=$(_s red)              ; local _tool_status="unmanaged"
#       - not installed at all       : local _tool_status_style_=$_recipe_status_style_ ; local _tool_status=$_recipe_status_
#
#   $path_status
#   Status of the tool in your $PATH
#       - no tool in path            : local _path_status_style_=$(_s grey b)           ; local _path_status_="not found"
#       - not a docker-tools tool    : local _path_status_style_=$(_s red bt b)         ; local _path_status_="unmanaged"
#       - this recipe                : local _path_status_style_=$_recipe_status_style_ ; local _path_status_=$_recipe_status_
#       - other recipe installed     : local _path_status_style_=$(_s green dk b)       ; local _path_status_="recipe: $path_recipe_name"
#       - 1-off unsaved tool         : local _path_status_style_=$(_s orange)           ; local _path_status_="1-off"
#
#
function __recipe_describe {
    IFS=$"$__recipe_delimiter__"

    local ret_val

    if [ "" == "$1" ]; then
        echo "A recipe must be specified"
        exit 1
    fi

    # recipe & tool data
    local recipe=$1 #$(__recipe_get $@)
    local -a recipe_parts=($recipe)
    local recipe_name=${recipe_parts[0]}

    local tool_name=${recipe_parts[1]}
    local tool_prefix=${recipe_parts[2]}
    local tool_template=${recipe_parts[3]}
    local docker_image=${recipe_parts[4]}
    local image_tag=${recipe_parts[5]}
    local entrypoint=${recipe_parts[6]}
    local cmd=${recipe_parts[7]}
    local volumes=${recipe_parts[8]}
    local docker_options=${recipe_parts[9]}
    local recipe_note=${recipe_parts[10]}
    local recipe_source=${recipe_parts[11]}

    # defaults
    if [ "" == "$tool_prefix" ]; then tool_prefix=$DOCKER_TOOLS_PREFIX; fi
    if [ "" == "$image_tag" ];   then image_tag="latest";               fi
    if [ "" == "$cmd" ];         then cmd="n/a";                        fi
    if [ "" == "$entrypoint" ];  then entrypoint="n/a";                 fi

    # full tool path
    local tool_path="$(eval "echo $tool_prefix/$tool_name")"

    # recipe_status[0] - style
    # recipe_status[1] - text
    local -a recipe_status=($(__recipe_get_status $recipe_name $tool_path))


    # source_status[0] - style
    # source_status[1] - text
    local -a source_status=($(__recipe_get_source_status $recipe_source))

    # tool_status[0] - style
    # tool_status[1] - text
    local -a tool_status=($(__recipe_get_tool_status $recipe_name $tool_path ${recipe_status[0]} ${recipe_status[1]}))

    # path_status[0] - style
    # path_status[1] - text
    local -a path_status=($(__recipe_get_path_status $recipe_name $tool_path ${recipe_status[0]} ${recipe_status[1]}))

    #
    # Template vars
    #
    # ✔ ☑ ☆ ★ ✧ ¤ * ｡ﾟ. ☆  ☺  ☻ ☸ ‣
    # ⦿ ⁌ ●
    # 😀 😊 😐 😶 😮 😲 😖 😡 😷 💀
    #
    local _recipe_source_="${source_status[0]}${source_status[1]}$(_s r white)"
    local _recipe_status_="${recipe_status[0]}${recipe_status[1]}$(_s r white)"
    local _tool_status_="${tool_status[0]}${tool_status[1]}$(_s r white)"
    local _path_status_="${path_status[0]}${path_status[1]}$(_s r white)"
    local _recipe_name_="$(_s u b)${recipe_name}$(_s r white)"
    local _t_hr_="$(tput setaf 235)────────────────$(_s r white)"
    local _i_hr_="$(tput setaf 235)───────────────$(_s r white)"
    local _c_hr_="$(tput setaf 235)─────────────────$(_s r white)"
    local _e_hr_="$(tput setaf 235)──────────$(_s r white)"
    local _v_hr_=""
    local _tool_path_="$(_s grey i)${tool_path}$(_s r white)"
    local _image_="$(_s grey i)${docker_image}:${image_tag}$(_s r white)"
    local _cmd_="$(_s grey i)${cmd}$(_s r white)"
    local _entrypoint_="$(_s grey i)${entrypoint}$(_s r white)"

    # break down and style the volume data
    local _volumes_=
    if [ "" != "$volumes" ]; then
        local -a vols=($(echo "$volumes" | tr ";" "|"))
        local -a vol_parts
        local _src_
        local _dest_
        local _mode_

        for volume in $(echo "$volumes" | tr ";" "|"); do
            vol_parts=($(echo "$volume" | tr ":" "|"))
            _src_="$(_s white lt i)${vol_parts[0]}$(_s r white)"
            _dest_="$(_s white lt i)${vol_parts[1]}$(_s r white)"
            local mode=${vol_parts[2]}
            local mode_part
            local mode_icon
            local -a modes
            local a=0
            for mode_part in $(echo $mode | tr "," "|" | sort | uniq); do
                case $mode_part in
                    rw)
                        mode_part="$(_s red bt)${mode_part}$(_s r white)"
                        mode_icon="⇆"
                        ;;

                    ro)
                        mode_part="$(_s blue dk)${mode_part}$(_s r white)"
                        if [ "" == "$mode_icon" ]; then mode_icon="→"; fi
                        ;;

                    *)
                        mode_part="$(_s grey)${mode_part}$(_s r white)"
                        if [ "" == "$mode_icon" ]; then mode_icon="⇼"; fi
                        ;;
                esac
                if [ "" != "$mode_part" ]; then
                    modes[$a]=$mode_part
                    a=$((a + 1))
                fi
            done
            _mode_="$(printf "${modes[@]}" | tr " " ",")"
            if [ "" == "${_mode_}" ]; then
                _mode_="$(_s red bt)rw$(_s r white)"
            fi

            _volumes_="${_volumes_}
          ${mode_icon} ${_src_} : ${_dest_} : ${_mode_}"
            mode_part=
            mode_icon=
        done
    fi

    #
    # Template
    #
    ret_val="$(_s white)

 ${_recipe_source_} ${_recipe_name_} ${_recipe_status_}
     ├─ tool ${_t_hr_} ${_tool_path_} ${_tool_status_}
     ├─ image ${_i_hr_} ${_image_}
     ├─ cmd ${_c_hr_} ${_cmd_}
     ├─ entrypoint ${_e_hr_} ${_entrypoint_}
     └─ volumes${_v_hr_} ${_volumes_}"
#"──────────────────────────────────────────────────────────────────────────────"

    printf "$ret_val"

}


