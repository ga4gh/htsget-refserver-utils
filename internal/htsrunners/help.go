// Package htsrunners contains cli subcommands
//
// Module help prints program help message to stdout
package htsrunners

import "fmt"

// helpMessage message to be displayed on help
var helpMessage = `htsget-refserver-utils - cli app for streaming genomic data according to GA4GH htsget protocol

Usage:
htsget-refserver-utils <COMMAND> <ARG1> <ARG2> ...

Commands:
modify-sam	include/exclude fields and tags from SAM stdin stream
`

// Help prints command help message
func Help() int {
	fmt.Print(helpMessage)
	return 0
}
