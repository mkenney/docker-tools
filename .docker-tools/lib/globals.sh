
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

export DOCKER_TOOLS_VERSION=v0.0.1

#####################################
# Configuration directories and files
#####################################

# Allow overrides
if [ "" == "$DOCKER_TOOLS_CONFIG_DIR" ]; then
    DOCKER_TOOLS_CONFIG_DIR=$HOME/.docker-tools
fi
if [ "" == "$DOCKER_TOOLS_PREFIX" ]; then
    DOCKER_TOOLS_PREFIX=/usr/local/bin
fi

export DOCKER_TOOLS_CONFIG_DIR
export DOCKER_TOOLS_PREFIX
export DOCKER_TOOLS_LIB_DIR=$DOCKER_TOOLS_CONFIG_DIR/lib
export DOCKER_TOOLS_CONFIG=$DOCKER_TOOLS_CONFIG_DIR/config
export DOCKER_TOOLS_REGISTRY=$DOCKER_TOOLS_CONFIG_DIR/registry
export DOCKER_TOOLS_RECIPES=$DOCKER_TOOLS_CONFIG_DIR/recipes
