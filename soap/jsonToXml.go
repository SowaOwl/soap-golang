package soap

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Element struct {
	XMLName xml.Name
	Attrs   []xml.Attr    `xml:",attr,omitempty"`
	Content []interface{} `xml:",any"`
	Text    string        `xml:",chardata"`
}

func toXMLElement(key string, value interface{}) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		var elements []interface{}
		for k, val := range v {
			elements = append(elements, toXMLElement(k, val))
		}
		return Element{XMLName: xml.Name{Local: key}, Content: elements}

	case []interface{}:
		var elements []interface{}
		for _, item := range v {
			elements = append(elements, toXMLElement(key, item))
		}
		return elements

	default:
		return Element{XMLName: xml.Name{Local: key}, Text: fmt.Sprintf("%v", v)}
	}
}

func JsonToXML(jsonData []byte, rootName string) ([]byte, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	root := Element{XMLName: xml.Name{Local: rootName}}
	for k, v := range data {
		root.Content = append(root.Content, toXMLElement(k, v))
	}

	return xml.MarshalIndent(root, "", "  ")
}
