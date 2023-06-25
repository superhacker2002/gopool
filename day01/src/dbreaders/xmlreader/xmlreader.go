package xmlreader

import (
	"day01/entity"
	"encoding/xml"
	"io"
)

type XmlReader struct{}

type CakesRecipes struct {
	XMLName xml.Name `xml:"recipes"`
	Recipes []Recipe `xml:"cake"`
}

type Recipe struct {
	Name        string      `xml:"name"`
	Time        string      `xml:"stovetime"`
	Ingredients Ingredients `xml:"ingredients"`
}

type Ingredients struct {
	XMLName xml.Name `xml:"ingredients"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Name    string   `xml:"itemname"`
	Count   string   `xml:"itemcount"`
	Unit    string   `xml:"itemunit,omitempty"`
}

func (x XmlReader) Read(reader io.Reader) (entity.CakeRecipes, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return entity.CakeRecipes{}, err
	}

	cr := CakesRecipes{}

	err = xml.Unmarshal(data, &cr)
	if err != nil {
		return entity.CakeRecipes{}, err
	}

	return recipeToEntity(cr), nil
}

func recipeToEntity(recipes CakesRecipes) entity.CakeRecipes {
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
