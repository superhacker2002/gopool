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
