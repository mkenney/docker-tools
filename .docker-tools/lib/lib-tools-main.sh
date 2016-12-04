
##############################################################################
##############################################################################
##
##     ###########  ##        ##  ##########
##   #############  ###      ###  ############
##  ###             ####    ####  ##        ###
##  ##              ## ##  ## ##  ##         ##
##  ##              ##  ####  ##  ##         ##
##  ##              ##   ##   ##  ##         ##
##  ##              ##        ##  ##         ##
##  ###             ##        ##  ##        ###
##   #############  ##        ##  ############
##     ###########  ##        ##  ##########
##
##############################################################################
##############################################################################

##############################################################################
##############################################################################
##
##  docker-tools commands
##
##############################################################################
##############################################################################

#
# config
#
function __docker_tools_config {
    local config_var=$(__get_arg 1 $@)
    local config_val=$(__get_arg 2 $@)
    local reset=$(__get_flag r $@)
    local -a vars
    local -a parts
    local a=0

    # "Reset" a specific value
    if [ 1 -eq "$(__get_flag r $@)" ]; then
        __config_delete $config_var

    # List current values
    elif [ "" == "$config_var" ]; then
        __config_list

    # List a specific value
    elif [ "" == "$config_val" ]; then
        echo "$config_var=$(__config_get $config_var)"

    # Set a value
    else
        __config_set $config_var $config_val
    fi
}



