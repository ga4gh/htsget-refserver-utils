// Package htsutils contains miscellaneous, static util methods to be used
// throughout the program
//
// Module utils contains static utility methods
package htsutils

import (
	"github.com/golang-collections/collections/set"
)

// CreateInterfaceList converts all strings within a list to empty interfaces,
// returning them as a list
func CreateInterfaceList(old []string) []interface{} {

	new := make([]interface{}, len(old))
	for i, item := range old {
		new[i] = item
	}
	return new
}

// CreateSet creates a set from a list of strings, so that set operations can
// be performed on them
func CreateSet(arr []string) *set.Set {
	interfaces := CreateInterfaceList(arr)
	return set.New(interfaces...)
}
