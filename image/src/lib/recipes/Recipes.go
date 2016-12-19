package recipes

import (
	"encoding/csv"
	"fmt"
	"glog"
	"io"
	"os"
	"path"
)

/*
Recipes is an array of recipie pointers containing all tool recipes available to
the system
*/
type Recipes [1000]*Recipe

/*
New initializes and returns a pointer to a new instance of Recipies
*/
func New() *Recipes {
	recipes := new(Recipes)
	return recipes
}

/*
HasRecipe returns whether a specified recipe exists in the index
*/
func (rcps *Recipes) HasRecipe(recipeName, recipeSource string) (retval bool) {
	for _, recipe := range rcps {
		if recipe.RecipeName == recipeName {
			if recipe.RecipeSource == recipeSource {
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
SetRecipe will add a recipe to the index or update an existing recipe
*/
func (rcps *Recipes) SetRecipe(recipeName, recipeSource string, recipe *Recipe) *Recipes {
	var updated bool
	var key int
	var curRecipe *Recipe

	for key, curRecipe = range rcps {
		if curRecipe.RecipeName == recipeName && curRecipe.RecipeSource == recipeSource {
			rcps[key] = recipe
			updated = true
			break
		}
	}
	if !updated {
		fmt.Printf("Last key: %v", key)
	}
	return rcps
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
func (rcps *Recipes) Load(recipeFile string) (reterr error) {

	if _, err := os.Stat(recipeFile); os.IsNotExist(err) {
		reterr = err
		glog.Fatalf("File not found '%v'", recipeFile)

	} else {
		file, err := os.Open(recipeFile)
		if nil != err {
			reterr = fmt.Errorf("Error opening recipe file '%v'", recipeFile)
			glog.Fatalf("Error opening recipe file '%v'", recipeFile)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		recipeidx := 0
		lineidx := 0
		for {
			line, err := reader.Read()
			if io.EOF == err {
				break
			} else if nil != err {
				reterr = fmt.Errorf("Error reading recipe file '%v': %v", recipeFile, err)
				glog.Fatalf("Error reading recipe file '%v': %v", recipeFile, err)
			}
			if 0 < lineidx { // Skip the header row
				record := make([]string, 13)
				for k, v := range line {
					record[k] = v
				}
				record[12] = path.Base(recipeFile) // Append the recipe source
				recipe := NewRecipe(record)
				rcps[recipeidx] = recipe
				recipeidx++
			}
			lineidx++
		}
	}

	return reterr
}
