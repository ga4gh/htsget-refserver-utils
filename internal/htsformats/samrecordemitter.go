// Package htsformats contains objects modeling entities in genomic file formats
// and associated behaviors
//
// Module samrecordemitter emits individual alignments with fields and
// tags included/excluded according to custom parameters
package htsformats

import (
	"errors"
	"strings"

	"github.com/ga4gh/htsget-refserver-utils/internal/htsutils"
	"github.com/golang-collections/collections/set"
)

// samFields map of canonical field name to its column position in a SamRecord
var samFields = map[string]int{
	"QNAME": 0,
	"FLAG":  1,
	"RNAME": 2,
	"POS":   3,
	"MAPQ":  4,
	"CIGAR": 5,
	"RNEXT": 6,
	"PNEXT": 7,
	"TLEN":  8,
	"SEQ":   9,
	"QUAL":  10,
}

// samFieldReplacements if a field is marked to be excluded, the true value will
// be replaced by the value within the list corresponding to the column position
var samFieldReplacements = []string{
	"*",   // QNAME
	"0",   // FLAG
	"*",   // RNAME
	"0",   // POS
	"255", // MAPQ
	"*",   // CIGAR
	"*",   // RNEXT
	"0",   // PNEXT
	"0",   // TLEN
	"*",   // SEQ
	"*",   // QUAL
}

// SamRecordEmitter performs custom field/tag emitting of SamRecords based on
// requested properties. Can include/exclude specific fields and tags
type SamRecordEmitter struct {
	emitAllFields bool
	emitAllTags   bool
	inclusionEmit bool
	fields        []bool
	tags          []string
	notags        []string
}

// NewSamRecordEmitter constructs and configures a SamRecordEmitter
func NewSamRecordEmitter(fields string, tags string, notags string) (*SamRecordEmitter, error) {
	samRecordEmitter := new(SamRecordEmitter)

	// setup fields-related attributes
	err := samRecordEmitter.setupFields(fields)
	if err != nil {
		return nil, err
	}

	// setup tags and notags related attributes
	err = samRecordEmitter.setupTagsNotags(tags, notags)
	if err != nil {
		return nil, err
	}
	return samRecordEmitter, nil
}

// setupFields configures field-related properties of the SamRecordEmitter,
// including an indication of whether all fields should be emitted, and what
// fields should be included/excluded
func (samRecordEmitter *SamRecordEmitter) setupFields(fields string) error {

	if fields == "" {
		// emit all fields if the requested fields was an empty string
		samRecordEmitter.emitAllFields = true
	} else {
		// otherwise, set each field's emittance to off (false), only setting it
		// on (true) if the canonical field name was found in the fields list
		samRecordEmitter.emitAllFields = false
		samRecordEmitter.fields = make([]bool, 11)
		requestedFields := strings.Split(fields, ",")
		for _, requestedField := range requestedFields {
			if val, ok := samFields[requestedField]; ok {
				samRecordEmitter.fields[val] = true
			} else {
				return errors.New("Invalid field: '" + requestedField + "'")
			}
		}
	}
	return nil
}

// setupTagsNotags configures tag emitting properties of the SamRecordEmitter
// including an indication of whether all fields should be emitted
func (samRecordEmitter *SamRecordEmitter) setupTagsNotags(tags string, notags string) error {
	samRecordEmitter.tags = []string{}
	samRecordEmitter.notags = []string{}

	if tags == "" && notags == "" {
		// if neither tags nor notags specified, emit all tags
		samRecordEmitter.emitAllTags = true
	} else {
		// if 'tags' was specified, then don't emit by default (inclusionEmit = true),
		// only emitting by those tags specified/included by the tags string
		// if 'tags' wasn't specified, emit everything by default (inclusionEmit = false),
		// only excluding what's specified in notags
		samRecordEmitter.inclusionEmit = false
		if tags != "" {
			samRecordEmitter.inclusionEmit = true
			samRecordEmitter.tags = strings.Split(tags, ",")
		}

		if notags != "" {
			samRecordEmitter.notags = strings.Split(notags, ",")
		}

		// check for overlap between tags and notags
		tagsSet := htsutils.CreateSet(samRecordEmitter.tags)
		notagsSet := htsutils.CreateSet(samRecordEmitter.notags)
		if tagsSet.Intersection(notagsSet).Len() > 0 {
			return errors.New("Overlap between 'tags' and 'notags'")
		}
	}
	return nil
}

// CustomEmit accepts an unmodified SamRecord, and returns a string representing
// the SamRecord after modification by field and tag inclusion/exclusion
func (samRecordEmitter *SamRecordEmitter) CustomEmit(samRecord *SamRecord) string {

	customEmit := []string{}

	// emit fields
	if samRecordEmitter.emitAllFields {
		emitFields := samRecord.emitFields()
		customEmit = append(customEmit, emitFields...)
	} else {
		customEmitFields := samRecordEmitter.customEmitFields(samRecord)
		customEmit = append(customEmit, customEmitFields...)
	}

	// emit tags
	if samRecordEmitter.emitAllTags {
		emitTags := samRecord.emitTags()
		customEmit = append(customEmit, emitTags...)
	} else {
		customEmitTags := samRecordEmitter.customEmitTags(samRecord)
		customEmit = append(customEmit, customEmitTags...)
	}

	return strings.Join(customEmit, "\t")
}

// customEmitFields returns only the specified fields of a SamRecord, replacing
// excluded fields with their appropriate non-specified values
func (samRecordEmitter *SamRecordEmitter) customEmitFields(samRecord *SamRecord) []string {

	// for each field, check if the corresponding column position is set to
	// true (emit real value). If true, get and emit the real value, if false,
	// emit the replacement value
	customEmitFields := make([]string, 11)
	for i := 0; i < len(samRecordEmitter.fields); i++ {
		if samRecordEmitter.fields[i] {
			customEmitFields[i] = samRecord.getField(i)
		} else {
			customEmitFields[i] = samFieldReplacements[i]
		}
	}
	return customEmitFields
}

// customEmitTags returns only the specified tags of a SamRecord, excluding everything
// either not specified by 'tags' or specified by 'notags'
func (samRecordEmitter *SamRecordEmitter) customEmitTags(samRecord *SamRecord) []string {

	customEmitTags := []string{}
	allTagKeysSet := htsutils.CreateSet(samRecord.tagKeys)
	var tagsToEmit *set.Set

	if samRecordEmitter.inclusionEmit {
		// (inclusionEmit = true) means emit ONLY the tags specified by 'tags'
		// get overlap between available tags in sam record, and those specified in tags
		requestedTagsSet := htsutils.CreateSet(samRecordEmitter.tags)
		tagsToEmit = allTagKeysSet.Intersection(requestedTagsSet)
	} else {
		// (inclusionEmit = false) means emit everything EXCEPT the tags specified by 'notags'
		// get the difference between available tags in sam record, and those specified in notags
		requestedNoTagsSet := htsutils.CreateSet(samRecordEmitter.notags)
		tagsToEmit = allTagKeysSet.Difference(requestedNoTagsSet)
	}

	// for each tag in the set, lookup the tag value and emit
	for _, tagKey := range samRecord.tagKeys {
		if tagsToEmit.Has(tagKey) {
			customEmitTags = append(customEmitTags, samRecord.getTag(tagKey))
		}
	}

	return customEmitTags
}
