package edi

import (
	"github.com/nacimux/go-edi/common"
	"github.com/nacimux/go-edi/edifact"
	"github.com/nacimux/go-edi/x12"
)

type ParserType int

const (
	X12 ParserType = iota
	EDIFACT
)

func NewParser(parserType ParserType) common.Parser {
	switch parserType {
	case X12:
		return &x12.X12Parser{}
	case EDIFACT:
		return &edifact.EDIFACTParser{}
	default:
		return nil
	}
}
