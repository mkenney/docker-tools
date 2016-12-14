package dt
import (
    "bufio"
    "encoding/json"
    "fmt"
    "lib/recipes"
    "lib/config"
    "lib/ui"
    "log"
    "os"
    "os/exec"
    "regexp"
//    "strconv"
    "strings"
)


/**
 * List current recipes to stdout
 * @param  {[type]} this *DockerTools) ListRecipes( [description]
 * @return {[type]}      [description]
 */
func (this *DockerTools) ListRecipes() {
    var listing string

    for _, recipe := range this.Recipes {
        if nil == recipe {break}
        listing += this.render(this.templateVars(recipe))
    }
//ui.Test()
//os.Exit(0)
    pipe_stdin, pipe_stdout, err := os.Pipe()
    if err != nil {
        panic("Could not create pipe")
    }

    stdout := os.Stdout
    os.Stdout = pipe_stdout

    pager := exec.Command("less","-r")
    pager.Stdin = pipe_stdin
    pager.Stdout = stdout // the pager uses the original stdout, not the pipe
    pager.Stderr = os.Stderr

    defer func() {
        // Close the pipe
        pipe_stdout.Close()
        // Run the pager
        if err := pager.Run(); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        // restore stdout
        os.Stdout = stdout
    }()

    fmt.Println("\n\n"+listing)
}

/**
 * states:
 *  1. recipe installed
 *      1. up to date
 *      2. out of date
 *  2. recipe not installed
 *      1. no tool installed
 *      2. docker-tools tool installed - recipe name
 *      3. unmanaged tool installed (any other file)
 *
 */
func (this *DockerTools) render(d map[string]string) (string) {
    type_hr := ui.Grey("────────")
    tool_hr := ui.Grey("────────")
    image_hr := ui.Grey("───────")
    cmd_hr := ui.Grey("─────────")
    entrypoint_hr := ui.Grey("──")
    volumes_hr := ""
    env_hr := ""

    return fmt.Sprintf(`
 `+d["RecipeSource"]+` `+d["RecipeName"]+` `+d["RecipeStatus"]+`
`+d["RecipeNotes"]+`
   Details
     ├─ type `+type_hr+` `+d["ToolTemplate"]+`
     ├─ tool `+tool_hr+` `+d["ToolPrefix"]+`/`+d["ToolName"]+` `+d["ToolStatus"]+`
     ├─ image `+image_hr+` `+d["DockerImage"]+`:`+d["DockerTag"]+`
     ├─ cmd `+cmd_hr+` `+d["ContainerCmd"]+`
     ├─ entrypoint `+entrypoint_hr+` `+d["ContainerEntrypoint"]+`
     ├─ env `+env_hr+` `+d["ContainerEnv"]+`
     └─ volumes`+volumes_hr+` `+d["ContainerVolumes"]+`

`)
//──────────────────────────────────────────────────────────────────────────────

}

/**
 */
func (this *DockerTools) templateVars(recipe *recipes.Recipe) (map[string]string) {
    ret_val := make(map[string]string)

    ret_val["RecipeSource"]        = this._recipeSource(recipe)
    ret_val["RecipeName"]          = this._recipeName(recipe)
    ret_val["RecipeNotes"]         = this._recipeNotes(recipe)
    ret_val["RecipeStatus"]        = this._recipeStatus(recipe)
    ret_val["ToolTemplate"]        = this._toolTemplate(recipe)
    ret_val["ToolPrefix"]          = this._toolPrefix(recipe)
    ret_val["ToolName"]            = this._toolName(recipe)
    ret_val["ToolStatus"]          = this._toolStatus(recipe)
    ret_val["DockerImage"]         = this._dockerImage(recipe)
    ret_val["DockerTag"]           = this._dockerTag(recipe)
    ret_val["ContainerCmd"]        = this._containerCmd(recipe)
    ret_val["ContainerEntrypoint"] = this._containerEntrypoint(recipe)
    ret_val["ContainerVolumes"]    = this._containerVolumes(recipe)
    ret_val["ContainerEnv"]        = this._containerEnv(recipe)
//    ret_val["DockerOptions"]       = this._dockerOptions(recipe)

    return ret_val
}

