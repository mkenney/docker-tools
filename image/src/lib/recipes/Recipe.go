/*
Package recipes defines data structures and methods for managing the tool recipe
database.

This portion of the recipe package contains the recipe data structure definition
*/
package recipes

import (
	"bufio"
	"bytes"
	"fmt"
	"lib/config"
	"lib/templates/tool"
	"lib/templates/service"
	"lib/ui"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/golang/glog"
)

/*
Recipe defines a set of properties that define a tool or service management
script
*/
type Recipe struct {
	RecipeName string   // 0.  Recipe name
	ToolName   string   // 1.  Tool name
	Prefix     string   // 2.  Tool install prefix
	Template   string   // 3.  Tool template name (tool, service (daemons))
	Image      string   // 4.  Docker image name
	Tag        string   // 5.  Docker image tag
	Volume     []string // 6.  Volume mount strings
	Env        []string // 7.  Environment variables to pass
	Entrypoint string   // 8.  Container entrypoint
	Cmd        string   // 9.  Tool/container commands
	Option     []string // 10. Additional `docker run` cli arguments
	Notes      string   // 11. Any notes about the recipe
	Source     string   // 12. The name of the sourcefile, either 'recipes' or 'registry'
}

/*
NewRecipe will generate a Recipe model from a data array
*/
func NewRecipe() (recipe *Recipe) {
	recipe = new(Recipe)
	return
}

/*
AddEnv accepts a key=value pair
*/
func (recipe *Recipe) AddEnv(env string) (reterr error) {
	envParts := strings.Split(env, "=")
	if 2 != len(envParts) {
		reterr = fmt.Errorf("Invalid environment variable definition '%s', both `key` and `value` are required", env)
	}
	if nil == reterr {
		for _, curEnv := range recipe.Env {
			if curEnv == env {
				reterr = fmt.Errorf("Environment variable already exists in this recipe '%s'", curEnv)
				break
			}
		}
	}
	if nil == reterr {
		recipe.Env = append(recipe.Env, env)
	}
	return
}

/*
AddOption accepts a key or a key=value pair
*/
func (recipe *Recipe) AddOption(option string) (reterr error) {
	if "" == option {
		reterr = fmt.Errorf("Invalid option definition '%s', a valid `docker run` CLI option is required", option)
	}
	if nil == reterr {
		for _, curOpt := range recipe.Option {
			if curOpt == option {
				reterr = fmt.Errorf("Option variable already exists in this recipe '%s'", curOpt)
				break
			}
		}
	}
	if nil == reterr {
		recipe.Option = append(recipe.Option, option)
	}
	return
}

/*
AddVolume accepts a volume mount string
*/
func (recipe *Recipe) AddVolume(volume string) (reterr error) {
	volParts := strings.Split(volume, ":")
	if 2 > len(volParts) {
		reterr = fmt.Errorf("Invalid volume definition '%s', both `src` and `dest` are required", volume)
	}
	if nil == reterr {
		for _, vol := range recipe.Volume {
			if vol == volume {
				reterr = fmt.Errorf("Volume already exists in this recipe '%s'", vol)
				break
			}
		}
	}
	if nil == reterr {
		recipe.Volume = append(recipe.Volume, volume)
	}
	return
}

/*
Render states:
 1. recipe installed
     1. up to date
     2. out of date
 2. recipe not installed
     1. no tool installed
     2. docker-tools tool installed - recipe name
     3. unmanaged tool installed (any other file)
*/
func (recipe *Recipe) Render() string {

	template, err := template.New("recipe").Parse(`
 {{.Source}} {{.RecipeName}}
{{.Notes}}
   Details
     ├─ status ` + ui.Grey("──────") + ` {{.RecipeStatus}}
     ├─ type ` + ui.Grey("────────") + ` {{.Template}}
     ├─ tool ` + ui.Grey("────────") + ` {{.Prefix}}/{{.ToolName}}
     ├─ image ` + ui.Grey("───────") + ` {{.Image}}:{{.Tag}}
     ├─ cmd ` + ui.Grey("─────────") + ` {{.Cmd}}
     ├─ entrypoint ` + ui.Grey("──") + ` {{.Entrypoint}}
     ├─ env {{.Env}}
     └─ volumes {{.Volumes}}
`)
	if nil != err {
		glog.Fatalf("Could not parse Recipe template: %s", err)
	}

	var retBuffer bytes.Buffer
	err = template.Execute(&retBuffer, recipe.renderVars())
	if nil != err {
		glog.Fatalf("Could not execute Recipe template: %s", err)
	}

	return retBuffer.String()
}

