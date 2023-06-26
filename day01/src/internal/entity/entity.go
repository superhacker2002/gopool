package entity

type CakeRecipes struct {
	Recipes []Recipe
}

type Recipe struct {
	Name        string
	Time        string
	Ingredients []Ingredient
}

type Ingredient struct {
	Name  string
	Count string
	Unit  string
}

func (c CakeRecipes) Recipe(name string) Recipe {
	for _, recipe := range c.Recipes {
		if recipe.Name == name {
			return recipe
		}
	}
	return Recipe{}
}

func (r Recipe) Ingredient(name string) Ingredient {
	for _, ingredient := range r.Ingredients {
		if ingredient.Name == name {
			return ingredient
		}
	}
	return Ingredient{}
}
