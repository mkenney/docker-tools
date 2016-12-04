
##############################################################################
##############################################################################
##
##  CLI - Common functions useful for parsing cli arguments, etc.
##
##
##
##############################################################################
##############################################################################

#
# In case of poorly constructed generated code...
#
function __ {
    echo "Invalid function '__'"
    exit 1
}

#
# Get the position of the argument number in the argument list
# Useful for shifting arguments around
#
# @param filter
# @param $@
# @return Number or ""
#
function __get_arg_pos {
    local filter=$1
    shift

    local ret_val
    local arg_num=1
    local param_num=1
    local var

    for var in "$@"; do
        if [ "-" != "${var:0:1}" ]; then
            if [ "$filter" == "$arg_num" ]; then
                ret_val=$param_num
                break;
            fi
            arg_num=$((arg_num + 1))
        fi
        param_num=$((param_num + 1))
    done

    printf "$ret_val"
}

#
# Create a `set` command to shift positional arguments out of the set.
#
# Accepts any number of arguments. Like `shift` but removes specified arguments
# from the middle of the arguments list when eval'd
#
# i.e. remove argumets 2 and 4 from the parameter list:
#   before: "$@" = a b c d e
#       `eval $(__shift_args 2 4)`
#   after: "$@" = a c e
#
function __shift_args {
    local ret_val
    local arg_ids=$(echo "$@" | tr " " "\n" | sort -g | uniq)
    local arg_id
    local start=1
    local arg_count

    for arg_id in $arg_ids; do
        arg_count=$((arg_id - $((start))))
        if [ $arg_count -gt 0 ]; then
            ret_val="$ret_val \${@:$start:$arg_count}"
        fi
        start=$((arg_id + 1))
    done
    ret_val="$ret_val \${@:$start}"

    printf "set -- \"$ret_val\""
}

#
# Get command-line arguments
#
# @param $@
# @return a space-delimited list of each option in the order it was given.
#         Duplicates are included.
#
function __get_args {
    ret_val=
    for var in "$@"; do
        if [ "-" != "${var:0:1}" ]; then
            if [ "" == "$ret_val" ]; then
                ret_val="$var"
            else
                ret_val="$ret_val $var"
            fi
        fi
    done

    printf "$(echo "$ret_val" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')"
}

#
# Get a command-line argument by index
#
# @param argument number
# @param $@
# @return a space-delimited list of each option in the order it was given.
#         Duplicates are included.
#
function __get_arg {
    local argnum=$1
    shift
    local ret_val
    local a=1
    for var in $(__get_args $@); do
        if [ $a -eq $argnum ]; then
            ret_val=$var
            break
        fi
        a=$((a + 1))
    done

    printf "$ret_val" #| sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//'
}

#
# Get the index of a specified command-line argument
#
# @param argument
# @param $@
# @return A number or "" if the argument doesn't exist
#
function __get_argnum {
    if [ "" == "$1" ]; then
        echo "An argument is required"
    fi

    local arg_filter=$1
    local arg
    local a=0

    for arg in "$(__get_args)"; do
        a=$((a + 1))
        if [ "$arg_filter" == "$arg" ]; then
            break
        fi
    done

    printf "$a"
}

#
# Get command-line options
# An option is any argument that begins with '--'
#
# @param $@
# @return a unique, sorted, space-delimited list of each option
#
function __get_opts {
    local ret_val=
    for var in "$@"; do
        if [ "--" == "${var:0:2}" ]; then
            if [ "" == "$ret_val" ]; then
                ret_val="${var:2}"
            else
                ret_val="$ret_val ${var:2}"
            fi
        fi
    done

    printf "$(echo $ret_val | tr " " "\n" | sort | uniq | tr "\n" " " | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')"
}

#
# Get the value of a specified command-line option
# An option is any argument that begins with '--'
# A value is anything associated to the argument with '='
#
# @param string The name of the option to retrieve the value for
# @param $@
# @return The specified value or "" if it doesn't exist
#
function __get_opt {
    if [ "" == "$1" ]; then
        echo "No option provided"
        exit 1
    fi

    local ret_val
    local opt
    local option
    local value

    for opt in "$(__get_opts $@)"; do
        option=${opt%%=*}
        value=${opt#*=}
        if [ "" == "$value" ]; then value=1; fi
        case $option in
            $1)
                ret_val="$value"
                ;;
        esac
    done

    printf "$ret_val"
}

#
# Get command-line flags
# Flags are made up of argument that begins with '-'
#
# @param $@
# @return a unique, sorted, space-delimited list of each flag
#
function __get_flags {
    local -a ret_val
    local var
    local flags
    local a=0
    local b=0

    for var in "$@"; do
        if [ "-" == "${var:0:1}" ] && [ "-" != "${var:1:1}" ]  && [ "" != "${var:1}" ]; then
            flags=${var:1}
            for (( a=0; a<${#flags}; a++ )); do
                echo "b: $b; flags:$a:1: ${flags:$a:1}"
                ret_val[$b]=${flags:a:1}
                b=$((b + 1))
            done
        fi
    done

    printf "%s\n" "${ret_val[@]}" | sort | uniq | tr "\n" " " | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//'
}

#
# See if a flag was passed
#
# @param flag to search for
# @param $@
# @return 1 for true, 0 for false
#
function __get_flag {
    local ret_val=0
    local flag=$1
    shift
    local flags="$(__get_flags $@)"

    if [[ $flags == *"$flag"* ]]; then
        ret_val=1
    fi
    printf "$ret_val"
}

#
# Get the tool install prefix based on a combination of command arguments and
# the user $PATH.
#
# @param Name of the tool being installed
# @param $@
# @return The logical installation path
#
function __get_install_prefix {
    local tool="$(__get_arg 1)"
    local prefix="$(__get_opt prefix $@)"

    if [ "" == "$prefix" ]; then
        if [ "" != "$tool" ]; then
            tmp=$(which $tool)
            if [ 0 -eq $? ] && [ "" != "$tmp" ]; then
                prefix=$(dirname $tmp)
            fi
        fi
    fi

    printf "$prefix"
}
