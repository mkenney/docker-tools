
package docs

import (
    "bytes"
    "fmt"
    "lib/cli"
    "lib/ui"

	"github.com/golang/glog"
)

/*
DockerToolsHelp will render and display full documentation for the docker-tools commands
*/
func DockerToolsHelp() {
    var retBuffer bytes.Buffer

    err := compileHelpTemplate(&retBuffer, "dockerToolsHelp", getDockerToolsTemplateData())
    if nil != err {glog.Fatalf("Failed to execute template '_head': %s", err)}

	cli.Page(retBuffer.String())
}

/*
DockerToolsUsage will render simple command usage
*/
func DockerToolsUsage() {
    var retBuffer bytes.Buffer

    err := compileUsageTemplate(&retBuffer, "dockerToolsUsage", getDockerToolsTemplateData())
    if nil != err {glog.Fatalf("Failed to execute template '_head': %s", err)}

	fmt.Println(retBuffer.String())
}

func getDockerToolsTemplateData() (retval map[string]string) {
    retval = ui.GetTemplateVars()

    // Document labels
    retval["labelDOCKER_TOOLS_CONFIG_DIR"]     = ui.B("DOCKER_TOOLS_CONFIG_DIR")
    retval["labelDOCKER_TOOLS_PREFIX"]         = ui.B("DOCKER_TOOLS_PREFIX")

    // Document highlight words
    retval["highlightConfig"]                  = ui.WhiteBt("config")
    retval["highlightSelfUpdate"]              = ui.WhiteBt("self-update")
    retval["highlightCreate"]                  = ui.WhiteBt("create")
    retval["highlightList"]                    = ui.WhiteBt("list")
    retval["highlightDelete"]                  = ui.WhiteBt("delete")
    retval["highlightInstall"]                 = ui.WhiteBt("install")
    retval["highlightUninstall"]               = ui.WhiteBt("uninstall")
    retval["highlightUpdate"]                  = ui.WhiteBt("update")
    retval["highlightDOCKER_TOOLS_CONFIG_DIR"] = ui.WhiteBt("DOCKER_TOOLS_CONFIG_DIR")
    retval["highlightDOCKER_TOOLS_PREFIX"]     = ui.WhiteBt("DOCKER_TOOLS_PREFIX")

    // Command examples
    retval["exampleList"]                      = ui.I("docker-tools "+ui.U(ui.I("list")))
    retval["exampleListRegistry"]              = ui.I("docker-tools "+ui.U(ui.I("list"))+" "+ui.I("--source=registry"))
    retval["exampleListRecipes"]               = ui.I("docker-tools "+ui.U(ui.I("list"))+" "+ui.I("--source=recipes"))
    retval["exampleInstallGulp"]               = ui.I("docker-tools "+ui.U(ui.I("install"))+" "+ui.I("gulp"))
    retval["exampleInstallGulpRegistry"]       = ui.I("docker-tools "+ui.U(ui.I("install"))+" "+ui.I("gulp --source=registry"))
    retval["exampleGenerateRegistryRecipe"]    = ui.I("docker-tools "+ui.U(ui.I("generate"))+" "+ui.I("some-registered-recipe --source=registry"))
    retval["exampleGenerateCustomRecipe"]      = ui.I("docker-tools "+ui.U(ui.I("generate"))+" "+ui.I("my-custom-recipe  --source=recipes"))
    retval["exampleGenerateRecipeCustomTag"]   = ui.I("docker-tools "+ui.U(ui.I("generate"))+" "+ui.I("registry some-registered-recipe --tag=different-image-tag"))
    retval["exampleInstallGulpOverrideTag"]    = ui.I("docker-tools "+ui.U(ui.I("install"))+" "+ui.I("gulp --tag=7.0-alpine"))

    retval["exampleMoreHelp"]                  = ui.I("docker-tools "+ui.U("COMMAND")+" --help")
    return
}
