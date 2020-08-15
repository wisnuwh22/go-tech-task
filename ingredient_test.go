package main

import (
	"reflect"
	"testing"
)

func TestReadIngredientJsonFile(t *testing.T) {
	testCases := map[string]struct {
		index int
		title string
	}{
		"TestIndex0": {index: 0, title: "Ham"},
		"TestIndex1": {index: 1, title: "Cheese"},
		"TestIndex2": {index: 2, title: "Bread"},
		"TestIndex3": {index: 3, title: "Butter"},
		"TestIndex4": {index: 4, title: "Bacon"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			is := readIngredientJSONFile()
			i := is.Ingredients
			firstCut := i[:tc.index+1]
			item := firstCut[len(firstCut)-1]
			if !reflect.DeepEqual(item.Title, tc.title) {
				t.Errorf("Expected %v Got %v", tc.title, item.Title)
			}
		})
	}
}
