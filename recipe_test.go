package main

import (
	"reflect"
	"testing"
)

func TestReadRecipeJsonFile(t *testing.T) {
	testCases := map[string]struct {
		index int
		title string
	}{
		"TestIndex0": {index: 0, title: "Ham and Cheese Toastie"},
		"TestIndex1": {index: 1, title: "Fry-up"},
		"TestIndex2": {index: 2, title: "Salad"},
		"TestIndex3": {index: 3, title: "Hotdog"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			rs := readRecipeJSONFile()
			r := rs.Recipes
			firstCut := r[:tc.index+1]
			item := firstCut[len(firstCut)-1]
			if !reflect.DeepEqual(item.Title, tc.title) {
				t.Errorf("Expected %v Got %v", tc.title, item.Title)
			}
		})
	}
}