/*
ToString returns a string representation of the recipe (a `docker run` command)
*/
func (recipe *Recipe) ToString() string {
	var template *template.Template

	if "tool" == recipe.Template {
		template = tool.Template
	} else if "service" == recipe.Template {
		template = service.Template
	}

	var retBuffer bytes.Buffer
	err := template.Execute(&retBuffer, recipe.toStringVars())
	if nil != err {
		glog.Fatalf("Could not execute Recipe string template: %s", err)
	}

	return retBuffer.String()
}

/*
renderVars
*/
func (recipe *Recipe) renderVars() (retval map[string]string) {
	retval = make(map[string]string)

	retval["Source"] = recipe.renderSource()
	retval["RecipeName"] = recipe.renderRecipeName()
	retval["Notes"] = recipe.renderNotes()
	retval["RecipeStatus"] = recipe.renderStatus()
	retval["Template"] = recipe.renderTemplate()
	retval["Prefix"] = recipe.renderPrefix()
	retval["ToolName"] = recipe._toolName()
	//retval["ToolStatus"]          = recipe._toolStatus()
	retval["Image"] = recipe.renderImage()
	retval["Tag"] = recipe.renderTag()
	retval["Cmd"] = recipe.renderCmd()
	retval["Entrypoint"] = recipe.renderEntrypoint()
	retval["Volumes"] = recipe.renderVolumes()
	retval["Env"] = recipe.renderEnv()
	//retval["Options"]       = recipe.renderOption()

	return
}

/*
renderVars
*/
func (recipe *Recipe) toStringVars() (retval map[string]string) {
	retval = make(map[string]string)

	retval["DockerToolsVersion"] = escapeBashVar(config.DockerToolsVersion)
	retval["DefaultPrefix"] = escapeBashVar(config.DockerToolsDefaultPrefixDir)
	retval["RecipeName"] = escapeBashVar(recipe.stringRecipeName())
	retval["ToolName"] = escapeBashVar(recipe.stringToolName())
	retval["Prefix"] = escapeBashVar(recipe.stringPrefix())
	retval["Template"] = escapeBashVar(recipe.stringTemplate())
	retval["Image"] = escapeBashVar(recipe.stringImage())
	retval["Tag"] = escapeBashVar(recipe.stringTag())
	retval["Volume"] = escapeBashVar(recipe.stringVolume())
	retval["Env"] = escapeBashVar(recipe.stringEnv())
	retval["Entrypoint"] = escapeBashVar(recipe.stringEntrypoint())
	retval["Cmd"] = escapeBashVar(recipe.stringCmd())
	retval["Option"] = escapeBashVar(recipe.stringOption())
	retval["Notes"] = escapeBashVar(recipe.stringNotes())
	retval["Source"] = escapeBashVar(recipe.stringSource())

	return
}

func escapeBashVar(str string) (retval string) {
	retval = str
	retval = strings.Replace(retval, "\"", "\\\"", -1)
	return
}

func (recipe *Recipe) stringSource() string {
	return escapeBashVar(recipe.Source)
}
func (recipe *Recipe) stringNotes() string {
	return escapeBashVar(recipe.Notes)
}
func (recipe *Recipe) stringCmd() string {
	return escapeBashVar(recipe.Cmd)
}
func (recipe *Recipe) stringTemplate() string {
	return escapeBashVar(recipe.Template)
}
func (recipe *Recipe) stringRecipeName() string {
	return escapeBashVar(recipe.RecipeName)
}
func (recipe *Recipe) stringToolName() string {
	return escapeBashVar(recipe.ToolName)
}
func (recipe *Recipe) stringPrefix() string {
	retval := recipe.Prefix
	if "" == retval {
		retval = config.DockerToolsDefaultPrefixDir
	}
	return escapeBashVar(retval)
}
func (recipe *Recipe) stringImage() string {
	return escapeBashVar(recipe.Image)
}
func (recipe *Recipe) stringTag() (retval string) {
	retval = escapeBashVar(recipe.Tag)
	if "" == retval {
		retval = "latest"
	}
	return
}
func (recipe *Recipe) stringEntrypoint() (retval string) {
	if "" != recipe.Entrypoint {
		retval = "--entrypoint=\""+escapeBashVar(recipe.Entrypoint)+"\""
	}
	return
}
func (recipe *Recipe) stringVolume() (retval string) {
	for _, volume := range recipe.Volume {
		retval += "-v \""+escapeBashVar(volume)+"\" "
	}
	if "docker-tools" == recipe.RecipeName && "registry" == recipe.Source {
		for _, path := range config.HostPath {
			retval += "-v \""+path+":/host"+path+":ro\" "
		}
	}
	return
}
func (recipe *Recipe) stringEnv() (retval string) {
	for _, env := range recipe.Env {
		retval += "-e \""+escapeBashVar(env)+"\" "
	}
	return
}
func (recipe *Recipe) stringOption() string {
	return strings.Join(recipe.Option, " ")
}

