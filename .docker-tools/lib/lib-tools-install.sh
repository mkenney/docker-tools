
##############################################################################
##############################################################################
##
##  Functions to manage the docker-tools "install" command
##
##  Also includes documentation
##
##############################################################################
##############################################################################


#########################
#
#  install
#
#########################

function __install_usage {
        echo "
 Usage: $(_s b)docker-tools install$(_s r) <recipe name> [options]...
    Or: $(_s b)docker-tools install$(_s r) --image=IMAGE_NAME --name=TOOL_NAME [options]...

    <recipe name>       Optional. If included, the specified recipe will be
                        installed. Any recipe value can be overridden with
                        an option

    --cmd               Specify a prefix to any tool arguments
    --entrypoint        Specify an entrypoint for the Docker container
    --image             Required if a recipe isn't specified. Specify an image
                        for the Docker container
    --tag               Specify an image tag for the Docker container, default
                        'latest'
    --name              Required if a recipe isn't specified. Define the tool
                        name. This will be the name of the installed script.
    --prefix            Specify the install location
    --volumes           Specify any volumes that should be mounted. This is a
                        semicolon delimited list of \`docker run\` volume
                        mount strings
"
}

function __install_help {
    echo "
$(_s b)NAME$(_s r)

    \`$(_s b)docker-tools install$(_s r)\` -- Install a tool

$(_s b)USAGE$(_s r)

    docker-tools install [RECIPE_NAME] [options]

$(_s b)DESCRIPTION$(_s r)

    Install a tool on the local system. A tool can be constructed from a defined recipe or directly
    using command options. If a recipe name is not provided then $(_s i)--image$(_s r) and $(_s i)--name$(_s r) options are
    required, otherwise all options are... optional.

$(_s b)OPTIONS$(_s r)

    $(_s b)--cmd=$(_s r u)cmd$(_s r)
        Specify a command prefix. This command or string will be prefixed to any arguments passed
        to the tool being defined. This is useful  for specifying a command in containers that use
        a $(_s u)cmd$(_s r) instead of an $(_s u)entrypoint$(_s r) or for tailoring your tool for a specific purpose. See the
        $(_s u)cmd$(_s r) reference for more information.

        See https://docs.docker.com/engine/reference/run/#/cmd-default-command-or-options

    $(_s b)--entrypoint=$(_s r u)entrypoint$(_s r)
        Specify an entrypoint for the container executed by the tool. See the $(_s u)entrypoint$(_s r) reference
        for more information.

        https://docs.docker.com/engine/reference/run/#/entrypoint-default-command-to-execute-at-runtime

    $(_s b)--image=$(_s r u)image$(_s r)
        Specify the Docker $(_s u)image$(_s r) to use. If a URL is not provided then the Docker HUB repository is
        used by default.

        https://docs.docker.com/docker-hub/repos/

    $(_s b)--tag=$(_s r u)tag$(_s r)
        Specify the image $(_s u)tag$(_s r) use in the container executed by the tool. Default 'latest'.

    $(_s b)--name=$(_s r u)name$(_s r)
        Specify the tool $(_s u)name$(_s r), required if a recipe is not provided. This is the name to call the
        tool that's being installed (\`npm\`, \`php\`, etc.)

    $(_s b)--prefix=$(_s r u)path$(_s r)
        Default '$(_s u)/usr/local/bin$(_s r)'. Specify the location to install the tool. The default location
        can be overridden by defining DOCKER_TOOLS_PREFIX in your .bashrc or similar.

        \`export DOCKER_TOOLS_PREFIX=$(_s u)/default/install/path$(_s r)\`

    $(_s b)--save_as=$(_s r u)name$(_s r)
        Save this tool install statement as a recipe. If saved, this tool can be installed just by
        specifying the recipe $(_s u)name$(_s r), and your recipe file
        ($(_s i)\$DOCKER_TOOLS_CONFIG_DIR/recipies$(_s r)) can be easily copied or shared.

    $(_s b)--volumes=$(_s r u)volumes$(_s r)
        Specify any $(_s u)volumes$(_s r) to be mounted into the container run by the tool. This is a
        $(_s i)semicolon$(_s r) separated list of \`docker run\` $(_s u)volume$(_s r) mount strings.

        \`--volumes=[host-src:]container-dest[:<options>][;[host-src:]container-dest[:<options>]...]\`

        https://docs.docker.com/engine/reference/run/#volume-shared-filesystems

$(_s b)EXAMPLES$(_s r)

    Install a tool from a stored recipe but change it's name and command:
        $ $(_s i)docker-tools install my-npm-recipe --name=gulp --command=gulp$(_s r)

    Install a tool from scratch:
        $ $(_s i)docker-tools install \\
          --name=gulp \\
          --prefix=/usr/local/bin \\
          --image=mkenney/npm \\
          --tag=7.0-alpine \\
          --entrypoint='/usr/local/bin/gulp' \\
          --volumes='\\\$(pwd):/src:rw$(_s r)'

    Save it in your recipes:
          $(_s i)--save_as=my-phpunit-recipe$(_s r)"
}

#
# install
#
function __docker_tools_install {
    local old_IFS=IFS

    ########################
    # 2 modes
    #   if no recipe is specified, then at minimum --image and --name options
    #   are required
    #
    #   if a recipe is specified, no options are required, but any may be
    #   passed as a 1-time override of the corresponding recipe value
    #########################

    local recipe_name="$(__get_arg 1 $@)"
    local opt_name="$(__get_opt name $@)"
    local opt_prefix="$(__get_opt prefix $@)"
    local opt_type="$(__get_opt type $@)"
    local opt_image="$(__get_opt image $@)"
    local opt_tag="$(__get_opt tag $@)"
    local opt_entrypoint="$(__get_opt entrypoint $@)"
    local opt_cmd="$(__get_opt cmd $@)"
    local opt_volumes="$(__get_opt volumes $@)"

    if [ "" == "$recipe_name" ]; then
        if [ "" == "$opt_image" ] || [ "" == "$opt_name" ]; then
            __show_usage install " Error: A recipe name must be specified or --image AND --name options are
 required"
            exit 1
        fi
    fi
 Error: A recipe name must be specified or --image AND --name options are
 required
    # Tool recipe
    local -a recipe
    IFS=$"$__recipe_delimiter__"
    recipe=($(__recipe_get $recipe_name))
    IFS=$old_IFS
    if [ "" != "$opt_cmd" ];        then recipe[6]="$opt_cmd";        fi
    if [ "" != "$opt_entrypoint" ]; then recipe[5]="$opt_entrypoint"; fi
    if [ "" != "$opt_image" ];      then recipe[3]="$opt_image";      fi
    if [ "" != "$opt_tag" ];        then recipe[4]="$opt_tag";        fi
    if [ "" != "$opt_name" ];       then recipe[1]="$opt_name";       fi
    if [ "" != "$opt_prefix" ];     then recipe[2]="$opt_prefix";     fi
    if [ "" != "$opt_type" ];       then recipe[3]="$opt_prefix";     fi
    if [ "" != "$opt_volumes" ];    then recipe[7]="$opt_volumes";    fi

    if [ "" == "$opt_name" ]; then
        echo "Error - a tool name was not provided"
        exit 1
    fi


    # Tool tempfile
    cp $DOCKER_TOOLS_LIB_DIR/templates/tool.sh $__INSTALL_TMPFILE__

        # tool info
    sed -i "s/declare __TOOLS_VERSION__=/declare __TOOLS_VERSION__=$DOCKER_TOOLS_VERSION/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_NAME__=/declare __RECIPE_NAME__=$recipe_name/" $__INSTALL_TMPFILE__
    sed -i "s/declare __TOOLS_LIB_DIR__=/declare __TOOLS_LIB_DIR__=$DOCKER_TOOLS_LIB_DIR/" $__INSTALL_TMPFILE__

        # tool recipe
    sed -i "s/declare __RECIPE_CMD__=/declare __RECIPE_CMD__=$opt_cmd/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_ENTRYPOINT__=/declare __RECIPE_ENTRYPOINT__=$opt_entrypoint/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_IMAGE__=/declare __RECIPE_IMAGE__=$opt_image/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_IMAGE_TAG__=/declare __RECIPE_IMAGE_TAG__=$opt_tag/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_TOOL__=/declare __RECIPE_TOOL__=$opt_name/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_PREFIX__=/declare __RECIPE_PREFIX__=$opt_prefix/" $__INSTALL_TMPFILE__
    sed -i "s/declare __RECIPE_VOLUMES__=/declare __RECIPE_VOLUMES__=$opt_volumes/" $__INSTALL_TMPFILE__

    if [ "" == "$opt_prefix" ]; then opt_prefix=$DOCKER_TOOLS_PREFIX; fi
    #result=$( (cat ${__INSTALL_TMPFILE__} > $opt_prefix/$opt_name) 2>${__ERROR_TMPFILE__})
    #exit_code=$?
    #errors=$(< ${__ERROR_TMPFILE__})


echo "result:$result"
echo "exit_code:$exit_code"
echo "errors:$errors"
#sed -i "s/php_value newrelic.appname emt_web-dev/php_value newrelic.appname $APP_NAME/" /var/www/html/documentroot/.htaccess


echo "__docker_tools_install:recipe:${recipe[@]}"




}

