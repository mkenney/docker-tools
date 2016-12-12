package recipes

import (
    "encoding/csv"
    "errors"
    "fmt"
    "io"
    "log"
    "os"
    "path"
)

/**
 * Represents all tool recipes available to the system
 */
type Recipes [1000]*Recipe
//type Recipes struct  {
//    recipes map[string]*Recipe
//    files   map[string]bool // True if loaded, else false
//}

/**
 *
 */
func New() *Recipes {
    recipes := new(Recipes)
    return recipes
}

/**
 * See if a recipe exists in the index
 *
 * @param  {string} recipe_name
 * @param  {string} recipe_source   Either 'registry' or 'recipes'
 * @return {bool}   True if a matching recipe was found, else false
 */
func (this *Recipes) HasRecipe(recipe_name, recipe_source string) (bool) {
    ret_val := false
    for _, recipe := range this {
        if recipe.RecipeName == recipe_name {
            if recipe.RecipeSource == recipe_source {
                ret_val = true
                break
            } else if "" == recipe_source {
                ret_val = true
                break
            }
        }
    }
    return ret_val
}

/**
 * See if a recipe exists in the index
 *
 * @param  {string} recipe_name
 * @param  {string} recipe_source   Either 'registry' or 'recipes'
 * @return {bool}   True if a matching recipe was found, else false
 */
func (this *Recipes) SetRecipe(recipe_name, recipe_source string, recipe *Recipe) (*Recipes) {
    var updated bool
    var key int
    var cur_recipe *Recipe

    for key, cur_recipe = range this {
        if cur_recipe.RecipeName == recipe_name && cur_recipe.RecipeSource == recipe_source {
            this[key] = recipe
            updated = true
            break
        }
    }
    if !updated {
        fmt.Printf("Last key: %v", key)
    }
    return this
}

/**
 * Load all recipes defined in a specified file
 *
 * Recipe files are headerless CSV files with the following structure:
 *
 *     0: RecipeName
 *     1: ToolName
 *     2: ToolPrefix
 *     3: ToolTemplate
 *     4: DockerImage
 *     5: DockerTag
 *     6: ContainerVolumes
 *     7: ContainerEnv
 *     8: ContainerEntrypoint
 *     9: ContainerCmd
 *    10: DockerOptions
 *    11: RecipeNotes
 *
 * @param  {string} recipe_file
 * @return {bool}   true on success, else false
 * @return {error}  error on error, else nil
 */
func (this *Recipes) Load(recipe_file string, ret_chan chan bool, err_chan chan error) {
    var ret_val bool
    var ret_err error

    if _, err := os.Stat(recipe_file); os.IsNotExist(err) {
        ret_val = false
        ret_err = err
        log.Fatalf("File not found '%v'", recipe_file)

    } else {
        file, err := os.Open(recipe_file)
        if nil != err {
            ret_val = false
            ret_err = errors.New(fmt.Sprintf("Error opening recipe file '%v'", recipe_file))
            log.Fatalf("Error opening recipe file '%v'", recipe_file)
        }
        defer file.Close()

        reader := csv.NewReader(file)
        recipe_idx := 0
        for {
            line, err := reader.Read()
            if io.EOF == err  {
                break
            } else if nil != err {
                ret_val = false
                ret_err = errors.New(fmt.Sprintf("Error reading recipe file '%v': %v", recipe_file, err))
                log.Fatalf("Error reading recipe file '%v': %v", recipe_file, err)
            }

            record := make([]string, 13)
            for k, v := range line {record[k] = v}
            // Append the filename
            record[12] = path.Base(recipe_file)
            recipe := createRecipeFromData(record)

            this[recipe_idx] = recipe
            recipe_idx++
        }
    }

    ret_chan<- ret_val
    err_chan<- ret_err
}
