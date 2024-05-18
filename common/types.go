package common

type Segment struct {
	ID       string
	Elements []string
}

type Message struct {
	Segments []Segment
}

type Parser interface {
	Parse(message string) (Message, error)
	Serialize(message Message) (string, error)
}