/*
renderVolumes
*/
func (recipe *Recipe) renderVolumes() string {
	var status string

	if 0 < len(recipe.Volume) {
		for _, volstr := range recipe.Volume {
			volparts := strings.Split(volstr, ":")

			volmode := ""
			modeicon := ""
			if 3 == len(volparts) {
				modeparts := strings.Split(volparts[2], ",")
				modelen := 0
				for k, v := range modeparts {
					switch v {
					case "rw":
						modeparts[k] = ui.RedBt("rw")
						modeicon = "⇿"
					case "ro":
						modeparts[k] = ui.BlueDk("ro")
						if "" == modeicon {
							modeicon = "⇾"
						}
					default:
						modeparts[k] = ui.Grey(v)
						if "⇿" != modeicon {
							modeicon = "⇝"
						}
					}
					modelen += len(v)
				}
				if 0 < len(modeparts) {
					modelen += (len(modeparts) - 1)
				}
				volmode = strings.Join(modeparts, ", ")
			}

			status += fmt.Sprintf("\n         %s  %s:%s %s", modeicon, volparts[0], volparts[1], volmode)
		}

	} else {
		status = ui.Grey("───── n/a")
	}

	return status
}

/*
renderEnv
*/
func (recipe *Recipe) renderEnv() string {
	var status string

	// Env is a JSON byte array
	if 0 < len(recipe.Env) {
		for _, envstr := range recipe.Env {
			envvar := strings.Split(envstr, "=")
			if 1 == len(envvar) {
				glog.Fatalf("Invalid environment variable '%s' in recipe %s:%s", envstr, recipe.RecipeName, recipe.Source)
			}
			status += fmt.Sprintf("\n     │    - %v=%v", envvar[0], envvar[1])
		}

	} else {
		status = ui.Grey("───────── n/a")
	}

	return status
}

/*
renderEntrypoint
*/
func (recipe *Recipe) renderEntrypoint() string {
	status := recipe.Entrypoint
	if "" == status {
		status = ui.Grey("n/a")
	}
	return status
}

/*
renderCmd
*/
func (recipe *Recipe) renderCmd() string {
	status := recipe.Cmd
	if "" == status {
		status = ui.Grey("n/a")
	}
	return status
}

/*
renderImage
*/
func (recipe *Recipe) renderImage() string {
	return recipe.Image
}

/*
renderTag
*/
func (recipe *Recipe) renderTag() (retval string) {
	retval = recipe.Tag
	if "" == retval {
		retval = "latest"
	}
	return
}

/*
renderSource
*/
func (recipe *Recipe) renderSource() string {
	styled := ui.Indigo("●")
	if "recipes" == recipe.Source {
		styled = ui.Custom(52, 0, "●")
	}
	return styled
}

/*
renderRecipeName
*/
func (recipe *Recipe) renderRecipeName() string {
	var style string
	if "recipes" == recipe.Source {
		style = ui.WhiteBt(ui.U(ui.I(recipe.RecipeName)))
	} else {
		style = ui.B(ui.U(recipe.RecipeName))
	}
	return style
}

/*
renderNotes will word-wrap multi-line text blobs at column `maxlen` with a `margin` indent
*/
func (recipe *Recipe) renderNotes() string {
	formatted := "\n"

	margin := "   "
	maxlen := 80

	for _, line := range strings.Split(recipe.Notes, "\n") {
		tmpline := ""
		for _, word := range strings.Split(line, " ") {
			if maxlen < len(margin)+len(tmpline)+len(word) {
				formatted += margin + tmpline + "\n"
				tmpline = ""
			}
			tmpline += word + " "
		}
		formatted += margin + tmpline + "\n"
		tmpline = ""
	}
	if "" == strings.Trim(formatted, " \n") {
		formatted = ""
	}

	return formatted
}

