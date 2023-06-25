package jsonreader

import (
	"day01/entity"
	"encoding/json"
	"io"
)

type JsonReader struct{}

type cakesRecipes struct {
	Recipes []recipe `json:"cake"`
}

type recipe struct {
	Name        string `json:"name"`
	Time        string `json:"time"`
	Ingredients []ingredient
}

type ingredient struct {
	Name  string `json:"ingredient_name"`
	Count string `json:"ingredient_count"`
	Unit  string `json:"ingredient_unit,omitempty"`
}

func (j JsonReader) Read(reader io.Reader) (entity.CakeRecipes, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return entity.CakeRecipes{}, err
	}

	cr := cakesRecipes{}

	err = json.Unmarshal(data, &cr)
	if err != nil {
		return entity.CakeRecipes{}, err
	}

	return recipeToEntity(cr), nil
}

func recipeToEntity(recipes cakesRecipes) entity.CakeRecipes {
	var outputCakes entity.CakeRecipes
	for _, jsonRecipe := range recipes.Recipes {
		var entityIngredients []entity.Ingredient
		for _, jsonIngredient := range jsonRecipe.Ingredients {
			entityIngredient := entity.Ingredient{
				Name:  jsonIngredient.Name,
				Count: jsonIngredient.Count,
				Unit:  jsonIngredient.Unit,
			}
			entityIngredients = append(entityIngredients, entityIngredient)
		}
		entityRecipe := entity.Recipe{
			Name:        jsonRecipe.Name,
			Time:        jsonRecipe.Time,
			Ingredients: entityIngredients,
		}
		outputCakes.Recipes = append(outputCakes.Recipes, entityRecipe)
	}

	return outputCakes
}
