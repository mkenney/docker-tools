package dt

import (
	"fmt"
	"lib/cli"
	"lib/config"
	"lib/recipes"
	"log"
	"os"
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
	dtools := new(DockerTools)
	// Shift the cli options and flags for the docker-tools command off the
	// stack and store them locally
	var command cli.Command
	var err error
	cli.Commands, command, err = cli.Commands.Shift()
	if nil != err {
		log.Fatalf("Error shifting commands")
	}
	dtools.Commands = cli.Commands
	dtools.Name = command.Name
	dtools.Opts = command.Opts
	dtools.Flags = command.Flags

	// Init all recipe files
	dtools.Recipes = recipes.NewRecipes()
	erra := dtools.Recipes.Load(config.Config.ConfPath + "/registry")
	errb := dtools.Recipes.Load(config.Config.ConfPath + "/recipes")
	if nil != erra {
		log.Output(1, fmt.Sprintf("Error loading recipes from file '%v'", erra))
	} else if nil != errb {
		log.Output(1, fmt.Sprintf("Error loading recipes from file '%v'", errb))
	}

	return dtools
}

/*
Run executes the docker-tools program
*/
func (dt *DockerTools) Run() {
	fmt.Printf("Commands: %v", dt.Commands)
	os.Exit(0)
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
