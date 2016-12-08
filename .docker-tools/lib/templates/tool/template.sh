
# Tool script template

if [ -f "${__TEMPLATE_DIR__}/libs.sh" ]; then source "${__TEMPLATE_DIR__}/libs.sh"; fi
if [ -f "${__TEMPLATE_DIR__}/data.sh" ]; then source "${__TEMPLATE_DIR__}/data.sh"; fi

if [ "self-update" = "$1" ]; then
    __update_image__  # docker pull $DOCKER_TOOL_IMAGE:$DOCKER_TOOL_TAG
    __update_script__ # curl -f -L -s $INSTALL_SCRIPT | sh -s $SCRIPT $TAG $(dirname $0) && exit 0

else
    docker run --rm -i $(__term__) $(__volumes__) $(__entrypoint__) $(__image__):$(__tag__) $(__command__) $@
fi