#
# list
#
# @option --source
#
function __docker_tools_list {
    local ret_val
    local -a parts

    local -a recipe_files=("$DOCKER_TOOLS_RECIPES" "$DOCKER_TOOLS_REGISTRY")
    local recipe_file
    local recipe_line
    local recipe_name
    local -a recipes
    local recipe
    local a=0

    for recipe_file in "${recipe_files[@]}"; do
        while read recipe_line; do
            if [ "" != "${recipe_line:0:1}" ] && [ "#" != "${recipe_line:0:1}" ]; then
                parts=($(echo "$recipe_line" | tr "$__recipe_delimiter__" "\n"))
                recipe_name="${parts[0]}"
                recipes[$a]="${recipe_name}:$(basename ${recipe_file})"
                a=$((a + 1))
            fi
        done < $recipe_file
    done

    for recipe in $(echo "${recipes[@]}" | tr " " "\n" | sort | uniq); do
        recipe_name=${recipe%%:*}
        recipe_file=${recipe#*:}
        ret_val="$ret_val $(__recipe_describe $recipe_name --source=$recipe_file)"
    done

    echo "$ret_val" | less -r
}

#
# install
#
function __docker_tools_install {
    local old_IFS=IFS

    ########################
    # 2 modes
    #   if no recipe is specified, then at minimum --image and --name options
    #   are required
    #
    #   if a recipe is specified, no options are required, but any may be
    #   passed as a 1-time override of the corresponding recipe value
    #########################

    local recipe_name="$(__get_arg 1 $@)"
    local opt_name="$(__get_opt name $@)"
    local opt_prefix="$(__get_opt prefix $@)"
    local opt_type="$(__get_opt type $@)"
    local opt_image="$(__get_opt image $@)"
    local opt_tag="$(__get_opt tag $@)"
    local opt_entrypoint="$(__get_opt entrypoint $@)"
    local opt_cmd="$(__get_opt cmd $@)"
    local opt_volumes="$(__get_opt volumes $@)"

    if [ "" == "$recipe_name" ]; then
        if [ "" == "$opt_image" ] || [ "" == "$opt_name" ]; then
            __show_usage install " Error: A recipe name must be specified or --image AND --name options are
 required"
            exit 1
        fi
    fi
 Error: A recipe name must be specified or --image AND --name options are
 required
    # Tool recipe
    local -a recipe
    IFS=$"$__recipe_delimiter__"
    recipe=($(__recipe_get $recipe_name))
    IFS=$old_IFS
    if [ "" != "$opt_cmd" ];        then recipe[6]="$opt_cmd";        fi
    if [ "" != "$opt_entrypoint" ]; then recipe[5]="$opt_entrypoint"; fi
    if [ "" != "$opt_image" ];      then recipe[3]="$opt_image";      fi
    if [ "" != "$opt_tag" ];        then recipe[4]="$opt_tag";        fi
    if [ "" != "$opt_name" ];       then recipe[1]="$opt_name";       fi
    if [ "" != "$opt_prefix" ];     then recipe[2]="$opt_prefix";     fi
    if [ "" != "$opt_type" ];       then recipe[3]="$opt_prefix";     fi
    if [ "" != "$opt_volumes" ];    then recipe[7]="$opt_volumes";    fi

    if [ "" == "$opt_name" ]; then
        echo "Error - a tool name was not provided"
        exit 1
    fi


    # Tool tempfile
    cp $DOCKER_TOOLS_LIB_DIR/templates/tool.sh $__INSTALL_TMPFILE__

        # tool info
    sed -i "s/declare __TOOLS_VERSION__=/declare __TOOLS_VERSION__=$DOCKER_TOOLS_VERSION/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_NAME__=/declare __RECIPE_NAME__=$recipe_name/" $__INSTALL_TMPFILE__
    sed -i "s/declare __TOOLS_LIB_DIR__=/declare __TOOLS_LIB_DIR__=$DOCKER_TOOLS_LIB_DIR/" $__INSTALL_TMPFILE__

        # tool recipe
    sed -i "s/declare __RECIPE_CMD__=/declare __RECIPE_CMD__=$opt_cmd/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_ENTRYPOINT__=/declare __RECIPE_ENTRYPOINT__=$opt_entrypoint/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_IMAGE__=/declare __RECIPE_IMAGE__=$opt_image/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_IMAGE_TAG__=/declare __RECIPE_IMAGE_TAG__=$opt_tag/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_TOOL__=/declare __RECIPE_TOOL__=$opt_name/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_PREFIX__=/declare __RECIPE_PREFIX__=$opt_prefix/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_VOLUMES__=/declare __RECIPE_VOLUMES__=$opt_volumes/" $__INSTALL_TMPFILE__

    if [ "" == "$opt_prefix" ]; then opt_prefix=$DOCKER_TOOLS_PREFIX; fi
    #result=$( (cat ${__INSTALL_TMPFILE__} > $opt_prefix/$opt_name) 2>${__ERROR_TMPFILE__})
    #exit_code=$?
    #errors=$(< ${__ERROR_TMPFILE__})


echo "result:$result"
echo "exit_code:$exit_code"
echo "errors:$errors"
#sed -i "s/php_value newrelic.appname emt_web-dev/php_value newrelic.appname $APP_NAME/" /var/www/html/documentroot/.htaccess







echo "__docker_tools_install:recipe:${recipe[@]}"
}



#
# Supporting directories and files
#

#
# @todo if these don't exist, run first-run routines...
# @todo write first-run routines:
#       - Check for each required file, if it doesn't exist download it... maybe
#         __init should always do all that and should run before these are sourced
#



function __main {
    local cmd_stdout
    local cmd_stderr
    local cmd_exit

    __init

    # define the command and remove it from the argument list
    local COMMAND="$(__get_arg 1 $@)"
    eval $(__shift_args $(__get_arg_pos 1 $@))

    # show usage
    if [ 1 -eq $(__get_flag h $@) ]; then
        __show_usage $COMMAND; exit 0
    fi

    # show help
    if [ "" != "$(__get_opt help $@)" ]; then
        __show_help $COMMAND; exit 0
    fi

    case $COMMAND in
        # execute pre-defined commands
        config|list|install|remove|update|update)
            cmd_stdout=$( (__docker_tools_${COMMAND} $@) 2> ${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err)
            cmd_exit=$?
            cmd_stderr=$(< ${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err)
            ;;

        # show usage / help
        "")
            if [ 1 -eq $(__get_flag h $@) ]; then
                __show_usage
                exit 0
            fi
            __show_help
            ;;

        # invalid commands
        *)
            echo "Unknown command: '$COMMAND'"
            __usage
            exit 1
    esac

    echo $cmd_stdout
    >&2 echo $cmd_stderr
    exit $cmd_exit
}
