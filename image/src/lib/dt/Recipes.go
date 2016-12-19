/*
This portion of the dt package contains methods for managing recipe-related
commands
*/
package dt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"glog"
	"lib/config"
	"lib/recipes"
	"lib/ui"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

/*
ListRecipes iterates through all known recipies and passes their data to a
formatting method which returns CLI-formatted output. It then passes that output
through a pager (`less -r`).
*/
func (dt *DockerTools) ListRecipes() {
	var listing string

	for _, recipe := range dt.Recipes {
		if nil == recipe {
			break
		}
		listing += dt.render(dt.templateVars(recipe))
	}

	pipestdin, pipestdout, err := os.Pipe()
	if err != nil {
		panic("Could not create pipe")
	}

	stdout := os.Stdout
	os.Stdout = pipestdout

	pager := exec.Command("less", "-r")
	pager.Stdin = pipestdin
	pager.Stdout = stdout // the pager uses the original stdout, not the pipe
	pager.Stderr = os.Stderr

	defer func() {
		// Close the pipe
		pipestdout.Close()
		// Run the pager
		if err := pager.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// restore stdout
		os.Stdout = stdout
	}()

	fmt.Println("\n\n" + listing)
}

/*
render states:
 1. recipe installed
     1. up to date
     2. out of date
 2. recipe not installed
     1. no tool installed
     2. docker-tools tool installed - recipe name
     3. unmanaged tool installed (any other file)
*/
func (dt *DockerTools) render(d map[string]string) string {
	statusHr := ui.Grey("──────")
	typeHr := ui.Grey("────────")
	toolHr := ui.Grey("────────")
	imageHr := ui.Grey("───────")
	cmdHr := ui.Grey("─────────")
	entrypointHr := ui.Grey("──")
	volumesHr := ""
	envHr := ""

	return fmt.Sprintf(`
 ` + d["RecipeSource"] + ` ` + d["RecipeName"] + `
` + d["RecipeNotes"] + `
   Details
     ├─ status ` + statusHr + ` ` + d["RecipeStatus"] + `
     ├─ type ` + typeHr + ` ` + d["ToolTemplate"] + `
     ├─ tool ` + toolHr + ` ` + d["ToolPrefix"] + `/` + d["ToolName"] + `
     ├─ image ` + imageHr + ` ` + d["DockerImage"] + `:` + d["DockerTag"] + `
     ├─ cmd ` + cmdHr + ` ` + d["ContainerCmd"] + `
     ├─ entrypoint ` + entrypointHr + ` ` + d["ContainerEntrypoint"] + `
     ├─ env ` + envHr + ` ` + d["ContainerEnv"] + `
     └─ volumes` + volumesHr + ` ` + d["ContainerVolumes"] + `

`)
	//──────────────────────────────────────────────────────────────────────────────

}

/*
templateVars
*/
func (dt *DockerTools) templateVars(recipe *recipes.Recipe) map[string]string {
	retval := make(map[string]string)

	retval["RecipeSource"] = dt._recipeSource(recipe)
	retval["RecipeName"] = dt._recipeName(recipe)
	retval["RecipeNotes"] = dt._recipeNotes(recipe)
	retval["RecipeStatus"] = dt._recipeStatus(recipe)
	retval["ToolTemplate"] = dt._toolTemplate(recipe)
	retval["ToolPrefix"] = dt._toolPrefix(recipe)
	retval["ToolName"] = dt._toolName(recipe)
	//    retval["ToolStatus"]          = dt._toolStatus(recipe)
	retval["DockerImage"] = dt._dockerImage(recipe)
	retval["DockerTag"] = dt._dockerTag(recipe)
	retval["ContainerCmd"] = dt._containerCmd(recipe)
	retval["ContainerEntrypoint"] = dt._containerEntrypoint(recipe)
	retval["ContainerVolumes"] = dt._containerVolumes(recipe)
	retval["ContainerEnv"] = dt._containerEnv(recipe)
	//    retval["DockerOptions"]       = dt._dockerOptions(recipe)

	return retval
}

