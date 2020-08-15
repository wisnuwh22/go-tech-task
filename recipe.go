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

type recipes []recipe

func readRecipeJsonFile() recipes {
	bs, err := ioutil.ReadFile("recipess/data.json")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(0)
	}

	r := recipes{}
	err = json.Unmarshal([]byte(bs), &r)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(0)
	}

	return r

}