/**
 */
func (this *DockerTools) _containerVolumes(recipe *recipes.Recipe) (string) {
    var status string = ""

    // ContainerVolumes is a JSON byte array
    if 0 < len(recipe.ContainerVolumes) {
        var data interface{}
        json.Unmarshal(recipe.ContainerVolumes, &data)
        for _, vol_str := range data.([]interface{}) {
            vol_parts := strings.Split(vol_str.(string), ":")

            vol_mode := ""
            mode_icon := ""
            if 3 == len(vol_parts) {
                mode_parts := strings.Split(vol_parts[2], ",")
                mode_len := 0
                for k, v := range mode_parts {
                    switch v {
                        case "rw":
                            mode_parts[k] = ui.RedBt("rw")
                            mode_icon = "⇿"
                        case "ro":
                            mode_parts[k] = ui.BlueDk("ro")
                            if "" == mode_icon {mode_icon = "⇾"}
                        default:
                            mode_parts[k] = ui.Grey(v)
                            if "⇿" != mode_icon {mode_icon = "⇝"}
                    }
                    mode_len += len(v)
                }
                if 0 < len(mode_parts) {
                    mode_len += (len(mode_parts) - 1)
                }
                //vol_mode = fmt.Sprintf("% "+strconv.Itoa(9 - mode_len)+"s", "")+strings.Join(mode_parts, ", ")
                vol_mode = strings.Join(mode_parts, ", ")
            }

            status += fmt.Sprintf("\n         %s  %s:%s %s", mode_icon, vol_parts[0], vol_parts[1], vol_mode)
        }

    } else {
        status = ui.Grey("───── n/a")
    }

    return status
}

/**
 */
func (this *DockerTools) _containerEnv(recipe *recipes.Recipe) (string) {
    var status string = ""

    // ContainerEnv is a JSON byte array
    if 0 < len(recipe.ContainerEnv) {
        var data interface{}
        json.Unmarshal(recipe.ContainerEnv, &data)
        for _, env_str := range data.([]interface{}) {
            env_var := strings.Split(env_str.(string), "=")
            status += fmt.Sprintf("\n     │    - %v=%v", env_var[0], env_var[1])
        }

    } else {
        status = ui.Grey("──────── n/a")
    }

    return status
}

/**
 */
func (this *DockerTools) _containerEntrypoint(recipe *recipes.Recipe) (string) {
    var status string = recipe.ContainerEntrypoint
    if "" == status {
        status = ui.Grey("n/a")
    }
    return status
}

/**
 */
func (this *DockerTools) _containerCmd(recipe *recipes.Recipe) (string) {
    var status string = recipe.ContainerCmd
    if "" == status {
        status = ui.Grey("n/a")
    }
    return status
}

/**
 */
func (this *DockerTools) _dockerImage(recipe *recipes.Recipe) (string) {
    return recipe.DockerImage
}

/**
 */
func (this *DockerTools) _dockerTag(recipe *recipes.Recipe) (string) {
    return recipe.DockerTag
}

/**
 */
func (this *DockerTools) _recipeSource(recipe *recipes.Recipe) (string) {
//ui.Test()
//os.Exit(0)
    var styled string = ui.Indigo("●")
    if "recipes" == recipe.RecipeSource {
//        styled = ui.GreenDk("●")
        styled = ui.Custom(52, 0, "●")
    }
    return styled
}

/**
 */
func (this *DockerTools) _recipeName(recipe *recipes.Recipe) (string) {
    var style string
    if "recipes" == recipe.RecipeSource {
        style = ui.WhiteBt(ui.U(ui.I(recipe.RecipeName)))
    } else {
        style = ui.B(ui.U(recipe.RecipeName))
    }
    return style
}

/**
 * Word-wrap multi-line text blobs at column `max_len` with a `margin` indent
 *
 * @param  recipes.Recipe
 * @return string
 */
