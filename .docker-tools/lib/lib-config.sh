
##############################################################################
##############################################################################
##
##  Manage the `docker-tools` configuration file
##  The confguration schema is set and should only allow modifications to
##  values already defined in the configuration file
##
##############################################################################
##############################################################################

#
# Delete a config value
# Deleting just means seting it to ""
#
# @param Required, the name of the value to delete
#
function __config_delete {
    if [ "" == "$1" ]; then
        echo "__config_delete: You must specify a config value to delete"
        exit 1
    fi

    local -a ret_val
    local a=0
    local update_count=0

    while read line; do
        if [ "" != "${line:0:1}" ] \
            && [ "#" != "${line:0:1}" ] \
            && [ "$__DT_VERSION" != "${line%%=*}" ]
        then
            key=${line%%=*}
            key=${key/declare __DOCKER_TOOLS_/}
            key=${key/__/}

            if [ "$key" == "$1" ]; then
                update_count=$((update_count + 1))
                line="declare __DOCKER_TOOLS_${key}__"
            fi
        fi
        ret_val[$a]=$line
        a=$((a + 1))
    done < $DOCKER_TOOLS_CONFIG

    printf "%s\n" "${ret_val[@]}" > $DOCKER_TOOLS_CONFIG
    source $DOCKER_TOOLS_CONFIG
}

#
# Get a config value
#
# @param Required, the name of the value to get
#
function __config_get {
    if [ "" == "$1" ]; then
        echo "__config_get: You must specify a config value to get"
        exit 1
    fi

    local ret_val
    local a=0
    local set_count=0

    while read line; do
        if [ "" != "${line:0:1}" ] \
            && [ "#" != "${line:0:1}" ]
        then
            key=${line%%=*}
            key=${key/declare __DOCKER_TOOLS_/}
            key=${key/__/}

            value=${line/declare __DOCKER_TOOLS_${key}__/}
            value=${value#=*}
            if [ "$key" == "$1" ]; then
                if [ "" == "$value" ]; then
                    local varname="__DOCKER_TOOLS_${key}_DEFAULT__"
                    value=${!varname}
                fi
                ret_val=$value
                break
            fi

        fi
    done < $DOCKER_TOOLS_CONFIG

    if [ "" == "$ret_val" ]; then
        >&2 printf "Config value '$1' not defined"
        exit 1
    fi

    echo $ret_val
}

#
# Set a config value
# Only allow updates to existing values, not creation of new values
#
# @param Required, the name of the value to set
# @param Required, the value to set
#
function __config_set {
    if [ "" == "$1" ]; then
        echo "__config_set: You must specify a config value"
        exit 1
    fi
    if [ "" == "$2" ]; then
        echo "__config_set: You must provide a config value"
        exit 2
    fi

    local -a ret_val
    local a=0
    local update_count=0

    while read line; do
        if [ "" != "${line:0:1}" ] \
            && [ "#" != "${line:0:1}" ]
        then
            key=${line%%=*}
            key=${key/declare __DOCKER_TOOLS_/}
            key=${key/__/}

            value=${line/declare __DOCKER_TOOLS_${key}__/}
            value=${value#=*}

            if [ "$key" == "$1" ]; then
                line="declare __DOCKER_TOOLS_${key}__='${2/\'/\\\'}'"
                update_count=$((update_count + 1))
            fi
        fi
        ret_val[$a]=$line
        a=$((a + 1))
    done < $DOCKER_TOOLS_CONFIG

    if [ 0 -eq $update_count ]; then
        echo "__config_set: Key not found '$1'"
        exit 1
    fi

    printf "%s\n" "${ret_val[@]}" > $DOCKER_TOOLS_CONFIG

    source $DOCKER_TOOLS_CONFIG
}

#
# List current config values
#
# @return A newline delimited list of 'key=value' pairs
#
function __config_list {

    local -a ret_val
    local a=0

    while read line; do
        if [ "" != "${line:0:1}" ] \
            && [ "#" != "${line:0:1}" ]
        then
            key=${line%%=*}
            key=${key/declare __DOCKER_TOOLS_/}
            key=${key/__/}

            value=${line/declare __DOCKER_TOOLS_${key}__/}
            value=${value#=*}

            if [ "" == "$value" ]; then
                local varname="__DOCKER_TOOLS_${key}_DEFAULT__"
                value=${!varname}
            fi

            ret_val[$a]="${key}=${value}"
            a=$((a + 1))
        fi
    done < $DOCKER_TOOLS_CONFIG

    printf "%s\n" "${ret_val[@]}"
}


