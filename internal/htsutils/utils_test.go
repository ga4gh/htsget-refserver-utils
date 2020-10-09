// Package htsutils contains miscellaneous, static util methods to be used
// throughout the program
//
// Module utils_test tests utils
package htsutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// utilsCreateInterfaceListTC test cases for CreateInterfaceList
var utilsCreateInterfaceListTC = []struct {
	input []string
	exp   []interface{}
}{
	{
		[]string{"MD"},
		[]interface{}{"MD"},
	},
	{
		[]string{"HI", "MD"},
		[]interface{}{"HI", "MD"},
	},
	{
		[]string{"MD", "NH", "NM", "HI"},
		[]interface{}{"MD", "NH", "NM", "HI"},
	},
}

// utilsCreateSetTC test cases for CreateSet
var utilsCreateSetTC = []struct {
	input []string
}{
	{[]string{"MD"}},
	{[]string{"HI", "MD"}},
	{[]string{"MD", "NH", "NM", "HI"}},
}

// TestUtilsCreateInterfaceList tests CreateInterfaceList function
func TestUtilsCreateInterfaceList(t *testing.T) {
	for _, tc := range utilsCreateInterfaceListTC {
		assert.Equal(t, tc.exp, CreateInterfaceList(tc.input))
	}
}

// TestUtilsCreateSet tests CreateSet function
func TestUtilsCreateSet(t *testing.T) {
	for _, tc := range utilsCreateSetTC {
		set := CreateSet(tc.input)
		for _, item := range tc.input {
			hasItem := set.Has(item)
			assert.True(t, hasItem)
		}
	}
}
