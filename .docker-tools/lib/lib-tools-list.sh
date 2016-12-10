
##############################################################################
##############################################################################
##
##  Functions to manage the docker-tools "list" command
##
##  Also includes documentation
##
##############################################################################
##############################################################################

#########################
#
#  documentation
#
#########################

function __list_usage {
        echo "
usage: docker-tools list ...
        "
}

function __list_help {
    echo "
    $(_s u)Name:$(_s r)

        \`$(_s b)docker-tools list$(_s r)\` -- list available recipes

    $(_s u)Usage:$(_s r)

        docker-tools list

    $(_s u)Description:$(_s r)

        List available recipes.

    $(_s u)Options:$(_s r)
        --source

    $(_s u)Examples:$(_s r)

        $ docker-tools list
"
}


#
# Define the 'docker-tools list' command interface
#
# @option --source
#
function __docker_tools_list {
    local ret_val
    local -a parts

    local -a recipe_files=("$DOCKER_TOOLS_RECIPES" "$DOCKER_TOOLS_REGISTRY")
    local recipe_file
    local recipe_line
    local recipe_name
    local -a recipes
    local recipe
    local a=0

    for recipe_file in "${recipe_files[@]}"; do
        while read recipe_line; do
            if [ "" != "${recipe_line:0:1}" ] && [ "#" != "${recipe_line:0:1}" ]; then
                ret_val="$ret_val $(__recipe_describe "${recipe_line}|$(basename ${recipe_file})" )"
            fi
        done < $recipe_file
    done

    printf "$ret_val"
}
