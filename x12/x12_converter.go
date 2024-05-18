package x12

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

func ConvertToCSV(x12Message string) (string, error) {

	segments := strings.Split(x12Message, "~")

	var csvData strings.Builder
	writer := csv.NewWriter(&csvData)
	writer.Comma = ','

	for _, segment := range segments {
		fields := strings.Split(segment, "*")
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

	var x12Message strings.Builder
	for _, record := range records {
		x12Message.WriteString(strings.Join(record, "*"))
		x12Message.WriteString("~")
	}

	return x12Message.String(), nil
}

// ConvertToJSON converts X12 message to JSON format
func ConvertToJSON(x12Message string) (string, error) {
	// Split X12 message into segments
	segments := strings.Split(x12Message, "~")

	// Create JSON representation
	var jsonData []map[string]interface{}
	for _, segment := range segments {
		fields := strings.Split(segment, "*")
		segmentMap := make(map[string]interface{})
		segmentMap["ID"] = fields[0]
		segmentMap["Elements"] = fields[1:]
		jsonData = append(jsonData, segmentMap)
	}

	// Marshal JSON data
	jsonString, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

// ConvertFromJSON converts JSON data to X12 format
func ConvertFromJSON(jsonData string) (string, error) {
	// Unmarshal JSON data
	var segments []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &segments); err != nil {
		return "", err
	}

	// Convert JSON data to X12 message
	var x12Message strings.Builder
	for _, segment := range segments {
		id := segment["ID"].(string)
		elements := segment["Elements"].([]interface{})
		var fields []string
		fields = append(fields, id)
		for _, element := range elements {
			fields = append(fields, element.(string))
		}
		x12Message.WriteString(strings.Join(fields, "*"))
		x12Message.WriteString("~")
	}

	return x12Message.String(), nil
}

// ConvertToXML converts X12 message to XML format
func ConvertToXML(x12Message string) (string, error) {
	// Split X12 message into segments
	segments := strings.Split(x12Message, "~")

	// Create XML representation
	var xmlData strings.Builder
	xmlData.WriteString("<X12Message>")
	for _, segment := range segments {
		fields := strings.Split(segment, "*")
		xmlData.WriteString("<Segment>")
		xmlData.WriteString(fmt.Sprintf("<ID>%s</ID>", fields[0]))
		for _, element := range fields[1:] {
			xmlData.WriteString(fmt.Sprintf("<Element>%s</Element>", element))
		}
		xmlData.WriteString("</Segment>")
	}
	xmlData.WriteString("</X12Message>")

	return xmlData.String(), nil
}

// ConvertFromXML converts XML data to X12 format
func ConvertFromXML(xmlData string) (string, error) {
	// Parse XML data
	type Segment struct {
		ID       string   `xml:"ID"`
		Elements []string `xml:"Element"`
	}
	type X12Message struct {
		Segments []Segment `xml:"Segment"`
	}
	var x12Msg X12Message
	if err := xml.Unmarshal([]byte(xmlData), &x12Msg); err != nil {
		return "", err
	}

	// Convert XML data to X12 message
	var x12Message strings.Builder
	for _, segment := range x12Msg.Segments {
		x12Message.WriteString(segment.ID)
		for _, element := range segment.Elements {
			x12Message.WriteString("*")
			x12Message.WriteString(element)
		}
		x12Message.WriteString("~")
	}

	return x12Message.String(), nil
}
