// Package main contains main program execution
//
// Module main contains main program execution
package main

import (
	"os"

	"github.com/ga4gh/htsget-refserver-utils/internal/htsrunners"
)

// main execution of the program
func main() {
	os.Exit(run(os.Args[1:])) // strip 'htsget-refserver-utils' from passed args
}

// run convenience function, passes exit code generated from program to main
func run(args []string) int {

	// program requires a subcommand, if one isn't specified print the help
	// message
	if len(args) < 1 {
		htsrunners.Help()
		return 1
	}

	// execute the correct subcommand based on what is specified on the command
	// line. show help message if there was no matching subcommand
	subcommand := args[0]
	passedArgs := args[1:]
	switch subcommand {
	case "modify-sam":
		return htsrunners.ModifySam(passedArgs, os.Stdin)
	case "help":
		return htsrunners.Help()
	default:
		htsrunners.Help()
		return 1
	}
}
