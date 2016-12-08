
##############################################################################
##############################################################################
##
##  docker-tools command runner
##
##############################################################################
##############################################################################

function __main {
    local cmd_stdout
    local cmd_stderr
    local cmd_exit
    local pager

    # Initialize the application, ensure all required directories and files exist
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
        config|list|install|remove|update)
            cmd_stdout=$( (__docker_tools_${COMMAND} $@) 2> ${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err)
            cmd_exit=$?
            cmd_stderr=$(< ${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err)
            rm -f ${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err &>2 /dev/null
            if [ "list" == "$COMMAND" ]; then
                pager="less -r"
            fi
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

    if [ "" != "$pager" ]; then
        printf "$cmd_stdout" | $pager
    else
        printf "$cmd_stdout"
    fi

    if [ "" != "$cmd_stderr" ]; then
        >&2 printf "$cmd_stderr"
    fi
    exit $cmd_exit
}
