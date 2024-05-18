package common_test

import (
	"testing"

	"github.com/nacimux/go-edi/edi/common"
)

func TestSegment(t *testing.T) {
	elements := []string{"element1", "element2"}
	segment := common.Segment{
		ID:       "SEG",
		Elements: elements,
	}

	// Test ID
	if segment.ID != "SEG" {
		t.Errorf("Expected segment ID to be SEG, but got %s", segment.ID)
	}

	// Test Elements
	if len(segment.Elements) != len(elements) {
		t.Errorf("Expected segment elements to have length %d, but got %d", len(elements), len(segment.Elements))
	}
	for i := range elements {
		if segment.Elements[i] != elements[i] {
			t.Errorf("Expected element at index %d to be %s, but got %s", i, elements[i], segment.Elements[i])
		}
	}
}

func TestMessage(t *testing.T) {
	segments := []common.Segment{
		{ID: "SEG1", Elements: []string{"element1", "element2"}},
		{ID: "SEG2", Elements: []string{"element3", "element4"}},
	}
	message := common.Message{
		Segments: segments,
	}

	// Test Segments
	if len(message.Segments) != len(segments) {
		t.Errorf("Expected message segments to have length %d, but got %d", len(segments), len(message.Segments))
	}
	for i := range segments {
		if message.Segments[i].ID != segments[i].ID {
			t.Errorf("Expected segment ID at index %d to be %s, but got %s", i, segments[i].ID, message.Segments[i].ID)
		}
		// Test elements within segment
		if len(message.Segments[i].Elements) != len(segments[i].Elements) {
			t.Errorf("Expected segment elements to have length %d, but got %d", len(segments[i].Elements), len(message.Segments[i].Elements))
		}
		for j := range segments[i].Elements {
			if message.Segments[i].Elements[j] != segments[i].Elements[j] {
				t.Errorf("Expected element at index %d of segment %s to be %s, but got %s", j, segments[i].ID, segments[i].Elements[j], message.Segments[i].Elements[j])
			}
		}
	}
}

type MockParser struct{}

func (p *MockParser) Parse(message string) (common.Message, error) {
	return common.Message{}, nil
}

func (p *MockParser) Serialize(message common.Message) (string, error) {
	return "", nil
}

func TestParser(t *testing.T) {

	parser := &MockParser{}

	_, err := parser.Parse("test message")
	if err != nil {
		t.Errorf("Parse() returned an error: %v", err)
	}

	_, err = parser.Serialize(common.Message{})
	if err != nil {
		t.Errorf("Serialize() returned an error: %v", err)
	}
}
