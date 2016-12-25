
package docs

import (
    "bytes"
    "fmt"
    "lib/cli"
    "lib/ui"

	"github.com/golang/glog"
)

/*
GenerateHelp will render and display full documentation for the docker-tools commands
*/
func GenerateHelp() {
    var retBuffer bytes.Buffer

    err := compileHelpTemplate(&retBuffer, "generateHelp", getGenerateTemplateData())
    if nil != err {glog.Fatalf("Failed to execute template '_head': %s", err)}

	cli.Page(retBuffer.String())
}

/*
GenerateUsage will render simple command usage
*/
func GenerateUsage() {
    var retBuffer bytes.Buffer

    err := compileUsageTemplate(&retBuffer, "generateUsage", getGenerateTemplateData())
    if nil != err {glog.Fatalf("Failed to execute template 'generateUsage': %s", err)}

	fmt.Println(retBuffer.String())
}

func getGenerateTemplateData() (retval map[string]string) {
    retval = ui.GetTemplateVars()

    // Usage vars
    retval["usageCOMMAND"]                     = ui.U("generate")

    // Document command keywords
    retval["toolCommand"]                      = ui.U(ui.WhiteBt("generate"))

    // Document options--
    retval["optName"]                          = ui.WhiteBt("--name")
    retval["optPrefix"]                        = ui.WhiteBt("--prefix")
    retval["optTemplate"]                      = ui.WhiteBt("--template")
    retval["optImage"]                         = ui.WhiteBt("--image")
    retval["optTag"]                           = ui.WhiteBt("--tag")
    retval["optVolume"]                        = ui.WhiteBt("--volume")
    retval["optEnv"]                           = ui.WhiteBt("--env")
    retval["optEntrypoint"]                    = ui.WhiteBt("--entrypoint")
    retval["optCmd"]                           = ui.WhiteBt("--cmd")
    retval["optOption"]                        = ui.WhiteBt("--option")

    // Document highlight words
    retval["highlightRecipe"]                  = ui.WhiteBt(ui.U("recipe"))
    retval["highlightName"]                    = ui.WhiteBt(ui.U("name"))
    retval["highlightPath"]                    = ui.WhiteBt(ui.U("path"))
    retval["highlightTemplate"]                = ui.WhiteBt(ui.U("template"))
    retval["highlightImage"]                   = ui.WhiteBt(ui.U("image"))
    retval["highlightTag"]                     = ui.WhiteBt(ui.U("tag"))
    retval["highlightVolume"]                  = ui.WhiteBt(ui.U("volume"))
    retval["highlightEnv"]                     = ui.WhiteBt(ui.U("key=value"))
    retval["highlightEntrypoint"]              = ui.WhiteBt(ui.U("entrypoint"))
    retval["highlightCommand"]                 = ui.WhiteBt(ui.U("command"))
    retval["highlightOption"]                  = ui.WhiteBt(ui.U("option"))
    retval["highlightOptions"]                 = ui.WhiteBt(ui.U("options"))

    // Command examples
    retval["exampleInstallGulp"]               = ui.I("docker-tools generate gulp-recipe --tag=7.0-alpine")

    return
}
