package main

import "testing"

func TestCreateIngredient(t *testing.T) {
	line := "Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8"
	result := createIngredient(line)
	want := Ingredient{"Butterscotch:", -1, -2, 6, 3, 8}
	if result != want {
		t.Fatalf("Expected %v, but got %v", want, result)
	}
}

func TestGetScore(t *testing.T) {
	ingredients := map[string]Ingredient{
		"Butterscotch:": {"Butterscotch:", -1, -2, 6, 3, 8},
		"Cinnamon:":     {"Butterscotch:", 2, 3, -2, -1, 3},
	}
	recipe := map[string]int{
		"Butterscotch:": 44,
		"Cinnamon:":     56,
	}
	result := getScore(recipe, ingredients)
	want := 62842880
	if result != want {
		t.Fatalf("Expected %v, but got %v", want, result)
	}
}
