package edifact

import (
	"reflect"
	"testing"

	"github.com/nacimux/go-edi/common"
)

func TestEDIFACTParser_Parse(t *testing.T) {
	parser := &EDIFACTParser{}
	message := "UNH+1+ORDERS:D:96A:UN:EAN008'BGM+220+BKOD01+9'"
	expectedMessage := common.Message{
		Segments: []common.Segment{
			{
				ID:       "UNH",
				Elements: []string{"1", "ORDERS:D:96A:UN:EAN008"},
			},
			{
				ID:       "BGM",
				Elements: []string{"220", "BKOD01", "9"},
			},
		},
	}

	result, err := parser.Parse(message)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if !reflect.DeepEqual(result, expectedMessage) {
		t.Errorf("Parse() = %v, want %v", result, expectedMessage)
	}
}

func TestEDIFACTParser_Serialize(t *testing.T) {
	parser := &EDIFACTParser{}
	message := common.Message{
		Segments: []common.Segment{
			{
				ID:       "UNH",
				Elements: []string{"1", "ORDERS:D:96A:UN:EAN008"},
			},
			{
				ID:       "BGM",
				Elements: []string{"220", "BKOD01", "9"},
			},
		},
	}
	expectedMessage := "UNH+1+ORDERS:D:96A:UN:EAN008'BGM+220+BKOD01+9'"

	result, err := parser.Serialize(message)
	if err != nil {
		t.Fatalf("Serialize() error = %v", err)
	}
	if result != expectedMessage {
		t.Errorf("Serialize() = %v, want %v", result, expectedMessage)
	}
}
