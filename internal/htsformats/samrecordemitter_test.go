// Package htsformats contains objects modeling entities in genomic file formats
// and associated behaviors
//
// Module samrecordemitter_test tests samrecordemitter
package htsformats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// newSamRecordEmitterTC test cases for NewSamRecordEmitter
var newSamRecordEmitterTC = []struct {
	fields, tags, notags string
	expError             bool
}{
	{"", "", "", false},
	{"QUAL,FOO", "", "", true},
	{"QUAL,SEQ", "HI", "HI", true},
}

// samRecordEmitterSetupFieldsTC test cases for setupFields
var samRecordEmitterSetupFieldsTC = []struct {
	fields           string
	expError         bool
	expEmitAllFields bool
	expFields        []bool
}{
	{"", false, true, nil},
	{"QNAME", false, false, []bool{true, false, false, false, false, false, false, false, false, false, false}},
	{"QNAME,NONFIELD,RNAME", true, false, nil},
	{"FLAG,SEQ,TLEN,CIGAR,POS,PNEXT", false, false, []bool{false, true, false, true, false, true, false, true, true, true, false}},
	{"QUAL,MAPQ,RNAME,RNEXT", false, false, []bool{false, false, true, false, true, false, true, false, false, false, true}},
	{"TLEN,QNAME,POS,SEQ,CIGAR", false, false, []bool{true, false, false, true, false, true, false, false, true, true, false}},
}

// samRecordEmitterSetupTagsNotagsTC test cases for setupTagsNotags
var samRecordEmitterSetupTagsNotagsTC = []struct {
	tags, notags     string
	expError         bool
	expEmitAllTags   bool
	expInclusionEmit bool
	expTags          []string
	expNotags        []string
}{
	{"", "", false, true, false, []string{}, []string{}},
	{"MD,HI", "", false, false, true, []string{"MD", "HI"}, []string{}},
	{"", "MD,HI", false, false, false, []string{}, []string{"MD", "HI"}},
	{"MD,HI", "NM,NI", false, false, true, []string{"MD", "HI"}, []string{"NM", "NI"}},
	{"NM,NI,MD", "HI", false, false, true, []string{"NM", "NI", "MD"}, []string{"HI"}},
	{"HI", "HI", true, false, true, nil, nil},
}

// samRecordEmitterCustomEmitTC test cases for CustomEmit
var samRecordEmitterCustomEmitTC = []struct {
	rawRecord, fields, tags, notags, exp string
}{
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"",
		"",
		"",
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"QNAME,FLAG,RNAME",
		"",
		"",
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t0\t255\t*\t*\t0\t0\t*\t*\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"QNAME,FLAG,RNAME",
		"NONE",
		"",
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t0\t255\t*\t*\t0\t0\t*\t*",
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"QNAME,FLAG,RNAME",
		"",
		"HI,MD",
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t0\t255\t*\t*\t0\t0\t*\t*\tNH:i:1\tNM:i:0",
	},
}

// samRecordEmitterCustomEmitFieldsTC test cases for customEmitFields
var samRecordEmitterCustomEmitFieldsTC = []struct {
	rawRecord, fields string
	exp               []string
}{
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"QNAME,FLAG,RNAME",
		[]string{"A00111:67:H3M5YDMXX:2:1182:16125:23813", "147", "ERCC-00171", "0", "255", "*", "*", "0", "0", "*", "*"},
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"CIGAR,RNEXT,PNEXT,TLEN",
		[]string{"*", "0", "*", "0", "255", "100M", "=", "1", "-379", "*", "*"},
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"SEQ,QUAL",
		[]string{"*", "0", "*", "0", "255", "*", "*", "0", "0", "AACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT", "FFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF"},
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"MAPQ,POS,RNAME",
		[]string{"*", "0", "ERCC-00171", "280", "255", "*", "*", "0", "0", "*", "*"},
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"QUAL,TLEN,RNEXT,MAPQ,RNAME,QNAME",
		[]string{"A00111:67:H3M5YDMXX:2:1182:16125:23813", "0", "ERCC-00171", "0", "255", "*", "=", "0", "-379", "*", "FFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF"},
	},
}

