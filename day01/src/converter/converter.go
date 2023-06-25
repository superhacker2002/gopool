package converter

import (
	"day01/dbreaders/jsonreader"
	"day01/dbreaders/xmlreader"
	"day01/entity"
	"encoding/json"
	"encoding/xml"
)

func ToJson(recipes entity.CakeRecipes) (string, error) {
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

func ToXml(recipes entity.CakeRecipes) (string, error) {
	var outputXmlRecipes xmlreader.CakesRecipes
	for _, entityRecipe := range recipes.Recipes {
		var xmlIngredients xmlreader.Ingredients
		for _, entityIngredient := range entityRecipe.Ingredients {
			var xmlItem = xmlreader.Item{
				Name:  entityIngredient.Name,
				Count: entityIngredient.Count,
				Unit:  entityIngredient.Unit,
			}
			xmlIngredients.Items = append(xmlIngredients.Items, xmlItem)
		}
		var xmlRecipe = xmlreader.Recipe{
			Name:        entityRecipe.Name,
			Time:        entityRecipe.Time,
			Ingredients: xmlIngredients,
		}
		outputXmlRecipes.Recipes = append(outputXmlRecipes.Recipes, xmlRecipe)
	}
	xmlOutputStr, err := xml.MarshalIndent(outputXmlRecipes, "", "    ")

	return string(xmlOutputStr), err
}
