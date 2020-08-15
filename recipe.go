package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type recipe struct {
	Title       string   `json:"Title"`
	Ingredients []string `json:"Ingredients"`
}

type recipes struct {
	Recipes []recipe `json:"recipes"`
}

func readRecipeJsonFile() recipes {
	bs, err := ioutil.ReadFile("recipes/data.json")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	r := recipes{}
	err = json.Unmarshal([]byte(bs), &r)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	return r

}