// samRecordEmitterCustomEmitTagsTC test cases for customEmitTags
var samRecordEmitterCustomEmitTagsTC = []struct {
	rawRecord, tags, notags string
	exp                     []string
}{
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"NH,HI",
		"",
		[]string{"NH:i:1", "HI:i:1"},
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"",
		"MD",
		[]string{"NH:i:1", "HI:i:1", "NM:i:0"},
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"NH",
		"HI",
		[]string{"NH:i:1"},
	},
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"",
		"NH,HI,NM,MD",
		[]string{},
	},
}

// TestNewSamRecordEmitter tests NewSamRecordEmitter function
func TestNewSamRecordEmitter(t *testing.T) {
	for _, tc := range newSamRecordEmitterTC {
		_, err := NewSamRecordEmitter(tc.fields, tc.tags, tc.notags)
		if tc.expError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

// TestSamRecordEmitterSetupFields tests setupFields function
func TestSamRecordEmitterSetupFields(t *testing.T) {
	for _, tc := range samRecordEmitterSetupFieldsTC {
		samRecordEmitter := new(SamRecordEmitter)
		err := samRecordEmitter.setupFields(tc.fields)
		if tc.expError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.expEmitAllFields, samRecordEmitter.emitAllFields)
			if !tc.expEmitAllFields {
				assert.Equal(t, tc.expFields, samRecordEmitter.fields)
			}
		}
	}
}

// TestSamRecordEmitterSetupTagsNotags tests setupTagsNotags function
func TestSamRecordEmitterSetupTagsNotags(t *testing.T) {
	for _, tc := range samRecordEmitterSetupTagsNotagsTC {
		samRecordEmitter := new(SamRecordEmitter)
		err := samRecordEmitter.setupTagsNotags(tc.tags, tc.notags)
		if tc.expError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.expEmitAllTags, samRecordEmitter.emitAllTags)
			assert.Equal(t, tc.expInclusionEmit, samRecordEmitter.inclusionEmit)
			assert.Equal(t, tc.expTags, samRecordEmitter.tags)
			assert.Equal(t, tc.expNotags, samRecordEmitter.notags)
		}
	}
}

// TestSamRecordEmitterCustomEmit tests CustomEmit function
func TestSamRecordEmitterCustomEmit(t *testing.T) {
	for _, tc := range samRecordEmitterCustomEmitTC {
		samRecordEmitter, _ := NewSamRecordEmitter(tc.fields, tc.tags, tc.notags)
		samRecord := NewSamRecord(tc.rawRecord)
		actual := samRecordEmitter.CustomEmit(samRecord)
		assert.Equal(t, tc.exp, actual)
	}
}

// TestSamRecordEmitterCustomEmitFields tests customEmitFields function
func TestSamRecordEmitterCustomEmitFields(t *testing.T) {
	for _, tc := range samRecordEmitterCustomEmitFieldsTC {
		samRecordEmitter, _ := NewSamRecordEmitter(tc.fields, "", "")
		samRecord := NewSamRecord(tc.rawRecord)
		actual := samRecordEmitter.customEmitFields(samRecord)
		assert.Equal(t, tc.exp, actual)
	}
}

// TestSamRecordEmitterCustomEmitTags tests customEmitTags function
func TestSamRecordEmitterCustomEmitTags(t *testing.T) {
	for _, tc := range samRecordEmitterCustomEmitTagsTC {
		samRecordEmitter, _ := NewSamRecordEmitter("", tc.tags, tc.notags)
		samRecord := NewSamRecord(tc.rawRecord)
		actual := samRecordEmitter.customEmitTags(samRecord)
		assert.Equal(t, tc.exp, actual)
	}
}
