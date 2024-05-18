package edifact

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

// ConvertToCSV converts EDIFACT message to CSV format
func ConvertToCSV(edifactMessage string) (string, error) {
	// Split EDIFACT message into segments
	segments := strings.Split(edifactMessage, "'")

	var csvData strings.Builder
	writer := csv.NewWriter(&csvData)
	writer.Comma = ','

	for _, segment := range segments {
		fields := strings.Split(segment, "+")
		if err := writer.Write(fields); err != nil {
			return "", err
		}
	}

	writer.Flush()
	return csvData.String(), nil
}

func ConvertFromCSV(csvData string) (string, error) {

	reader := csv.NewReader(strings.NewReader(csvData))
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	var edifactMessage strings.Builder
	for _, record := range records {
		edifactMessage.WriteString(strings.Join(record, "+"))
		edifactMessage.WriteString("'")
	}

	return edifactMessage.String(), nil
}

func ConvertToJSON(edifactMessage string) (string, error) {
	segments := strings.Split(edifactMessage, "'")

	var jsonData []map[string]interface{}
	for _, segment := range segments {
		fields := strings.Split(segment, "+")
		segmentMap := make(map[string]interface{})
		segmentMap["ID"] = fields[0]
		segmentMap["Elements"] = fields[1:]
		jsonData = append(jsonData, segmentMap)
	}

	jsonString, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

func ConvertFromJSON(jsonData string) (string, error) {
	var segments []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &segments); err != nil {
		return "", err
	}

	var edifactMessage strings.Builder
	for _, segment := range segments {
		id := segment["ID"].(string)
		elements := segment["Elements"].([]interface{})
		var fields []string
		fields = append(fields, id)
		for _, element := range elements {
			fields = append(fields, element.(string))
		}
		edifactMessage.WriteString(strings.Join(fields, "+"))
		edifactMessage.WriteString("'")
	}

	return edifactMessage.String(), nil
}

func ConvertToXML(edifactMessage string) (string, error) {

	segments := strings.Split(edifactMessage, "'")

	var xmlData strings.Builder
	xmlData.WriteString("<EDIFACTMessage>")
	for _, segment := range segments {
		fields := strings.Split(segment, "+")
		xmlData.WriteString("<Segment>")
		xmlData.WriteString(fmt.Sprintf("<ID>%s</ID>", fields[0]))
		for _, element := range fields[1:] {
			xmlData.WriteString(fmt.Sprintf("<Element>%s</Element>", element))
		}
		xmlData.WriteString("</Segment>")
	}
	xmlData.WriteString("</EDIFACTMessage>")

	return xmlData.String(), nil
}

func ConvertFromXML(xmlData string) (string, error) {

	type Segment struct {
		ID       string   `xml:"ID"`
		Elements []string `xml:"Element"`
	}
	type EDIFACTMessage struct {
		Segments []Segment `xml:"Segment"`
	}
	var edifactMsg EDIFACTMessage
	if err := xml.Unmarshal([]byte(xmlData), &edifactMsg); err != nil {
		return "", err
	}

	var edifactMessage strings.Builder
	for _, segment := range edifactMsg.Segments {
		edifactMessage.WriteString(segment.ID)
		for _, element := range segment.Elements {
			edifactMessage.WriteString("+")
			edifactMessage.WriteString(element)
		}
		edifactMessage.WriteString("'")
	}

	return edifactMessage.String(), nil
}
