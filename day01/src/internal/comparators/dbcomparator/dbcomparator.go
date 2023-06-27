package dbcomparator

import (
	"day01/internal/entity"
	"fmt"
)

const (
	oldDB = 1
	newDB = 2
)

func Compare(old entity.CakeRecipes, new entity.CakeRecipes) string {
	uniqueRecipes := make(map[string]int)
	recipeInfo := ""
	additionalInfo := ""

	for _, recipe := range old.Recipes {
		uniqueRecipes[recipe.Name] = oldDB
	}
	for _, recipe := range new.Recipes {
		additionalInfo += addRecipe(uniqueRecipes, old.Recipe(recipe.Name), recipe)
	}

	for key, val := range uniqueRecipes {
		if val == oldDB {
			recipeInfo += fmt.Sprintf("REMOVED cake %q\n", key)
		}
		if val == newDB {
			recipeInfo += fmt.Sprintf("ADDED cake %q\n", key)
		}
	}

	return recipeInfo + additionalInfo
}

func addRecipe(uniqueRecipes map[string]int, old entity.Recipe, new entity.Recipe) string {
	outputStr := ""
	if _, ok := uniqueRecipes[new.Name]; ok {
		outputStr += addStoveTime(old.Time, new.Time, new.Name)
		outputStr += addIngredients(old, new, new.Name)
		delete(uniqueRecipes, new.Name)
	} else {
		uniqueRecipes[new.Name] = newDB
	}
	return outputStr
}

func addStoveTime(old string, new string, cakeName string) string {
	if old != new {
		return fmt.Sprintf("CHANGED cooking time for cake %q - %q instead of %q\n", cakeName, new, old)
	}
	return ""
}

func addIngredients(old entity.Recipe, new entity.Recipe, cakeName string) string {
	uniqueIngredients := make(map[string]int)
	ingredientInfo := ""
	additionalInfo := ""

	for _, ingredient := range old.Ingredients {
		uniqueIngredients[ingredient.Name] = oldDB
	}

	for _, ingredient := range new.Ingredients {
		additionalInfo += checkIngredient(uniqueIngredients, old.Ingredient(ingredient.Name), ingredient, cakeName)
	}

	for key, val := range uniqueIngredients {
		if val == oldDB {
			ingredientInfo += fmt.Sprintf("REMOVED ingredient %q for cake %q\n", key, cakeName)
		}
		if val == newDB {
			ingredientInfo += fmt.Sprintf("ADDED ingredient %q cake %q\n", key, cakeName)
		}
	}

	return ingredientInfo + additionalInfo
}

func checkIngredient(uniqueIngredients map[string]int, old entity.Ingredient,
	new entity.Ingredient, cakeName string) string {
	outputStr := ""
	if _, ok := uniqueIngredients[old.Name]; ok {
		outputStr += addIngredientCount(old, new, cakeName)
		outputStr += addIngredientUnit(old, new, cakeName)
		delete(uniqueIngredients, old.Name)
	} else {
		uniqueIngredients[old.Name] = newDB
	}

	return outputStr
}

func addIngredientCount(old entity.Ingredient, new entity.Ingredient, cakeName string) string {
	if old.Count != new.Count {
		return fmt.Sprintf("CHANGED unit count for ingredient %q for cake %q - "+
			"%q instead of %q\n", old.Name, cakeName, new.Count, old.Count)
	}

	return ""
}

func addIngredientUnit(old entity.Ingredient, new entity.Ingredient, cakeName string) string {
	if old.Unit == new.Unit {
		return ""
	}

	if old.Unit == "" && new.Unit != "" {
		return fmt.Sprintf("ADDED unit %q for ingredient %q for cake %q\n",
			old.Unit, old.Name, cakeName)
	}

	if old.Unit != "" && new.Unit == "" {
		return fmt.Sprintf("REMOVED unit %q for ingredient %q for cake %q\n",
			old.Unit, old.Name, cakeName)
	}

	if old.Unit != new.Unit {
		return fmt.Sprintf("CHANGED unit for ingredient %q for cake %q - %q instead of %q\n",
			old.Name, cakeName, new.Unit, old.Unit)
	}

	return ""
}