func (this *DockerTools) _recipeNotes(recipe *recipes.Recipe) (string) {
    var formatted string = "\n"

    var margin string = "   "
    var max_len int = 80

    for _, line := range strings.Split(recipe.RecipeNotes, "\n") {
        tmp_line := ""
        for _, word := range strings.Split(line, " ") {
            if max_len < len(margin) + len(tmp_line) + len(word) {
                formatted += margin+tmp_line+"\n"
                tmp_line = ""
            }
            tmp_line += word+" "
        }
        formatted += margin+tmp_line+"\n"
        tmp_line = ""
    }
    if "" == strings.Trim(formatted, " \n") {formatted = ""}

    return formatted
}

/**
 */
func (this *DockerTools) _toolTemplate(recipe *recipes.Recipe) (string) {
    return recipe.ToolTemplate
}

/**
 */
func (this *DockerTools) _toolPrefix(recipe *recipes.Recipe) (string) {
    recipe.ToolPrefix = strings.Replace(recipe.ToolPrefix, "$HOME", config.Config.HostHome, -1)
    return recipe.ToolPrefix
}

/**
 */
func (this *DockerTools) _toolName(recipe *recipes.Recipe) (string) {
    return recipe.ToolName
}

/**
 */
func (this *DockerTools) _toolStatus(recipe *recipes.Recipe) (string) {
    var status string

    host_prefix := strings.Replace(recipe.ToolPrefix, "$HOME", config.Config.HostHome, -1)
    if _, err := os.Stat("/host"+host_prefix+"/"+recipe.ToolName); os.IsNotExist(err) {
        status = ui.Grey(ui.B("not installed"))

    } else {
        status = ui.Red(ui.B("unmanaged"))

        if file, err := os.Open("/host"+host_prefix+"/"+recipe.ToolName); err != nil {
            log.Fatal(err)

        } else {
            defer file.Close()
            if stats, stats_err := file.Stat(); nil == stats_err && !stats.IsDir() {
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
                  log.Fatal(err)
                }
            }
        }
    }

    return status
}

/**
 *
 *  1. recipe installed
 *  2. recipe not installed
 *      1. no tool installed
 *      2. docker-tools tool installed
 *          1. up to date
 *          2. out of date
 *      3. unmanaged tool installed (any other file)
 *
 * @param  {[type]} recipe_name [description]
 * @param  {[type]} tool_path   string)       (string [description]
 * @return {[type]}             [description]
 */
func (this *DockerTools) _recipeStatus(recipe *recipes.Recipe) (string) {
    var status string

    host_prefix := strings.Replace(recipe.ToolPrefix, "$HOME", config.Config.HostHome, -1)
    if _, err := os.Stat("/host"+host_prefix+"/"+recipe.ToolName); os.IsNotExist(err) {
        status = ui.Grey(ui.B("not installed"))

    } else {
        if file, err := os.Open("/host"+host_prefix+"/"+recipe.ToolName); err != nil {
            log.Fatal(err)

        } else {
            defer file.Close()
            status = ui.Orange(ui.B("unmanaged file installed"))

            if stats, stats_err := file.Stat(); nil == stats_err && !stats.IsDir() {
                reader := bufio.NewReader(file)

                var line []byte
                var read_err error

                var is_tool bool
                var is_recipe bool
                var is_updated bool
                var tool_recipe string

                for read_err == nil {
                    line, _, read_err = reader.ReadLine()
                    if strings.Contains(string(line), "__RECIPE_NAME__") {
                        is_tool = true

                        // Get the recipe name
                        re := regexp.MustCompile("^__RECIPE_NAME__=(.*)$")
                        matches := re.FindStringSubmatch(string(line))
                        tool_recipe = matches[1]
                        if recipe.RecipeName == tool_recipe {
                            is_recipe = true
                        }
                    }
                    if strings.Contains(string(line), "__TOOLS_VERSION__="+config.Config.ToolsVersion) {
                        is_updated = true
                    }
                }

                if is_tool {
                    status = ui.GreenDk(ui.B("Recipe '"+tool_recipe+"' installed"))
                }
                if is_recipe {
                    status = ui.OrangeBt(ui.B("outdated"))
                }
                if is_updated && is_recipe {
                    status = ui.GreenBt(ui.B("installed"))
                }
            }
        }
    }

    return status
}
