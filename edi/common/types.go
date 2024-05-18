package common

// Segment represents an EDI segment
type Segment struct {
	ID       string
	Elements []string
}

// Message represents an EDI message
type Message struct {
	Segments []Segment
}

// Parser interface for different EDI standards
type Parser interface {
	Parse(message string) (Message, error)
	Serialize(message Message) (string, error)
}
