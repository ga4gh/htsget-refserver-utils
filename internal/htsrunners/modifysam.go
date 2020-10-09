// Package htsrunners contains cli subcommands
//
// Module modifysam contains the modify-sam subcommand, in which a SAM file is
// streamed from stdin, custom fields and tags are included/excluded and streamed
// to stdout
package htsrunners

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/ga4gh/htsget-refserver-utils/internal/htsformats"
)

// samRecordCustomEmit convenience method to emit a single SAM alignment/record
// based on how the samRecordEmitter has been configured
func samRecordCustomEmit(samRecordEmitter *htsformats.SamRecordEmitter, text string) {
	samRecord := htsformats.NewSamRecord(text)
	customEmit := samRecordEmitter.CustomEmit(samRecord)
	fmt.Println(customEmit)
}

// ModifySam runner for 'modify-sam' subcommand. Streams a SAM file from stdin,
// performs custom field/tag inclusion, and streams to stdout
func ModifySam(args []string, reader io.Reader) int {

	// parses cli args
	fieldsPtr := flag.String("fields", "", "comma-delimited list of fields to include in output SAM")
	tagsPtr := flag.String("tags", "", "comma-delimited list of tags to include in output SAM")
	notagsPtr := flag.String("notags", "", "comma-delimited list of tags to exclude from output SAM")
	flag.CommandLine.Parse(args)

	// configure the SamRecordEmitter
	samRecordEmitter, err := htsformats.NewSamRecordEmitter(*fieldsPtr, *tagsPtr, *notagsPtr)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return 1
	}

	// iterates over each line in the SAM
	header := true
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if header {
			// first, header lines are streamed from stdin to stdout without
			// modification. the program looks for the first non-header line,
			// at which point the SamRecord lines are emitted according to custom
			// rules
			if strings.HasPrefix(text, "@") {
				fmt.Println(text)
			} else {
				header = false
				samRecordCustomEmit(samRecordEmitter, text)
			}
		} else {
			samRecordCustomEmit(samRecordEmitter, text)
		}
	}
	return 0
}
