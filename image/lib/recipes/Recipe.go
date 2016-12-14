package recipes

import (
    "lib/config"
    "flag" // Used by glog
)

/**
 * A recipe is a set of properties that define a tool or service management
 * script
 */
type Recipe struct {
    RecipeName          string // Recipe name
    ToolName            string // Tool name
    ToolPrefix          string // Tool install prefix
    ToolTemplate        string // Tool template name (tool, service (daemons))
    DockerImage         string // Docker image name
    DockerTag           string // Docker image tag
    ContainerVolumes    []byte // Volumes to mount
    ContainerEnv        []byte // Environment variables to pass
    ContainerEntrypoint string // Container entrypoint
    ContainerCmd        string // Tool/container commands
    DockerOptions       string // Additional `docker run` cli arguments
    RecipeNotes         string // Any notes about the recipe
    RecipeSource        string // The name of the sourcefile, either 'recipes' or 'registry'
}

/**
 * Generate a Recipe model from a data array
 *
 * @param  {[]string} recipe_data
 * @return {*Recipe}
 */
func NewRecipe(recipe_data []string) (*Recipe) {
    flag.Parse() // required by glog

    recipe := new(Recipe)

    // Recipe name
    recipe.RecipeName = recipe_data[0]

    // Tool name
    recipe.ToolName = recipe_data[1]

    // Tool install prefix
    recipe.ToolPrefix = recipe_data[2]
    if "" == recipe.ToolPrefix {
        recipe.ToolPrefix = config.Config.DefaultPrefix
    }

    // Tool template name (tool, service (daemons))
    recipe.ToolTemplate = recipe_data[3]

    // Docker image name
    recipe.DockerImage = recipe_data[4]

    // Docker image tag
    recipe.DockerTag = recipe_data[5]
    if "" == recipe.DockerTag {
        recipe.DockerTag = "latest"
    }

    // Volumes to mount
    recipe.ContainerVolumes = []byte(recipe_data[6])

    // Environment variables to pass
    recipe.ContainerEnv = []byte(recipe_data[7])

    // Container entrypoint
    recipe.ContainerEntrypoint = recipe_data[8]

    // Tool/container commands
    recipe.ContainerCmd = recipe_data[9]

    // Additional `docker run` cli arguments
    recipe.DockerOptions = recipe_data[10]

    // Any notes about the recipe
    recipe.RecipeNotes = recipe_data[11]

    // The name of the sourcefile, either 'recipes' or 'registry'
    recipe.RecipeSource = recipe_data[12]

    return recipe
}


/**
 * Output a recipe model as an array
 *
 * @return {[13]string}
 */
func (this *Recipe) ToArray() [13]string {
    var recipe_data [13]string
    recipe_data[0]  = this.RecipeName          // Recipe name
    recipe_data[1]  = this.ToolName            // Tool name
    recipe_data[2]  = this.ToolPrefix          // Tool install prefix
    recipe_data[3]  = this.ToolTemplate        // Tool template name (tool, service (daemons))
    recipe_data[4]  = this.DockerImage         // Docker image name
    recipe_data[5]  = this.DockerTag           // Docker image tag
    recipe_data[6]  = string(this.ContainerVolumes)    // Volumes to mount
    recipe_data[7]  = string(this.ContainerEnv)        // Environment variables to pass
    recipe_data[8]  = this.ContainerEntrypoint // Container entrypoint
    recipe_data[9]  = this.ContainerCmd        // Tool/container commands
    recipe_data[10] = this.DockerOptions       // Additional `docker run` cli arguments
    recipe_data[11] = this.RecipeNotes         // Any notes about the recipe
    recipe_data[12] = this.RecipeSource        // The name of the sourcefile, either 'recipes' or 'registry'
    return recipe_data
}
