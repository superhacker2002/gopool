package jsonconverter

import (
	"day01/internal/dbreaders/jsonreader"
	"day01/internal/entity"
	"encoding/json"
)

type JsonConverter struct{}

func (j JsonConverter) Convert(recipes entity.CakeRecipes) (string, error) {
	var outputJsonRecipes jsonreader.CakesRecipes
	for _, entityRecipe := range recipes.Recipes {
		var outputJsonRecipe jsonreader.Recipe
		for _, entityIngredient := range entityRecipe.Ingredients {
			var outputJsonIngredient = jsonreader.Ingredient{
				Name:  entityIngredient.Name,
				Count: entityIngredient.Count,
				Unit:  entityIngredient.Unit,
			}
			outputJsonRecipe.Ingredients = append(outputJsonRecipe.Ingredients, outputJsonIngredient)
		}
		outputJsonRecipe.Name = entityRecipe.Name
		outputJsonRecipe.Time = entityRecipe.Time
		outputJsonRecipes.Recipes = append(outputJsonRecipes.Recipes, outputJsonRecipe)
	}
	jsonOutputStr, err := json.MarshalIndent(outputJsonRecipes, "", "    ")

	return string(jsonOutputStr), err
}
