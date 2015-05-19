package main

import (
	"os"
	"regexp"
)

type FileEvent struct {
	Source           *string `json:"source,omitempty"`
	Offset           int64   `json:"offset,omitempty"`
	Line             uint64  `json:"line,omitempty"`
	Text             *string `json:"text,omitempty"`
	Fields           *map[string]string
	FieldNames       []string `json:fields`
	FieldTypes       []string `json:fields`
	DelimiterRegexp  *regexp.Regexp
	QuoteChar        string
	FieldNamesLength int

	ileinfo  *os.FileInfo
	fileinfo *os.FileInfo
}
