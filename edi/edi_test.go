package edi

import (
	"fmt"
	"testing"

	"github.com/nacimux/go-edi/edi/common"
	"github.com/nacimux/go-edi/edi/edifact"
	"github.com/nacimux/go-edi/edi/x12"
)

func TestNewParser(t *testing.T) {
	tests := []struct {
		parserType ParserType
		expected   common.Parser
	}{
		{
			parserType: X12,
			expected:   &x12.X12Parser{},
		},
		{
			parserType: EDIFACT,
			expected:   &edifact.EDIFACTParser{},
		},
		{
			parserType: ParserType(99), // invalid type
			expected:   nil,
		},
	}

	for _, test := range tests {
		t.Run(testName(test.parserType), func(t *testing.T) {
			result := NewParser(test.parserType)
			if result == nil && test.expected != nil {
				t.Errorf("NewParser(%v) = nil, want %T", test.parserType, test.expected)
			} else if result != nil && test.expected == nil {
				t.Errorf("NewParser(%v) = %T, want nil", test.parserType, result)
			} else if result != nil && test.expected != nil && resultType(result) != resultType(test.expected) {
				t.Errorf("NewParser(%v) = %T, want %T", test.parserType, result, test.expected)
			}
		})
	}
}

func resultType(parser common.Parser) string {
	if parser == nil {
		return "nil"
	}
	return fmt.Sprintf("%T", parser)
}

func testName(parserType ParserType) string {
	switch parserType {
	case X12:
		return "X12"
	case EDIFACT:
		return "EDIFACT"
	default:
		return "Unknown"
	}
}