/*
_containerVolumes
*/
func (dt *DockerTools) _containerVolumes(recipe *recipes.Recipe) string {
	var status string

	// ContainerVolumes is a JSON byte array
	if 0 < len(recipe.ContainerVolumes) {
		var data interface{}
		json.Unmarshal(recipe.ContainerVolumes, &data)
		for _, volstr := range data.([]interface{}) {
			volparts := strings.Split(volstr.(string), ":")

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
_containerEnv
*/
func (dt *DockerTools) _containerEnv(recipe *recipes.Recipe) string {
	var status string

	// ContainerEnv is a JSON byte array
	if 0 < len(recipe.ContainerEnv) {
		var data interface{}
		json.Unmarshal(recipe.ContainerEnv, &data)
		for _, envstr := range data.([]interface{}) {
			envvar := strings.Split(envstr.(string), "=")
			status += fmt.Sprintf("\n     │    - %v=%v", envvar[0], envvar[1])
		}

	} else {
		status = ui.Grey("──────── n/a")
	}

	return status
}

/*
_containerEntrypoint
*/
func (dt *DockerTools) _containerEntrypoint(recipe *recipes.Recipe) string {
	status := recipe.ContainerEntrypoint
	if "" == status {
		status = ui.Grey("n/a")
	}
	return status
}

/*
_containerCmd
*/
func (dt *DockerTools) _containerCmd(recipe *recipes.Recipe) string {
	status := recipe.ContainerCmd
	if "" == status {
		status = ui.Grey("n/a")
	}
	return status
}

/*
_dockerImage
*/
func (dt *DockerTools) _dockerImage(recipe *recipes.Recipe) string {
	return recipe.DockerImage
}

/*
_dockerTag
*/
func (dt *DockerTools) _dockerTag(recipe *recipes.Recipe) string {
	return recipe.DockerTag
}

/*
_recipeSource
*/
func (dt *DockerTools) _recipeSource(recipe *recipes.Recipe) string {
	styled := ui.Indigo("●")
	if "recipes" == recipe.RecipeSource {
		styled = ui.Custom(52, 0, "●")
	}
	return styled
}

/*
_recipeName
*/
func (dt *DockerTools) _recipeName(recipe *recipes.Recipe) string {
	var style string
	if "recipes" == recipe.RecipeSource {
		style = ui.WhiteBt(ui.U(ui.I(recipe.RecipeName)))
	} else {
		style = ui.B(ui.U(recipe.RecipeName))
	}
	return style
}

/*
_recipeNotes will word-wrap multi-line text blobs at column `maxlen` with a `margin` indent
*/
func (dt *DockerTools) _recipeNotes(recipe *recipes.Recipe) string {
	formatted := "\n"

	margin := "   "
	maxlen := 80

	for _, line := range strings.Split(recipe.RecipeNotes, "\n") {
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
_toolTemplate
*/
func (dt *DockerTools) _toolTemplate(recipe *recipes.Recipe) string {
	return recipe.ToolTemplate
}

/*
_toolPrefix
*/
func (dt *DockerTools) _toolPrefix(recipe *recipes.Recipe) string {
	recipe.ToolPrefix = strings.Replace(recipe.ToolPrefix, "$HOME", config.Config.HostHome, -1)
	return recipe.ToolPrefix
}

/*
_toolName
*/
func (dt *DockerTools) _toolName(recipe *recipes.Recipe) string {
	return recipe.ToolName
}

/*
_toolStatus
*/
func (dt *DockerTools) _toolStatus(recipe *recipes.Recipe) string {
	var status string

	hostprefix := strings.Replace(recipe.ToolPrefix, "$HOME", config.Config.HostHome, -1)
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
							if strings.Contains(scanner.Text(), "__TOOLS_VERSION__="+config.Config.ToolsVersion) {
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
_recipeStatus
 1. recipe installed
 2. recipe not installed
     1. no tool installed
     2. docker-tools tool installed
         1. up to date
         2. out of date
     3. unmanaged tool installed (any other file)
*/
func (dt *DockerTools) _recipeStatus(recipe *recipes.Recipe) string {
	var status string

	hostprefix := strings.Replace(recipe.ToolPrefix, "$HOME", config.Config.HostHome, -1)
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
					if strings.Contains(string(line), "__TOOLS_VERSION__="+config.Config.ToolsVersion) {
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
