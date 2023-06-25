package xmlreader

import (
	"day01/entity"
	"encoding/xml"
	"io"
)

type XmlReader struct{}

type cakesRecipes struct {
	Recipes []recipe `xml:"cake"`
}

type recipe struct {
	Name        string      `xml:"name"`
	Time        string      `xml:"stovetime"`
	Ingredients ingredients `xml:"ingredients"`
}

type ingredients struct {
	XMLName xml.Name `xml:"ingredients"`
	Items   []item   `xml:"item"`
}

type item struct {
	XMLName xml.Name `xml:"item"`
	Name    string   `xml:"itemname"`
	Count   string   `xml:"itemcount"`
	Unit    string   `xml:"itemunit"`
}

func (x XmlReader) Read(reader io.Reader) (entity.CakeRecipes, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return entity.CakeRecipes{}, err
	}

	cr := cakesRecipes{}

	err = xml.Unmarshal(data, &cr)
	if err != nil {
		return entity.CakeRecipes{}, err
	}

	return recipeToEntity(cr), nil
}

func recipeToEntity(recipes cakesRecipes) entity.CakeRecipes {
	var outputCakes entity.CakeRecipes
	for _, xmlRecipe := range recipes.Recipes {
		var entityIngredients []entity.Ingredient
		for _, jsonIngredient := range xmlRecipe.Ingredients.Items {
			entityIngredient := entity.Ingredient{
				Name:  jsonIngredient.Name,
				Count: jsonIngredient.Count,
				Unit:  jsonIngredient.Unit,
			}
			entityIngredients = append(entityIngredients, entityIngredient)
		}
		entityRecipe := entity.Recipe{
			Name:        xmlRecipe.Name,
			Time:        xmlRecipe.Time,
			Ingredients: entityIngredients,
		}
		outputCakes.Recipes = append(outputCakes.Recipes, entityRecipe)
	}

	return outputCakes
}
