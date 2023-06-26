package xmlconverter

import (
	"day01/internal/dbreaders/xmlreader"
	"day01/internal/entity"
	"encoding/xml"
)

type XmlConverter struct{}

func (x XmlConverter) Convert(recipes entity.CakeRecipes) (string, error) {
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