/*
renderTemplate
*/
func (recipe *Recipe) renderTemplate() string {
	return recipe.Template
}

/*
renderPrefix
*/
func (recipe *Recipe) renderPrefix() (retval string) {
	retval = strings.Replace(recipe.Prefix, "$HOME", config.HostHome, -1)
	if "" == retval {
		retval = config.DockerToolsDefaultPrefixDir
	}
	return
}

/*
_toolName
*/
func (recipe *Recipe) _toolName() string {
	return recipe.ToolName
}

/*
_toolStatus
*/
func (recipe *Recipe) _toolStatus() string {
	var status string

	hostprefix := strings.Replace(recipe.Prefix, "$HOME", config.HostHome, -1)
	if _, err := os.Stat("/host" + hostprefix + "/" + recipe.ToolName); os.IsNotExist(err) {
		status = ui.Grey(ui.B("not installed"))

	} else {
		status = ui.Red(ui.B("unmanaged"))

		if file, err := os.Open("/host" + hostprefix + "/" + recipe.ToolName); err != nil {
			glog.Fatalf("%s", err)

		} else {
			defer file.Close()
			if stats, statserr := file.Stat(); nil == statserr && !stats.IsDir() {
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					if strings.Contains(scanner.Text(), "TOOLS_VERSION") {
						if strings.Contains(scanner.Text(), "RECIPE_NAME=\""+recipe.RecipeName+"\"") {
							status = ui.OrangeBt(ui.B("outdated"))
							if strings.Contains(scanner.Text(), "TOOLS_VERSION=\""+config.DockerToolsVersion+"\"") {
								status = ui.GreenBt(ui.B("installed"))
							}
						}
					}
				}
				if err = scanner.Err(); err != nil {
					glog.Fatalf("%s", err)
				}
			}
		}
	}

	return status
}

/*
renderStatus
 1. recipe installed
 2. recipe not installed
     1. no tool installed
     2. docker-tools tool installed
         1. up to date
         2. out of date
     3. unmanaged tool installed (any other file)
*/
func (recipe *Recipe) renderStatus() string {
	var status string

	hostprefix := strings.Replace(recipe.Prefix, "$HOME", config.HostHome, -1)
	if _, err := os.Stat("/host" + hostprefix + "/" + recipe.ToolName); os.IsNotExist(err) {
		status = ui.Grey(ui.B("not installed"))

	} else {
		if file, err := os.Open("/host" + hostprefix + "/" + recipe.ToolName); err != nil {
			glog.Fatalf("%s", err)

		} else {
			defer file.Close()
			status = ui.Red(ui.B("unmanaged file installed"))

			if stats, statserr := file.Stat(); nil == statserr && !stats.IsDir() {
				reader := bufio.NewReader(file)

				var line []byte
				var readerr error

				var istool bool
				var isrecipe bool
				var isupdated bool
				var toolrecipe string

				for readerr == nil {
					line, _, readerr = reader.ReadLine()
					if strings.Contains(string(line), "RECIPE_NAME") {
						istool = true

						// Get the recipe name
						re := regexp.MustCompile("RECIPE_NAME=\"(.*)\"")
						matches := re.FindStringSubmatch(string(line))
						if 2 <= len(matches) {
							toolrecipe = matches[1]
							if recipe.RecipeName == toolrecipe {
								isrecipe = true
							}
						}
					}
					if strings.Contains(string(line), "DOCKER_TOOLS_VERSION=\""+config.DockerToolsVersion+"\"") {
						isupdated = true
					}
				}

				if isrecipe {
					status = ui.OrangeBt(ui.B("⚠  update available"))
					if isupdated {
						status = ui.GreenBt(ui.B("✓ installed"))
					}
				} else if istool {
					status = ui.GreenDk(ui.B("⇅ recipe '" + toolrecipe + "' installed"))
					if !isupdated {
						status = ui.YellowDk(ui.B("⇅ recipe '" + toolrecipe + "' installed"))
					}
				}
			}
		}
	}

	return status
}
