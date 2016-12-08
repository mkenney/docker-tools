
#
# Update all libraries and repos
#
function __update {
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

    # Download all managed files
    for file in "${__docker_tools_conf_files__[@]}"; do
        if [ ! -w "$DOCKER_TOOLS_CONFIG_DIR/$file" ]; then
            >&2 echo "Configuration file is not writable '$DOCKER_TOOLS_CONFIG_DIR/$file'"
            exit 5
        fi
        curl -f -L -s "${__DOCKER_TOOLS_CONFIG_URL__}/$file" > "$DOCKER_TOOLS_CONFIG_DIR/$file"
        exit_code=$?
        if [ 0 -ne $exit_code ]; then
            >&2 echo "Could not update '$DOCKER_TOOLS_CONFIG_DIR/$file'"
            exit 6
        fi
    done
}
