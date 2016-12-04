
##############################################################################
##############################################################################
##
##  Functions to manage the docker-tools "list" command
##
##  Also includes documentation
##
##############################################################################
##############################################################################

#########################
#
#  documentation
#
#########################

function __list_usage {
        echo "
usage: docker-tools list ...
        "
}

function __list_help {
    echo "
    $(_s u)Name:$(_s r)

        \`$(_s b)docker-tools list$(_s r)\` -- list available recipes

    $(_s u)Usage:$(_s r)

        docker-tools list

    $(_s u)Description:$(_s r)

        List available recipes.

    $(_s u)Options:$(_s r)
        --source

    $(_s u)Examples:$(_s r)

        $ docker-tools list
"
}


