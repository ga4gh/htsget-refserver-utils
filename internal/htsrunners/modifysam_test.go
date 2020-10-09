// Package htsrunners contains cli subcommands
//
// Module modifysam_test tests modifysam
package htsrunners

import (
	"crypto/md5"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

// modifySamTC test cases for ModifySam
var modifySamTC = []struct {
	args     []string
	expError bool
	filename string
}{
	// error case
	{
		[]string{"-fields", "QNAME,SEQ,BAD,QUAL"},
		true,
		"modify-sam.00.sam",
	},
	// no modification
	{
		[]string{},
		false,
		"modify-sam.00.sam",
	},
	// field modification only
	{
		[]string{"-fields", "QNAME,RNAME,MAPQ,RNEXT,TLEN,QUAL"},
		false,
		"modify-sam.01.sam",
	},
	{
		[]string{"-fields", "FLAG,POS,CIGAR,PNEXT,SEQ"},
		false,
		"modify-sam.02.sam",
	},
	// tags modification only
	{
		[]string{"-tags", "NH"},
		false,
		"modify-sam.03.sam",
	},
	{
		[]string{"-tags", "HI,NM,MD"},
		false,
		"modify-sam.04.sam",
	},
	{
		[]string{"-tags", "NONE"},
		false,
		"modify-sam.05.sam",
	},
	// notags modification only
	{
		[]string{"-notags", "HI,NM"},
		false,
		"modify-sam.06.sam",
	},
	{
		[]string{"-notags", "MD"},
		false,
		"modify-sam.07.sam",
	},
	// field, tags, notags modification
	{
		[]string{"-fields", "QNAME,FLAG,RNAME,POS", "-tags", "NH,HI", "-notags", "MD"},
		false,
		"modify-sam.08.sam",
	},
	{
		[]string{"-fields", "MAPQ,CIGAR,RNEXT,PNEXT", "-notags", "NH,HI"},
		false,
		"modify-sam.09.sam",
	},
	{
		[]string{"-fields", "TLEN,SEQ,QUAL", "-tags", "NONE"},
		false,
		"modify-sam.10.sam",
	},
	{
		[]string{"-fields", "CIGAR,QNAME,TLEN,SEQ,QUAL", "-tags", "MD"},
		false,
		"modify-sam.11.sam",
	},
}

// TestModifySam tests function ModifySam
func TestModifySam(t *testing.T) {

	for _, tc := range modifySamTC {
		// unset flag values between cases
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		// declare input/output files
		dataDir := "../../data/test"
		inputFp := dataDir + "/input/modify-sam.sam"
		outputFp := dataDir + "/output/" + tc.filename

		// declare wrapper function to capture stdout
		stdinReader, _ := os.Open(inputFp)

		if tc.expError {
			code := ModifySam(tc.args, stdinReader)
			assert.Equal(t, 1, code)
		} else {

			modifySamWrapper := func() {
				ModifySam(tc.args, stdinReader)
			}

			// load expected stdout
			stdoutReader, _ := os.Open(outputFp)
			expectedStdout, _ := ioutil.ReadAll(stdoutReader)
			expectedMD5 := md5.Sum(expectedStdout)

			// run function, capture stdout, and compare to expected
			actualStdout := capturer.CaptureStdout(modifySamWrapper)
			actualMD5 := md5.Sum([]byte(actualStdout))
			assert.Equal(t, expectedMD5, actualMD5)
		}
	}
}
