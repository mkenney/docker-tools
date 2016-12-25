/*
Package recipes defines data structures and methods for managing the tool recipe
database.

This portion of the recipe package contains the recipe data structure definition
*/
package recipes

import (
	"bufio"
	"bytes"
	"encoding/json"
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
func NewRecipe(recipeData []string) (recipe *Recipe) {
	recipe = new(Recipe)
	recipe.RecipeName = recipeData[0]
	recipe.ToolName = recipeData[1]
	recipe.Prefix = recipeData[2]
	if "" == recipe.Prefix {
		recipe.Prefix = config.Values.DefaultPrefix
	}
	recipe.Template = recipeData[3]
	recipe.Image = recipeData[4]
	recipe.Tag = recipeData[5]
	if "" == recipe.Tag {
		recipe.Tag = "latest"
	}

	var jsonData interface{}

	json.Unmarshal([]byte(recipeData[6]), &jsonData)
	if nil != jsonData {
		for _, volstr := range jsonData.([]interface{}) {
			recipe.Volume = append(recipe.Volume, volstr.(string))
		}
	}

	json.Unmarshal([]byte(recipeData[7]), &jsonData)
	if nil != jsonData {
		for _, volstr := range jsonData.([]interface{}) {
			recipe.Env = append(recipe.Env, volstr.(string))
		}
	}
	recipe.Entrypoint = recipeData[8]
	recipe.Cmd = recipeData[9]

	json.Unmarshal([]byte(recipeData[10]), &jsonData)
	if nil != jsonData {
		for _, volstr := range jsonData.([]interface{}) {
			recipe.Option = append(recipe.Option, volstr.(string))
		}
	}

	recipe.Notes = recipeData[11]
	recipe.Source = recipeData[12]
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
func (recipe *Recipe) renderVars() map[string]string {
	retval := make(map[string]string)

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
	return retval
}

/*
renderVars
*/
func (recipe *Recipe) toStringVars() (retval map[string]string) {
	retval = make(map[string]string)

	retval["DockerToolsVersion"] = escapeBashVar(config.DockerToolsVersion)
	retval["RecipeName"] = escapeBashVar(recipe.RecipeName)
	retval["ToolName"] = escapeBashVar(recipe.ToolName)
	retval["Prefix"] = escapeBashVar(recipe.Template)
	retval["Template"] = escapeBashVar(recipe.Template)
	retval["Image"] = escapeBashVar(recipe.Image)
	retval["Tag"] = escapeBashVar(recipe.Tag)
	retval["Volumes"] = escapeBashVar("-v "+strings.Join(recipe.Volume, " -v "))
	retval["Env"] = escapeBashVar("-e "+strings.Join(recipe.Env, " -e "))
	retval["Entrypoint"] = escapeBashVar(recipe.Entrypoint)
	retval["Cmd"] = escapeBashVar(recipe.Cmd)
	retval["Option"] = escapeBashVar(strings.Join(recipe.Option, " "))
	retval["Notes"] = escapeBashVar(recipe.Notes)
	retval["Source"] = escapeBashVar(recipe.Source)

	return
}

func escapeBashVar(str string) (retval string) {
	retval = str
	retval = strings.Replace(retval, "\"", "\\\"", -1)
	return
}

func (recipe *Recipe) stringImage() string {
	return recipe.Image
}
func (recipe *Recipe) stringTag() string {
	return recipe.Tag
}
func (recipe *Recipe) stringCmd() string {
	return recipe.Cmd
}
func (recipe *Recipe) stringEntrypoint() string {
	return "--entrypoint=\""+recipe.Entrypoint+"\""
}
func (recipe *Recipe) stringVolumes() (retval string) {
	for _, volume := range recipe.Volume {
		retval += "-v "+volume+" "
	}
	return
}
func (recipe *Recipe) stringEnv() (retval string) {
	for _, env := range recipe.Env {
		retval += "-e \""+env+"\" "
	}
	return
}
func (recipe *Recipe) stringOption() (retval string) {
	for _, opt := range recipe.Option {
		retval += "-e \""+opt+"\" "
	}
	return
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
				//volmode = fmt.Sprintf("% "+strconv.Itoa(9 - modelen)+"s", "")+strings.Join(modeparts, ", ")
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
func (recipe *Recipe) renderTag() string {
	return recipe.Tag
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
func (recipe *Recipe) renderPrefix() string {
	recipe.Prefix = strings.Replace(recipe.Prefix, "$HOME", config.Values.HostHome, -1)
	return recipe.Prefix
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

	hostprefix := strings.Replace(recipe.Prefix, "$HOME", config.Values.HostHome, -1)
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
					if strings.Contains(scanner.Text(), "__TOOLS_VERSION__") {
						if strings.Contains(scanner.Text(), "__RECIPE_NAME__="+recipe.RecipeName) {
							status = ui.OrangeBt(ui.B("outdated"))
							if strings.Contains(scanner.Text(), "__TOOLS_VERSION__="+config.DockerToolsVersion) {
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

	hostprefix := strings.Replace(recipe.Prefix, "$HOME", config.Values.HostHome, -1)
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
					if strings.Contains(string(line), "__RECIPE_NAME__") {
						istool = true

						// Get the recipe name
						re := regexp.MustCompile("^__RECIPE_NAME__=(.*)$")
						matches := re.FindStringSubmatch(string(line))
						toolrecipe = matches[1]
						if recipe.RecipeName == toolrecipe {
							isrecipe = true
						}
					}
					if strings.Contains(string(line), "__DOCKER_TOOLS_VERSION__="+config.DockerToolsVersion) {
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
