// Package htsformats contains objects modeling entities in genomic file formats
// and associated behaviors
//
// Module samrecord defines single alignments, as appearing in a SAM file
package htsformats

import (
	"strings"
)

// SamRecord holds all fields and tags from a single alignment of a SAM file
type SamRecord struct {
	raw     string
	qname   string
	flag    string
	rname   string
	pos     string
	mapq    string
	cigar   string
	rnext   string
	pnext   string
	tlen    string
	seq     string
	qual    string
	fields  []string
	tagKeys []string
	tags    map[string]string
}

// NewSamRecord constructs a SamRecord from a single line
func NewSamRecord(raw string) *SamRecord {
	samRecord := new(SamRecord)
	samRecord.raw = raw

	// split string and assign each field
	split := strings.Split(raw, "\t")
	samRecord.qname = split[0]
	samRecord.flag = split[1]
	samRecord.rname = split[2]
	samRecord.pos = split[3]
	samRecord.mapq = split[4]
	samRecord.cigar = split[5]
	samRecord.rnext = split[6]
	samRecord.pnext = split[7]
	samRecord.tlen = split[8]
	samRecord.seq = split[9]
	samRecord.qual = split[10]
	samRecord.fields = split

	// create a map of tag keys to tag value (entire tag), and also a list of
	// the tag keys within the map
	samRecord.tagKeys = []string{}
	samRecord.tags = make(map[string]string)
	for i := 11; i < len(split); i++ {
		tagval := split[i]
		tagkey := strings.Split(tagval, ":")[0]
		samRecord.tagKeys = append(samRecord.tagKeys, tagkey)
		samRecord.tags[tagkey] = tagval
	}

	return samRecord
}

// emitFields emits all SAM record fields without modification
func (samRecord *SamRecord) emitFields() []string {
	return []string{
		samRecord.qname,
		samRecord.flag,
		samRecord.rname,
		samRecord.pos,
		samRecord.mapq,
		samRecord.cigar,
		samRecord.rnext,
		samRecord.pnext,
		samRecord.tlen,
		samRecord.seq,
		samRecord.qual,
	}
}

// getField retrieves the corresponding field value given the column position
// (0-10 inclusive)
func (samRecord *SamRecord) getField(col int) string {
	return samRecord.fields[col]
}

// emitTags emits all tags within the SAM record without modification, in the
// same order they were parsed
func (samRecord *SamRecord) emitTags() []string {
	tags := []string{}
	for _, tagkey := range samRecord.tagKeys {
		tags = append(tags, samRecord.tags[tagkey])
	}
	return tags
}

// getTag gets a parsed tag value by its two-letter tag name/key
func (samRecord *SamRecord) getTag(key string) string {
	return samRecord.tags[key]
}

// String gets a string representation of the SamRecord
func (samRecord *SamRecord) String() string {
	return "[SamRecord qname=" + samRecord.qname + "]"
}
