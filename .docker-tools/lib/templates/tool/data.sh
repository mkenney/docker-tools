
# Tool script template data

function __command__ {
    printf "${__RECIPE_CMD__}"
}

function __entrypoint__ {
    local ret_val
    if [ "" != "${__RECIPE_ENTRYPOINT__}" ]; then
        ret_val="--entrypoint=\""${__RECIPE_ENTRYPOINT__}"\""
    fi
    printf "$ret_val"
}

function __image__ {
    if [ "" == "${__RECIPE_IMAGE__}" ]; then
        echo "Invalid docker image '$${__RECIPE_IMAGE__}', exiting"
        exit 1
    fi
    printf "${__RECIPE_IMAGE__}"
}

function __prefix__ {
    printf "${__RECIPE_PREFIX__}"
}

function __recipe__ {
    printf "${__RECIPE_RECIPE__}"
}

function __tag__ {
    if [ "" == "${__RECIPE_IMAGE_TAG__}" ]; then
        echo "Invalid image tag '$${__RECIPE_IMAGE_TAG__}', exiting"
        exit 1
    fi
    printf "${__RECIPE_IMAGE_TAG__}"
}

function __term__ {
    local ret_val
    if [ -t 0 ]; then
        ret_val="-t"
    fi
    printf -- "$ret_val"
}

function __update_image__ {
    docker pull $(__image__):$(__tag__)
}

function __update_script__ {
    printf "@todo write the __update_script__ function... ;-)\n"
}

function __volumes__ {
    local -a ret_val
    local volume
    local a=0
    while read volume; do
        if [ "" != "$volume" ]; then
            ret_val[$a]="-v $volume"
            a=$((a + 1))
        fi
    done <<< "$(echo "${__RECIPE_VOLUMES__}" | tr ";" "\n")"
    printf "%s " "${ret_val[@]}"
}

