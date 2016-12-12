package recipes

import ()

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
    ContainerVolumes    string // Volumes to mount
    ContainerEnv        string // Environment variables to pass
    ContainerEntrypoint string // Container entrypoint
    ContainerCmd        string // Tool/container commands
    DockerOptions       string // Additional `docker run` cli arguments
    RecipeNotes         string // Any notes about the recipe
    RecipeSource        string // The name of the sourcefile, either 'recipes' or 'registry'
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
    recipe_data[6]  = this.ContainerVolumes    // Volumes to mount
    recipe_data[7]  = this.ContainerEnv        // Environment variables to pass
    recipe_data[8]  = this.ContainerEntrypoint // Container entrypoint
    recipe_data[9]  = this.ContainerCmd        // Tool/container commands
    recipe_data[10] = this.DockerOptions       // Additional `docker run` cli arguments
    recipe_data[11] = this.RecipeNotes         // Any notes about the recipe
    recipe_data[12] = this.RecipeSource        // The name of the sourcefile, either 'recipes' or 'registry'
    return recipe_data
}

/**
 * Generate a Recipe model from a data array
 *
 * @param  {[]string} recipe_data
 * @return {*Recipe}
 */
func createRecipeFromData(recipe_data []string) (*Recipe) {
    recipe := new(Recipe)
    recipe.RecipeName          = recipe_data[0]  // Recipe name
    recipe.ToolName            = recipe_data[1]  // Tool name
    recipe.ToolPrefix          = recipe_data[2]  // Tool install prefix
    recipe.ToolTemplate        = recipe_data[3]  // Tool template name (tool, service (daemons))
    recipe.DockerImage         = recipe_data[4]  // Docker image name
    recipe.DockerTag           = recipe_data[5]  // Docker image tag
    recipe.ContainerVolumes    = recipe_data[6]  // Volumes to mount
    recipe.ContainerEnv        = recipe_data[7]  // Environment variables to pass
    recipe.ContainerEntrypoint = recipe_data[8]  // Container entrypoint
    recipe.ContainerCmd        = recipe_data[9]  // Tool/container commands
    recipe.DockerOptions       = recipe_data[10] // Additional `docker run` cli arguments
    recipe.RecipeNotes         = recipe_data[11] // Any notes about the recipe
    recipe.RecipeSource        = recipe_data[12] // The name of the sourcefile, either 'recipes' or 'registry'
    return recipe
}
