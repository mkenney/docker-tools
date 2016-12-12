package dt
import (
    "fmt"
    "lib/cli"
    "lib/config"
    "lib/recipes"
    "lib/ui"
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
     * Pointer to the application configuration instance
     */
    Config *config.Config

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
    dtools.Config = config.New()
    dtools.Name = command.Name
    dtools.Opts = command.Opts
    dtools.Flags = command.Flags

    // Init all recipe files
    dtools.Recipes = recipes.New()
    ret_chan := make(chan bool)
    err_chan := make(chan error)
    go dtools.Recipes.Load(dtools.Config.Path+"/registry", ret_chan, err_chan)
    go dtools.Recipes.Load(dtools.Config.Path+"/recipes", ret_chan, err_chan)
    if _, _, erra, errb := <-ret_chan, <-ret_chan, <-err_chan, <-err_chan; nil != erra {
        log.Output(1, fmt.Sprintf("Error loading recipes from file '%v'", erra))
    } else if nil != errb {
        log.Output(1, fmt.Sprintf("Error loading recipes from file '%v'", errb))
    }

    return dtools
}

/**
 * List current recipes to stdout
 * @param  {[type]} this *DockerTools) ListRecipes( [description]
 * @return {[type]}      [description]
 */
func (this *DockerTools) ListRecipes() {
    for _, recipe := range this.Recipes {
        if nil == recipe {break}
        fmt.Printf("RecipeName: %v\n", recipe.RecipeName)
    }
}

func (this *DockerTools) Run() {

    ui.Test()

    var err error
    switch {
//        case "config":
//            this.Config()
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

