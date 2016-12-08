
##############################################################################
##############################################################################
##
##  Functions to manage the `docker-tools` configuration file
##
##  The confguration schema is set and should only allow modifications to
##  values already defined in the configuration file
##
##  Also includes documentation
##
##############################################################################
##############################################################################

declare __DOCKER_TOOLS_VERSION_DEFAULT__=master
declare __DOCKER_TOOLS_PREFIX_DEFAULT__=/usr/local/bin
declare __DOCKER_TOOLS_INSTALL_TMPFILE_DEFAULT__=/tmp/docker-tools-install.tmp
declare __DOCKER_TOOLS_ERROR_TMPFILE_DEFAULT__=/tmp/docker-tools-errors.tmp

#########################
#
#  documentation
#
#########################

function __config_usage {
        echo "
 Usage: $(_s b)docker-tools config$(_s r) <var> [value]

    <var>   Optional. If included, the specified $(_s u)variable$(_s r) will be displayed or
            updated, otherwise all $(_s u)variables$(_s r) and their values are displayed.
            If an optional $(_s u)value$(_s r) is provided, then the specified $(_s u)variable$(_s r) wiil
            be updated with that $(_s u)value$(_s r)."
}

function __config_help {
    echo "
$(_s b)NAME$(_s r)

    \`$(_s b)docker-tools config$(_s r)\` -- Configure docker-tools

$(_s b)USAGE$(_s r)

    $(_s b)docker-tools config$(_s r) [<$(_s b)var$(_s r)> [$(_s u)value$(_s r)]]

$(_s b)DESCRIPTION$(_s r)

    View or update a $(_s b)docker-tools$(_s r) configuration variable. If called without arguments, all
    variables and their values are listed

$(_s b)ARGUMENTS$(_s r)

    [<$(_s b)var$(_s r)> [$(_s u)value$(_s r)]]
        Optional, the name of the $(_s u)variable$(_s r) being configured. Available $(_s u)variables$(_s r) are:

            $(_s b)VERSION$(_s r)
            $(_s b)PREFIX$(_s r)
            $(_s b)INSTALL_TEMPFILE$(_s r)
            $(_s b)ERROR_TEMPFILE$(_s r)

        If omitted, all variabls and their values are listed.

        If a $(_s u)value$(_s r) is provided in addition to a $(_s u)variable$(_s r), the specified $(_s u)variable$(_s r) will be updated
        with the provided $(_s u)value$(_s r) otherwise the specified $(_s u)variable$(_s r) will be updated with this $(_s u)value$(_s r).

$(_s b)FLAGS$(_s r)

    $(_s b)-r$(_s r)
        Reset the specified $(_s u)variable$(_s r) to it's default $(_s u)value$(_s r). Default $(_s u)values$(_s r) are:

            $(_s b)VERSION$(_s r): master
            $(_s b)PREFIX$(_s r): /usr/local/bin
            $(_s b)INSTALL_TEMPFILE$(_s r): /tmp/docker-tools-install.tmp
            $(_s b)ERROR_TEMPFILE$(_s r): /tmp/docker-tools-errors.tmp
"
}

#
# Define the 'docker-tools config' command interface
#
function __docker_tools_config {
    local config_var=$(__get_arg 1 $@)
    local config_val=$(__get_arg 2 $@)
    local reset=$(__get_flag r $@)
    local -a vars
    local -a parts
    local a=0

    # List all current values
    if [ "" == "$config_var" ]; then
        __config_list

    # "Reset" a specific value
    elif [ 1 -eq "$(__get_flag r $@)" ]; then
        __config_delete $config_var

    # List a specific value
    elif [ "" == "$config_val" ]; then
        config_val="$(__config_get $config_var)"
        exit_code=$?
        if [ 0 -ne $exit_code ]; then
            >&2 printf $config_val
            exit $exit_code
        fi
        printf "$config_var=\n"

    # Set a value
    else
        __config_set $config_var $config_val
    fi
}

