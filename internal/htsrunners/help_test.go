// Package htsrunners contains cli subcommands
//
// Module help_test tests module help
package htsrunners

import (
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

// TestHelp tests Help function
func TestHelp(t *testing.T) {
	helpWrapper := func() {
		Help()
	}

	stdout := capturer.CaptureStdout(helpWrapper)
	assert.Equal(t, stdout, helpMessage)
}
