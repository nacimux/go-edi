package edifact

import (
	"strings"

	"github.com/nacimux/go-edi/edi/common"
)

type EDIFACTParser struct{}

func (p *EDIFACTParser) Parse(message string) (common.Message, error) {
	segments := strings.Split(message, "'")
	var msg common.Message
	for _, segment := range segments {
		elements := strings.Split(segment, "+")
		if len(elements) > 0 {
			msg.Segments = append(msg.Segments, common.Segment{
				ID:       elements[0],
				Elements: elements[1:],
			})
		}
	}
	return msg, nil
}

func (p *EDIFACTParser) Serialize(message common.Message) (string, error) {
	var segments []string
	for _, segment := range message.Segments {
		segments = append(segments, segment.ID+"+"+strings.Join(segment.Elements, "+"))
	}
	return strings.Join(segments, "'"), nil
}
