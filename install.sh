#!/usr/bin/env bash

##############################################################################
##############################################################################
##
##  Globals and defaults
##
##############################################################################
##############################################################################

SELF=$0
DEFAULT_PATH=/usr/local/bin

#__DT_INSTALLER_URL=https://raw.githubusercontent.com/mkenney/docker-tools/stable/install.sh
__DT_INSTALLER_URL=https://raw.githubusercontent.com/mkenney/docker-tools/master/install.sh
DOCKER_TOOLS_INSTALLER_LOCAL=/tmp/docker-tools-intaller.sh

#DOCKER_TOOLS_EXECUTABLE_URL=https://raw.githubusercontent.com/mkenney/docker-tools/stable/bin/docker-tools
DOCKER_TOOLS_EXECUTABLE_URL=https://raw.githubusercontent.com/mkenney/docker-tools/master/bin/docker-tools
DOCKER_TOOLS_EXECUTABLE_TMPFILE=/tmp/docker-tools.tmp
DOCKER_TOOLS_EXECUTABLE_NAME=docker-tools
DOCKER_TOOLS_EXECUTABLE_PATH=


##############################################################################
##############################################################################
##
##  Lib
##
##############################################################################
##############################################################################

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
    echo $ret_val | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//'
}

function __get_opts {
    declare ret_val=
    for var in "$@"; do
        if [ "--" == "${var:0:2}" ]; then
            if [ "" == "$ret_val" ]; then
                ret_val="${var:2}"
            else
                ret_val="$ret_val ${var:2}"
            fi
        fi
    done
    echo $(echo $ret_val | tr " " "\n" | sort | uniq | tr "\n" " " | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')
}

function __get_flags {
    declare ret_val=
    for var in "$@"; do
        if [ "-" == "${var:0:1}" ] && [ "-" != "${var:1:1}" ]  && [ "" != "${var:1}" ]; then
            if [ "" == "$ret_val" ]; then
                ret_val="${var:1}"
            else
                ret_val="$ret_val ${var:1}"
            fi
        fi
    done
    echo $ret_val | tr " " "\n" | sort | uniq | tr "\n" " " | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//'
}

#
# Usage
#
function __usage {
    if [ "sh" == "$SELF" ] || [ "bash" == "$SELF" ]; then
        SELF="bash -s"
    fi

    echo "
    Usage: install.sh COMMAND [--prefix=PATH]

    Commands:
        install        Install the \`$DOCKER_TOOLS_EXECUTABLE_NAME\` command
        update         Update the \`$DOCKER_TOOLS_EXECUTABLE_NAME\` command
        remove         Delete the \`$DOCKER_TOOLS_EXECUTABLE_NAME\` command

    Options:
        --prefix=PATH  Optional, the location to use for the \`$DOCKER_TOOLS_EXECUTABLE_NAME\`
                       command. If omitted, your PATH will be searched.
                       Default '/usr/local/bin'.

    Examples:
        $ curl -L $__DT_INSTALLER_URL | bash -s install --path=\$HOME/bin
        $ curl -L $__DT_INSTALLER_URL | bash -s update
        $ curl -L $__DT_INSTALLER_URL | bash -s remove
"
}
function __help {
    if [ "sh" == "$SELF" ] || [ "bash" == "$SELF" ]; then
        SELF="bash -s"
    fi

    echo "
    Name
        install.sh -- \`$DOCKER_TOOLS_EXECUTABLE_NAME\` install manager

    Synopsys
        install.sh COMMAND [--prefix=PATH]

    Summary
        Manage your $DOCKER_TOOLS_EXECUTABLE_NAME installation

    Commands
        install
                Install the \`$DOCKER_TOOLS_EXECUTABLE_NAME\` command

        remove
                Install the \`$DOCKER_TOOLS_EXECUTABLE_NAME\` command

        update
                Update the \`$DOCKER_TOOLS_EXECUTABLE_NAME\` command

    Common Options
        All commands accept the options listed below

        --prefix=PATH
                Optional, the location to place the \`$DOCKER_TOOLS_EXECUTABLE_NAME\` executable into, defaults to
                '/usr/local/bin'

    Examples
        $ curl -L $__DT_INSTALLER_URL | bash -s install
"
}


##############################################################################
##############################################################################
##
##  Download the master install script and execute it locally
##
##  Loosely based on https://npmjs.org/install.sh, run as `curl | sh`
##  http://www.gnu.org/s/hello/manual/autoconf/Portable-Shell.html
##
##############################################################################
##############################################################################
if [ "sh" == "$SELF" ] || [ "bash" == "$SELF" ]; then

    #
    # Download and execute the install script
    #
    curl -f -L -s $__DT_INSTALLER_URL > $DOCKER_TOOLS_INSTALLER_LOCAL
    exit_code=$?
    if [ 0 -ne $exit_code -eq 0 ]; then
        echo
        echo "Install failed: Could not download 'install.sh' from $__DT_INSTALLER_URL" >&2
        exit $exit_code

    else
        if head $DOCKER_TOOLS_INSTALLER_LOCAL | grep -q '404: Not Found'; then
            echo "Install failed: The installation script could not be found at $__DT_INSTALLER_URL"  >&2
            rm -f $DOCKER_TOOLS_INSTALLER_LOCAL
            exit 404
        fi
        if ! [ -s $DOCKER_TOOLS_INSTALLER_LOCAL ]; then
            echo
            echo "Install failed: Invalid or empty script at $__DT_INSTALLER_URL" >&2
            exit 1
        fi
        (exit 0) # Reset $?
    fi

    bash $DOCKER_TOOLS_INSTALLER_LOCAL $@
    exit_code=$?
    rm -f $DOCKER_TOOLS_INSTALLER_LOCAL
    exit $exit_code
fi


##############################################################################
##############################################################################
##
##  asdf
##
##############################################################################
##############################################################################

#
INSTALL_ARG="$(__get_args $@)"
case $INSTALL_ARG in
    install|remove|update)
        ;;
    *)
        echo "Invalid command: '$INSTALL_ARG'"
        __usage
        exit 1
esac

INSTALL_OPT_PREFIX=
for opt in "$(__get_opts $@)"; do
    option=${opt%%=*}
    value=${opt#*=}

    case $option in

        # Sanitize any prefix input
        prefix)
            INSTALL_OPT_PREFIX=$opt
            dir=$(dirname $value)
            value=$(basename $value)
            if [ "/" == "$value" ]; then value= ; fi

            if [ "" == "$dir" ]; then
                dir='./'
            elif [ "/" != "$dir" ]; then
                dir="$dir/"
            fi
            DOCKER_TOOLS_EXECUTABLE_PATH="${dir}${value}"
            ;;

        *)
            if [ "" != "$opt" ]; then
                echo "Option provided but not defined: --$option"
                __usage
                exit 1
            fi
    esac
done

#
# No install prefix specified, search the user's path for an existing install
#
if [ "" == "$INSTALL_OPT_PREFIX" ]; then
    tmp=$(which $DOCKER_TOOLS_EXECUTABLE_NAME)
    if [ 0 -eq $? ] && [ "" != "$tmp" ]; then
        DOCKER_TOOLS_EXECUTABLE_PATH=$(dirname $tmp)
    fi
fi
if [ "" == "$DOCKER_TOOLS_EXECUTABLE_PATH" ]; then
    DOCKER_TOOLS_EXECUTABLE_PATH=$DEFAULT_PATH
fi

case $INSTALL_ARG in

    #
    # Install the `docker-tools` command
    #
    install|update)

        # Download and validate the script
        curl -f -L -s $DOCKER_TOOLS_EXECUTABLE_URL > $DOCKER_TOOLS_EXECUTABLE_TMPFILE
        exit_code=$?
        if [ $exit_code -ne 0 ]; then
            echo
            echo "$INSTALL_ARG failed: Could not download '$DOCKER_TOOLS_EXECUTABLE_URL'"
            exit $exit_code
        fi
        if grep -q '404: Not Found' $DOCKER_TOOLS_EXECUTABLE_TMPFILE; then
            usage
            echo
            echo "Not found: $DOCKER_TOOLS_EXECUTABLE_URL";
            exit 404
        fi
        if ! [ -s $DOCKER_TOOLS_EXECUTABLE_TMPFILE ]; then
            echo
            echo "$INSTALL_ARG failed: Invalid or empty script or download failed -- $DOCKER_TOOLS_EXECUTABLE_URL > $DOCKER_TOOLS_EXECUTABLE_TMPFILE"
            exit $exit_code
        fi
        (exit 0)

        # Create the installation directory
        mkdir -p "$DOCKER_TOOLS_EXECUTABLE_PATH"
        exit_code=$?
        if [ 0 -ne $exit_code ]; then
            echo
            echo "$INSTALL_ARG failed: Could not create directory '$DOCKER_TOOLS_EXECUTABLE_PATH'"
            exit $exit_code
        fi
        (exit 0)

        # Cat the tempfile into the command file instead of moving it so that
        # symlinkys aren't destroyed
        result=$( (cat $DOCKER_TOOLS_EXECUTABLE_TMPFILE > $DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME) 2>${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err)
        exit_code=$?
        errors=$(< ${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err)
        if [ 0 -ne $exit_code ] || [ "" != "$errors" ]; then
            echo "$errors"
            echo
            echo "$INSTALL_ARG failed: Could not write to '$DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME'"
            if [ "${my_string/denied}" == "$my_string" ]; then
                echo "   ...do you need more sudo?"
            fi
            rm -f ${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err
            exit $exit_code
        fi
        (exit 0)

        # Set the execute bit
        result=$( (chmod +x $DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME) 2>${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err)
        exit_code=$?
        errors=$(< ${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err)
        if [ 0 -ne $exit_code ] || [ "" != "$errors" ]; then
            echo "$errors"
            echo
            echo "$INSTALL_ARG failed: Could not set execute bit on '$DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME'"
            if [ "${my_string/permissions}" == "$my_string" ]; then
                echo "   ...do you need more sudo?"
            fi
            rm -f ${DOCKER_TOOLS_EXECUTABLE_TMPFILE}.err
            exit $exit_code
        fi
        (exit 0)

        # Cleanup the tempfile
        rm -f $DOCKER_TOOLS_EXECUTABLE_TMPFILE
        exit_code=$?
        if [ 0 -ne $exit_code ]; then
            echo
            echo "Error: Could not delete tempfile '$DOCKER_TOOLS_EXECUTABLE_TMPFILE'"
            exit $exit_code
        fi
        (exit 0)

        # w00t
        echo
        echo "$DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME: $INSTALL_ARG succeeded"
        exit 0
        ;;

    #
    # Remove the `docker-tools` command
    #
    remove)
        if [ ! -f "$DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME" ]; then
            echo "$DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME does not exist"
            exit 1
        fi
        if [ -L "$DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME" ]; then
            orig=$(readlink -f "$DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME")
            echo "$DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME is a symbolic link to $orig"
            echo "Removing $orig"
            rm -f "$orig"
            error_code=$?
            if [ 0 -ne $error_code ]; then
                echo "There was an error deleting $orig"
                exit $error_code
            fi
        fi
        echo "Removing $DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME"
        rm -f "$DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME"
            error_code=$?
            if [ 0 -ne $error_code ]; then
                echo "There was an error deleting $DOCKER_TOOLS_EXECUTABLE_PATH/$DOCKER_TOOLS_EXECUTABLE_NAME"
                exit $error_code
            fi
        echo "Done."
        exit 0
        ;;
esac
