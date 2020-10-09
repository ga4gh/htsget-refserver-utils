// Package main contains main program execution
//
// Module main_test tests main
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// runTC test cases for run
var runTC = []struct {
	args        []string
	expExitCode int
}{
	{
		[]string{},
		1,
	},
	{
		[]string{"modify-sam"},
		0,
	},
	{
		[]string{"help"},
		0,
	},
	{
		[]string{"non-command"},
		1,
	},
}

// TestRun tests run function
func TestRun(t *testing.T) {
	for _, tc := range runTC {
		actualExitCode := run(tc.args)
		assert.Equal(t, tc.expExitCode, actualExitCode)
	}
}
