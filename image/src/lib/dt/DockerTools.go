/*
Package dt defines the DockerTools instance and associated methods for managing
the `docker-tools` runtime

This portion of the dt package contains methods for initializing and executing
the `docker-tools` primary command and it's sub-commands
*/
package dt

import (
	"fmt"
	"lib/cli"
	"lib/config"
	"lib/recipes"
	"lib/templates/docs"

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
	Command cli.Command

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
	var err error
	dockerTools.Commands, dockerTools.Command, err = cli.Commands.Shift()
	if nil != err {
		glog.Fatalf("Error shifting commands")
	}

	// Load all recipe files
	dockerTools.Recipes = recipes.New()
	dockerTools.Recipes = dockerTools.Recipes.Load(config.DockerToolsRegistry)
	dockerTools.Recipes = dockerTools.Recipes.Load(config.ConfPath + "/recipes.yml")

	return dockerTools
}

/*
Run executes the docker-tools program
*/
func (dt *DockerTools) Run() {
	var command cli.Command

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
			dt.Commands, command, _ = dt.Commands.Shift()

			if 0 < len(command.Opts["help"]) {
				docs.GenerateHelp()
			} else if command.HasFlag("h") {
				docs.GenerateUsage()
			} else {
				dt.GenerateScript()
			}

		case "list" == dt.Commands[0].Name:
			dt.ListRecipes()

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
			docs.DockerToolsHelp()
		}
	} else {
		if 0 < len(dt.Command.Opts["help"]) {
			docs.DockerToolsHelp()
		} else {
			docs.DockerToolsUsage()
		}
	}
}

func (dt *DockerTools) usage(command string) {

}

/*
Generate usage:
	docker-tools generate RECIPE_SOURCE RECIPE_NAME [options]
*/
func (dt *DockerTools) GenerateScript() {
	var recipe *recipes.Recipe
	var recipeSource string

	commands, recipeCommand, err := dt.Commands.Shift()
	if nil != err {
		glog.Fatalf("Cannot generate script. Not really sure how I got here: %s", err)
	}

	if 0 == len(commands) {
		recipeSource = "registry"

	} else {
		if "recipes" != commands[0].Name && "registry" != commands[0].Name {
			glog.Fatalf("Unknown recipe source '%s'", commands[0].Name)
		}
		recipeSource = commands[0].Name
	}

	recipe = recipes.NewRecipe()
	if dt.Recipes.HasRecipe(recipeCommand.Name, recipeSource) {
		recipe = dt.Recipes.GetRecipe(recipeCommand.Name, recipeSource)
	}

	if recipeCommand.HasOpt("name")       {recipe.ToolName = recipeCommand.Opts["name"][0]}
	if recipeCommand.HasOpt("prefix")     {recipe.Prefix = recipeCommand.Opts["prefix"][0]}
	if recipeCommand.HasOpt("template")   {recipe.Template = recipeCommand.Opts["template"][0]}
	if recipeCommand.HasOpt("image")      {recipe.Image = recipeCommand.Opts["image"][0]}
	if recipeCommand.HasOpt("tag")        {recipe.Tag = recipeCommand.Opts["tag"][0]}
	if recipeCommand.HasOpt("volumes")    {
		for _, vol := range recipeCommand.Opts["volumes"] {
			recipe.AddVolume(vol)
		}
	}
	if recipeCommand.HasOpt("env")        {
		for _, env := range recipeCommand.Opts["env"] {
			recipe.AddEnv(env)
		}
	}
	if recipeCommand.HasOpt("entrypoint") {recipe.Entrypoint = recipeCommand.Opts["entrypoint"][0]}
	if recipeCommand.HasOpt("cmd")        {recipe.Cmd = recipeCommand.Opts["cmd"][0]}
	if recipeCommand.HasOpt("options")    {
		for _, opt := range recipeCommand.Opts["options"] {
			recipe.AddOption(opt)
		}
	}

	fmt.Printf("%s", recipe.ToString())
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

	cli.Page(listing)
}
