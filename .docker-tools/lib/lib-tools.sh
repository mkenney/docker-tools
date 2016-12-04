##############################################################################
##############################################################################
##
##  Documentation
##
##############################################################################
##############################################################################

#
# @param command
# @param string description of what triggered it, if any
#
function __show_usage {
    local usage_fnc="__usage"
    local usage_txt
    local pager="$(__get_opt pager $@)"

    if [ "" != "$1" ]; then usage_fnc="__${1}_usage"; fi
    if [ "" != "$2" ]; then
        usage_txt="\n${2}\n$($usage_fnc)"
    else
        usage_txt="$($usage_fnc)"
    fi

    if [ "" != "$pager" ]; then
        printf "$usage_txt" | $pager
    else
        printf "$usage_txt\n"
    fi
}

#
# @param command
# @param string description of what triggered it, if any
#
function __show_help {
    local usage_fnc="__help"
    local usage_txt
    local pager="$(__get_opt pager $@)"
    local header="                                  DOCKER-TOOLS Commands Manual"
    local footer="MIT                                       NOV 30, 2016                                          MIT"

    if [ "" != "$1" ]; then usage_fnc="__${1}_help"; fi
    if [ "" == "$pager" ]; then pager="less -r"; fi
    if [ "" != "$2" ]; then
        usage_txt="\n${2}\n$($usage_fnc)"
    else
        usage_txt="$($usage_fnc)"
    fi

    printf "${header}\n$usage_txt\n\n$footer" | $pager
}

#########################
#
#  docker-tools
#
#########################

function __usage {
        echo "
usage: docker-tools [COMMAND]
        "
}

function __help {
    echo "
$(_s b)NAME$(_s r)

    $(_s b)docker-tools$(_s r) -- shell script generator

$(_s b)USAGE$(_s r)

    $(_s b)docker-tools$(_s r) $(_s u)COMMAND$(_s r) [options]

$(_s b)DESCRIPTION$(_s r)

    Create and manage Docker container wrapper scripts ($(_s u)tools$(_s r)). Create, save and manage '$(_s u)tool$(_s r)'
    configurations (sets of values that define a \`docker run\` command) or \"$(_s u)recipes$(_s r)\" and install
    or uninstall shell scripts ($(_s u)tools$(_s r)) based on them. Generally useful for system utilities and
    dev tools (php-cli, node, python-cli, etc.).

$(_s b)COMMANDS$(_s r)

    See \`$(_s b)docker-tools$(_s r) $(_s u)COMMAND$(_s r) --help\` for command usage.

    Things you can do with \`$(_s b)docker-tools$(_s r)\`

        $(_s b)config$(_s r)
            Manage $(_s b)docker-tools$(_s r) configuration values

        $(_s b)self-update$(_s r)
            Update the $(_s b)docker-tools$(_s r) script

    Things you can do with recipes

        $(_s b)create$(_s r)
            Create or update a tool $(_s u)recipe$(_s r)

        $(_s b)list$(_s r)
            Display installed and/or registered $(_s u)recipes$(_s r)

        $(_s b)delete$(_s r)
            Delete a specified $(_s u)recipe$(_s r)

    Things you can do with tools

        $(_s b)install$(_s r)
            Install a $(_s u)tool$(_s r)

        $(_s b)uninstall$(_s r)
            Uninstall a $(_s u)tool$(_s r)

        $(_s b)update$(_s r)
            Update an existing $(_s u)tool$(_s r)


$(_s b)EXAMPLES$(_s r)

    $ $(_s i)docker-tools install gulp --tag=7.0-alpine$(_s r)

$(_s b)SHELL VARIABLES$(_s r)
    The following variables can be set in your shell to modify $(_s b)docker-tools$(_s r):

        $(_s b)DOCKER_TOOLS_CONFIG_DIR$(_s r)
            Define the path to the $(_s b)docker-tools$(_s r) confguration directory. The configuration directory
            is where $(_s b)docker-tools$(_s r) stores metadata, program libraries and stored tool $(_s u)recipes$(_s r) as well
            as the $(_s b)docker-tools$(_s r) $(_s u)recipe$(_s r) registry.

        $(_s b)DOCKER_TOOLS_PREFIX$(_s r)
            This is the default installation directory for tools unless specified as an argument or
            in a $(_s u)recipe$(_s r)."
}

