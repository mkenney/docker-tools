#!/usr/bin/env bash

##############################################################################
##############################################################################
##
##  This is the bootstrap file for a `docker run` wrapper script generated
##  using the `docker-tools` command and any modifications made here may be
##  overwritten without notice.
##
##  See `docker-tools --help` for more information.
##
##############################################################################
##############################################################################

# tool info
declare __TOOLS_VERSION__=
declare __RECIPE_NAME__=
declare __TOOLS_LIB_DIR__=

# tool recipe
declare __RECIPE_CMD__=
declare __RECIPE_ENTRYPOINT__=
declare __RECIPE_IMAGE__=
declare __RECIPE_IMAGE_TAG__=
declare __RECIPE_TOOL__=
declare __RECIPE_PREFIX__=
declare __RECIPE_VOLUMES__=

# tool resources
source "${__TOOLS_LIB_DIR__}/globals.sh"
source "${__TOOLS_LIB_DIR__}/lib-cli.sh"
source "${__TOOLS_LIB_DIR__}/lib-ui.sh"

##############################################################################
##############################################################################
##
##  Execute
##
##############################################################################
##############################################################################

# execute the tool
declare __TEMPLATE_DIR__="${__TOOLS_LIB_DIR__}/templates/tool"
if [ ! -f "${__TEMPLATE_DIR__}/template.sh" ]; then
    >&2 echo "Template not found: '${__TEMPLATE_DIR__}/template.sh'"
    exit 1
fi
source "${__TEMPLATE_DIR__}/template.sh"

# all errors should be "caught"... and sourced scripts should exit 0
declare exit_code=$?
if [ 0 -ne "$exit_code" ]; then
    echo "An unknown error occurred"
fi

exit $exit_code
