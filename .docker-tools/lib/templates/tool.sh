#!/usr/bin/env bash

##############################################################################
##############################################################################
##
##  This is the bootstrap file for a `docker run` wrapper script generated
##  using the `docker-tools` command. Any modifications made here may be
##  overwritten without notice.
##
##  See `docker-tools --help` for more information.
##
##############################################################################
##############################################################################

# docker-tools metadata
declare __TOOLS_VERSION__=
declare __TOOLS_LIB_DIR__=

# tool metadata
declare __TOOL_NAME__=
declare __TOOL_PREFIX__=
declare __TOOL_TEMPLATE__=

# recipe metadata
declare __RECIPE_NAME__=asdf
declare __RECIPE_IMAGE__=
declare __RECIPE_IMAGE_TAG__=
declare __RECIPE_ENTRYPOINT__=
declare __RECIPE_CMD__=
declare __RECIPE_VOLUMES__=
declare __RECIPE_OPTIONS__=
declare __RECIPE_NOTE__=

##############################################################################
##############################################################################
##
##  Execute
##
##############################################################################
##############################################################################

# import required resources
if [ ! -f "${__TOOLS_LIB_DIR__}/globals.sh" ]; then >&2 printf "Required resource missing: '${__TOOLS_LIB_DIR__}/globals.sh'\n"; exit 1; fi
if [ ! -f "${__TOOLS_LIB_DIR__}/lib-cli.sh" ]; then >&2 printf "Required resource missing: '${__TOOLS_LIB_DIR__}/lib-cli.sh'\n"; exit 2; fi
if [ ! -f "${__TOOLS_LIB_DIR__}/lib-ui.sh" ];  then >&2 printf "Required resource missing: '${__TOOLS_LIB_DIR__}/lib-ui.sh'\n";  exit 3; fi
source "${__TOOLS_LIB_DIR__}/globals.sh"
source "${__TOOLS_LIB_DIR__}/lib-cli.sh"
source "${__TOOLS_LIB_DIR__}/lib-ui.sh"

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
