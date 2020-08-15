package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ingredientPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	in := readIngredientJsonFile()
	json.NewEncoder(w).Encode(in)
}

func recipePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	re := readRecipeJsonFile()
	json.NewEncoder(w).Encode(re)
}

func handleRequest() {
	http.HandleFunc("/ingredient", ingredientPage)
	http.HandleFunc("/recipe", recipePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
