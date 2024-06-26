package x12

import (
	"strings"

	"github.com/nacimux/go-edi/common"
)

type X12Parser struct{}

func (p *X12Parser) Parse(message string) (common.Message, error) {
	segments := strings.Split(message, "~")
	var msg common.Message
	for _, segment := range segments {
		if segment == "" {
			continue // Skip empty segments
		}
		elements := strings.Split(segment, "*")
		if len(elements) > 0 {
			msg.Segments = append(msg.Segments, common.Segment{
				ID:       elements[0],
				Elements: elements[1:],
			})
		}
	}
	return msg, nil
}

func (p *X12Parser) Serialize(message common.Message) (string, error) {
	var segments []string
	for _, segment := range message.Segments {
		segments = append(segments, segment.ID+"*"+strings.Join(segment.Elements, "*"))
	}
	return strings.Join(segments, "~") + "~", nil // Adding the trailing ~
}
