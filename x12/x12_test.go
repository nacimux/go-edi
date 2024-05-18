package x12

import (
	"reflect"
	"testing"

	"github.com/nacimux/go-edi/common"
)

func TestParse(t *testing.T) {
	parser := &X12Parser{}
	message := "ISA*00*          *00*          *ZZ*123456789      *ZZ*987654321      *210101*1253*U*00401*000000001*0*T*:~GS*PO*123456789*987654321*20210101*1253*1*X*004010~"
	expectedMessage := common.Message{
		Segments: []common.Segment{
			{
				ID:       "ISA",
				Elements: []string{"00", "          ", "00", "          ", "ZZ", "123456789      ", "ZZ", "987654321      ", "210101", "1253", "U", "00401", "000000001", "0", "T", ":"},
			},
			{
				ID:       "GS",
				Elements: []string{"PO", "123456789", "987654321", "20210101", "1253", "1", "X", "004010"},
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

func TestSerialize(t *testing.T) {
	parser := &X12Parser{}
	message := common.Message{
		Segments: []common.Segment{
			{
				ID:       "ISA",
				Elements: []string{"00", "          ", "00", "          ", "ZZ", "123456789      ", "ZZ", "987654321      ", "210101", "1253", "U", "00401", "000000001", "0", "T", ":"},
			},
			{
				ID:       "GS",
				Elements: []string{"PO", "123456789", "987654321", "20210101", "1253", "1", "X", "004010"},
			},
		},
	}
	expectedMessage := "ISA*00*          *00*          *ZZ*123456789      *ZZ*987654321      *210101*1253*U*00401*000000001*0*T*:~GS*PO*123456789*987654321*20210101*1253*1*X*004010~"

	result, err := parser.Serialize(message)
	if err != nil {
		t.Fatalf("Serialize() error = %v", err)
	}
	if result != expectedMessage {
		t.Errorf("Serialize() = %v, want %v", result, expectedMessage)
	}
}
