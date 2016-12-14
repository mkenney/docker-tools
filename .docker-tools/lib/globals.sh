
##############################################################################
##############################################################################
##
##  User Configurable Variables
##
##  This is just a list for the sake of documentation. These should be set in
##  your .bashrc or similar, executing a `self-update` command will overwrite
##  this file.
##
##############################################################################
##############################################################################

# Define the location of the docker-tools metadata directory. The default
# value is `$HOME/.docker-tools`
#export DOCKER_TOOLS_CONFIG_DIR=[PATH]


##############################################################################
##############################################################################
##
##  Globals and defaults
##
##  Probably not a good idea to modify these values because executing a
##  `self-update` command will overwrite this file.
##
##############################################################################
##############################################################################

# docker-tools config
# Always allow ENV override

if [ "" == "$DOCKER_TOOLS_VERSION" ]; then
    if [ "" == "$__DOCKER_TOOLS_VERSION__" ]; then
        DOCKER_TOOLS_VERSION=master
    else
        DOCKER_TOOLS_VERSION=$__DOCKER_TOOLS_VERSION__
    fi
fi
if [ "" == "$DOCKER_TOOLS_PREFIX" ]; then
    if [ "" == "$__DOCKER_TOOLS_PREFIX__" ]; then
        DOCKER_TOOLS_PREFIX=/usr/local/bin
    else
        DOCKER_TOOLS_PREFIX=$__DOCKER_TOOLS_VERSION__
    fi
fi
if [ "" == "$DOCKER_TOOLS_INSTALL_TMPFILE" ]; then
    if [ "" == "$__DOCKER_TOOLS_INSTALL_TMPFILE__" ]; then
        DOCKER_TOOLS_INSTALL_TMPFILE=/tmp/docker-tools-install.tmp
    else
        DOCKER_TOOLS_INSTALL_TMPFILE=$__DOCKER_TOOLS_INSTALL_TMPFILE__
    fi
fi
if [ "" == "$DOCKER_TOOLS_ERROR_TMPFILE" ]; then
    if [ "" == "$__DOCKER_TOOLS_ERROR_TMPFILE__" ]; then
        DOCKER_TOOLS_ERROR_TMPFILE=/tmp/docker-tools-errors.tmp
    else
        DOCKER_TOOLS_ERROR_TMPFILE=$__DOCKER_TOOLS_ERROR_TMPFILE__
    fi
fi

#####################################
# Configuration directories and files
#####################################

# Allow overrides
if [ "" == "$DOCKER_TOOLS_CONFIG_DIR" ]; then
    DOCKER_TOOLS_CONFIG_DIR=$HOME/.docker-tools
fi

export DOCKER_TOOLS_VERSION
export DOCKER_TOOLS_CONFIG_DIR
export DOCKER_TOOLS_PREFIX
export DOCKER_TOOLS_LIB_DIR=$DOCKER_TOOLS_CONFIG_DIR/lib
export DOCKER_TOOLS_CONFIG=$DOCKER_TOOLS_CONFIG_DIR/config
export DOCKER_TOOLS_REGISTRY=$DOCKER_TOOLS_CONFIG_DIR/registry.old
export DOCKER_TOOLS_RECIPES=$DOCKER_TOOLS_CONFIG_DIR/recipes.old

##################
# Remote resources
##################
declare __DOCKER_TOOLS_INSTALLER_URL__=https://raw.githubusercontent.com/mkenney/docker-tools/$DOCKER_TOOLS_VERSION/install.sh
declare __DOCKER_TOOLS_CONFIG_URL__=https://raw.githubusercontent.com/mkenney/docker-tools/$DOCKER_TOOLS_VERSION/.docker-tools
