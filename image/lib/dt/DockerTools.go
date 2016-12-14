package dt
import (
    "fmt"
    "lib/cli"
    "lib/config"
    "lib/recipes"
//    "lib/ui"
    "log"
//    "os"
)

const DefaultPrefix = "/usr/local/bin"
const default_prefix = "/usr/local/bin"

/**
 * Represents the docker-tools program controller and it's options / flags
 */
type DockerTools struct {

    /**
     * The name of the command (docker-tools)
     */
    Name string

    /**
     * A key/value map of command-line options
     */
    Opts map[string]string

    /**
     * A list of command-line flags
     */
    Flags map[string]bool

    /**
     * All commands for use internally. Should generally only be 1 primary
     * command and sometimes a secondary command modifier (the name of a
     * recipe or tool for example)
     */
    Commands *cli.CliCommands

    /**
     * All available recipes
     */
    Recipes *recipes.Recipes
}

/**
 * Initialize a new DockerTools instance and return a pointer
 *
 * Defines the
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
    erra := dtools.Recipes.Load(config.Config.ConfPath+"/registry")
    errb := dtools.Recipes.Load(config.Config.ConfPath+"/recipes")
    if nil != erra {
        log.Output(1, fmt.Sprintf("Error loading recipes from file '%v'", erra))
    } else if nil != errb {
        log.Output(1, fmt.Sprintf("Error loading recipes from file '%v'", errb))
    }

    return dtools
}

/**
 * [func description]
 * @param  {[type]} this *DockerTools) Run( [description]
 * @return {[type]}      [description]
 */
func (this *DockerTools) Run() {

    var err error
    switch {
//        case "config" == this.Commands[0].Name:
//            this.ConfigureDockerTools()
//
//        case "self-update":
//            this.Update()
//
//        case "create":
//            this.CreateRecipe()
//
//        case "edit":
//            this.EditRecipe()

        case "list" == this.Commands[0].Name:
//ui.Test()
            if nil == err {
                this.ListRecipes()
            }

//        case "delete":
//            this.DeleteRecipe()
//
//        case "install":
//            this.InstallTool()
//
//        case "uninstall":
//            this.UninstallTool()
//
//        case "update":
//            this.UpdateTool()
        default:
            fmt.Printf("command.Name: %v", this.Commands[0].Name)
    }

}
