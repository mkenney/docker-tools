/*
Package recipes defines data structures and methods for managing the tool recipe
database.

This portion of the recipe package defines the structure of the recipe
collection and provides methods for managing the collection
*/
package recipes

import (
	"io/ioutil"
	"lib/config"
	"os"
	"path"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

/*
Recipes is an array of recipie pointers containing all tool recipes available to
the system
*/
type Recipes []*Recipe

/*
New initializes and returns a pointer to a new instance of Recipies
*/
func New() *Recipes {
	recipes := make(Recipes, 0)
	return &recipes
}

/*
GetRecipe returns a specified recipe
*/
func (rcps Recipes) GetRecipe(recipeName, recipeSource string) (retval *Recipe) {
	for _, recipe := range rcps {
		if recipe.RecipeName == recipeName {
			if recipe.Source == recipeSource {
				retval = recipe
				break
			} else if "" == recipeSource {
				retval = recipe
				break
			}
		}
	}
	return
}

/*
HasRecipe returns whether a specified recipe exists in the index
*/
func (rcps Recipes) HasRecipe(recipeName, recipeSource string) (retval bool) {
	for _, recipe := range rcps {
		if recipe.RecipeName == recipeName {
			if recipe.Source == recipeSource {
				retval = true
				break
			} else if "" == recipeSource {
				retval = true
				break
			}
		}
	}
	return
}

/*
Load all recipes defined in a specified file

Recipe files are CSV files with the following structure:

     0: RecipeName
     1: ToolName
     2: ToolPrefix
     3: ToolTemplate
     4: DockerImage
     5: DockerTag
     6: ContainerVolumes
     7: ContainerEnv
     8: ContainerEntrypoint
     9: ContainerCmd
    10: DockerOptions
    11: RecipeNotes
*/
func (rcps *Recipes) Load(recipeFile string) *Recipes {

	if _, err := os.Stat(recipeFile); os.IsNotExist(err) {
		glog.Fatalf("File not found '%s': %s", recipeFile, err)

	} else {

		// Load the file intoa byte array and unmarshal a Yaml representation
		fileBytes, err := ioutil.ReadFile(recipeFile)
		if nil != err {
			glog.Fatalf("Error reading recipe file '%v'", recipeFile)
		}


		// Merge the loaded recipies into the global object
		var fileRecipes Recipes
		yaml.Unmarshal(fileBytes, &fileRecipes)
		for _, tmpRecipe := range fileRecipes {
			tmpRecipe.Source = path.Base(recipeFile)[0:len(path.Base(recipeFile)) - len(path.Ext(recipeFile))]
			rcps = rcps.SetRecipe(tmpRecipe)
		}
	}

	return rcps
}

/*
Save saves all the stuff
*/
func (rcps *Recipes) Save() *Recipes {
	for _, outfile := range []string{config.DockerToolsRegistry, config.ConfPath+"/recipes.yml"} {
		saveSet := make(Recipes, 0)
		for _, recipe := range (*rcps) {
			if recipe.Source == outfile {
				saveSet = append(saveSet, recipe)
			}
		}

		yamlBytes, err := yaml.Marshal(saveSet)
		if nil != err {glog.Fatalf("Unable to serialize recipe data: %s", err)}

		err = ioutil.WriteFile(outfile, yamlBytes, 0644)
		if nil != err {
			glog.Fatalf("Error writing recipe data: %s", err)
		}
	}
	return rcps
}

/*
SetRecipe will add a recipe to the index or update an existing recipe
*/
func (rcps Recipes) SetRecipe(recipe *Recipe) *Recipes {
	var updated bool

	for key, curRecipe := range rcps {
		if curRecipe.RecipeName == recipe.RecipeName && curRecipe.Source == recipe.Source {
			rcps[key] = recipe
			updated = true
		}
	}
	if !updated {
		rcps = append(rcps, recipe)
	}

	return &rcps
}
