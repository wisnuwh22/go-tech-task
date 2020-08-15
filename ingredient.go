package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ingredient struct {
	Title      string `json:"title"`
	BestBefore string `json:"best-before"`
	UseBy      string `json:"use-by"`
}

type ingredients struct {
	Ingredient []ingredient `json:"ingredients"`
}

func readIngredientJsonFile() ingredients {
	bs, err := ioutil.ReadFile("ingredients/data.json")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	i := ingredients{}
	err = json.Unmarshal([]byte(bs), &i)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	return i

}
