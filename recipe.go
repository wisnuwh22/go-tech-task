package main

type Recipe struct {
	Title       string   `json:"Title"`
	Ingredients []string `json:"Ingredients"`
}

type Recipes []Recipe
