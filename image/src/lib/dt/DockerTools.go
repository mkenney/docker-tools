/*
Package dt defines the DockerTools instance and associated methods for managing
the `docker-tools` runtime

This portion of the dt package contains methods for initializing and executing
the `docker-tools` primary command and it's sub-commands
*/
package dt

import (
	"flag"
	"fmt"
	"lib/cli"
	"lib/config"
	"lib/recipes"
	"os"
	"os/exec"

	"github.com/golang/glog"
)

/*
DefaultPrefix is the default install path for generated scripts
*/
const DefaultPrefix = "/usr/local/bin"

/*
DockerTools represents the running docker-tools program
*/
type DockerTools struct {

	/*
		the name of the command (docker-tools)
	*/
	Name string

	/*
		A key/value map of command-line options passed to docker-tools
	*/
	Opts map[string]string

	/*
		A list of command-line flags passed to docker-tools
	*/
	Flags map[string]bool

	/*
		All commands for use internally. Should generally only be 1 primary
		command and sometimes a secondary command modifier (the name of a
		recipe or tool for example)
	*/
	Commands cli.CommandList

	/*
		All available recipes
	*/
	Recipes *recipes.Recipes
}

/*
New initialize a new DockerTools instance and return a pointer
*/
func New() *DockerTools {

	// Init docker-tools
	dockerTools := new(DockerTools)
	// Shift the cli options and flags for the docker-tools command off the
	// stack and store them locally
	var command cli.Command
	var err error
	cli.Commands, command, err = cli.Commands.Shift()
	if nil != err {
		glog.Fatalf("Error shifting commands")
	}
	dockerTools.Commands = cli.Commands
	dockerTools.Name = command.Name
	dockerTools.Opts = command.Opts
	dockerTools.Flags = command.Flags

	// Load all recipe files
	dockerTools.Recipes = recipes.New()
	dockerTools.Recipes = dockerTools.Recipes.Load(config.Values.ConfPath + "/registry.yml")
	dockerTools.Recipes = dockerTools.Recipes.Load(config.Values.ConfPath + "/recipes.yml")

	return dockerTools
}

/*
Run executes the docker-tools program
*/
func (dt *DockerTools) Run() {
	//fmt.Printf("Commands: %v", dt.Commands)
	//os.Exit(0)
	var err error
	if 0 < len(dt.Commands) {
		switch {
		//        case "config" == dt.Commands[0].Name:
		//            dt.ConfigureDockerTools()
		//
		//        case "self-update":
		//            dt.Update()
		//
		//        case "create":
		//            dt.CreateRecipe()
		//
		//        case "edit":
		//            dt.EditRecipe()

		case "generate" == dt.Commands[0].Name:
			if nil == err {
				dt.Generate()
			}

		case "list" == dt.Commands[0].Name:
			if nil == err {
				dt.ListRecipes()
			}

			//        case "delete":
			//            dt.DeleteRecipe()
			//
			//        case "install":
			//            dt.InstallTool()
			//
			//        case "uninstall":
			//            dt.UninstallTool()
			//
			//        case "update":
			//            dt.UpdateTool()
		default:
			fmt.Printf("command.Name: %v", dt.Commands[0].Name)
		}
	} else {
		dt.usage("dt")
	}
}

func (dt *DockerTools) usage(command string) {

}

/*
Generate usage:
	docker-tools generate RECIPE_SOURCE RECIPE_NAME [options]
*/
func (dt *DockerTools) Generate() {
	var recipe *recipes.Recipe

	commands, opts, err := dt.Commands.Shift()
	if nil != err {
		glog.Fatalf("Cannot generate script. Not really sure how I got here: %s", err)
	}

	if 0 == len(commands) {
		recipe = recipes.NewRecipe([]string{})

	} else {
		if "recipes" != commands[0].Name && "registry" != commands[0].Name {
			glog.Fatalf("Unknown recipe source '%s'", commands[0].Name)
		}
		if !dt.Recipes.HasRecipe(commands[1].Name, commands[0].Name) {
			glog.Fatalf("Unknown recipe '%s'", commands[1].Name)
		}
		recipe = dt.Recipes.GetRecipe(commands[1].Name, commands[0].Name)

		//recipe.SetCliVolumes(recipeData[6])
		//recipe.SetCliEnv(recipeData[7])
	}

	if opts.HasOpt("name")       {recipe.ToolName = opts.Opts["name"]}
	if opts.HasOpt("prefix")     {recipe.Prefix = opts.Opts["prefix"]}
	if opts.HasOpt("template")   {recipe.Template = opts.Opts["template"]}
	if opts.HasOpt("image")      {recipe.Image = opts.Opts["image"]}
	if opts.HasOpt("tag")        {recipe.Tag = opts.Opts["tag"]}
	if opts.HasOpt("volumes")    {recipe.SetVolumes(opts.Opts["volumes"])}
	if opts.HasOpt("env")        {recipe.SetEnv(opts.Opts["env"])}
	if opts.HasOpt("entrypoint") {recipe.Entrypoint = opts.Opts["entrypoint"]}
	if opts.HasOpt("cmd")        {recipe.Cmd = opts.Opts["cmd"]}
	//if opts.HasOpt("options")    {recipe.SetOptions(opts.Opts["options"])}


	fmt.Printf("Recipe: %s", recipe.ToString())
}

/*
ListRecipes iterates through all known recipies and passes their data to a
formatting method which returns CLI-formatted output. It then passes that output
through a pager (`less -r`).
*/
func (dt *DockerTools) ListRecipes() {
	var listing string

	for _, recipe := range (*dt.Recipes) {
		if nil == recipe {
			break
		}
		listing += recipe.Render() + "\n"
	}

	pipestdin, pipestdout, err := os.Pipe()
	if err != nil {
		panic("Could not create pipe")
	}

	stdout := os.Stdout
	os.Stdout = pipestdout

	pager := exec.Command("less", "-r")
	pager.Stdin = pipestdin
	pager.Stdout = stdout // the pager uses the original stdout, not the pipe...
	pager.Stderr = os.Stderr

	defer func() {
		pipestdout.Close()
		err := pager.Run()
		os.Stdout = stdout
		if err != nil {
			glog.Fatalf("%v", os.Stderr)
			glog.Fatalf("%s", err)
		}
	}()

	fmt.Println("\n\n" + listing)
}

func init() {
	flag.Parse() // Required for glog
}