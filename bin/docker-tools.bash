#!/usr/bin/env bash

##############################################################################
##############################################################################
##
##  docker-tools
##
##############################################################################
##############################################################################

if [ "" == "$DOCKER_TOOLS_CONFIG_DIR" ]; then
    #DOCKER_TOOLS_CONFIG_DIR=$HOME/.docker-tools
DOCKER_TOOLS_CONFIG_DIR=$HOME/my_projects/docker-tools/.docker-tools
fi

##############################################################################
# Do any required initialization
#
# - Ensure the minimum requirements exist and if any are missing attempt to
#   download them using `curl`
##############################################################################

declare -a __docker_tools_conf_dirs__=(
    "lib"
    "lib/templates"
    "lib/templates/tool"
    "lib/templates/service"
)
declare -a __docker_tools_conf_files__=(
    "lib/templates/tool.sh"
    "lib/templates/tool/data.sh"
    "lib/templates/tool/template.sh"
    "lib/globals.sh"
    "lib/lib-cli.sh"
    "lib/lib-config.sh"
    "lib/lib-recipes.sh"
    "lib/lib-registry.sh"
    "lib/lib-tools-install.sh"
    "lib/lib-tools-list.sh"
    "lib/lib-tools-main.sh"
    "lib/lib-tools.sh"
    "lib/lib-ui.sh"
    "config"
    "registry"
)
function __init {
    local conf_dir
    local conf_file

    # Make sure the configuration directory exists
    if [ ! -d "$DOCKER_TOOLS_CONFIG_DIR" ]; then
        mkdir -pv "$DOCKER_TOOLS_CONFIG_DIR"
        exit_code=$?
        if [ 0 -ne $exit_code ]; then
            >&2 echo "Could not create configuration directory '$DOCKER_TOOLS_CONFIG_DIR'"
            exit 1
        fi
    fi
    if [ ! -w "$DOCKER_TOOLS_CONFIG_DIR" ]; then
        >&2 echo "Configuration directory is not writable '$DOCKER_TOOLS_CONFIG_DIR'"
        exit 2
    fi

    # Make sure the full directory structure exists
    for dir in "${__docker_tools_conf_dirs__[@]}"; do
        if [ ! -d "$DOCKER_TOOLS_CONFIG_DIR/$dir" ]; then
            mkdir -pv "$DOCKER_TOOLS_CONFIG_DIR/$dir"
            exit_code=$?
            if [ 0 -ne $exit_code ]; then
                >&2 echo "Could not create configuration directory '$DOCKER_TOOLS_CONFIG_DIR/$dir'"
                exit 3
            fi
        fi
        if [ ! -w "$DOCKER_TOOLS_CONFIG_DIR/$dir" ]; then
                >&2 echo "Configuration directory is not writable '$DOCKER_TOOLS_CONFIG_DIR/$dir'"
                exit 4
        fi
    done

    # Make sure all required libraries exist
    for file in "${__docker_tools_conf_files__[@]}"; do
        if [ ! -f "$DOCKER_TOOLS_CONFIG_DIR/$file" ]; then
            curl -f -L -s "${__DOCKER_TOOLS_CONFIG_URL__}/$file" > "$DOCKER_TOOLS_CONFIG_DIR/$file"
            exit_code=$?
            if [ 0 -ne $exit_code ]; then
                >&2 echo "Could not create configuration file '$DOCKER_TOOLS_CONFIG_DIR/$file'"
                exit 5
            fi
        fi
        if [ ! -w "$DOCKER_TOOLS_CONFIG_DIR/$file" ]; then
            >&2 echo "Configuration file is not writable '$DOCKER_TOOLS_CONFIG_DIR/$file'"
            exit 6
        fi
    done
}

__init
exit_code=$?
if [ 0 -ne $exit_code ]; then
    >&2 echo "Initialization failed: $exit_code"
    exit $exit_code
fi

##############################################################################
# Import all libraries and supporting files and execute
##############################################################################

# Application configuration & globals
source "$DOCKER_TOOLS_CONFIG_DIR/config"
source "$DOCKER_TOOLS_CONFIG_DIR/lib/globals.sh"

# Import core libraries
source "$DOCKER_TOOLS_LIB_DIR/lib-ui.sh"
source "$DOCKER_TOOLS_LIB_DIR/lib-cli.sh"

# Import supporting libraries
source "$DOCKER_TOOLS_LIB_DIR/lib-config.sh"
source "$DOCKER_TOOLS_LIB_DIR/lib-recipes.sh"
source "$DOCKER_TOOLS_LIB_DIR/lib-registry.sh"

# Import docker-tools command libraries
source "$DOCKER_TOOLS_LIB_DIR/lib-tools.sh"
source "$DOCKER_TOOLS_LIB_DIR/lib-tools-config.sh"
source "$DOCKER_TOOLS_LIB_DIR/lib-tools-list.sh"
source "$DOCKER_TOOLS_LIB_DIR/lib-tools-install.sh"
source "$DOCKER_TOOLS_LIB_DIR/lib-tools-main.sh"

# Execute
__main $@
exit $?
