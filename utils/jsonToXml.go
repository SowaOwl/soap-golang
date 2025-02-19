package utils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
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

func JsonToXML(jsonData []byte, rootName string, rootAttr xml.Attr) ([]byte, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	root := Element{XMLName: xml.Name{Local: rootName}, Attrs: []xml.Attr{rootAttr}}
	for k, v := range data {
		root.Content = append(root.Content, toXMLElement(k, v))
	}

	return xml.MarshalIndent(root, "", "  ")
}

func JsonMapToXml(jsonMap map[string]interface{}, rootName string, rootAttr xml.Attr) ([]byte, error) {
	root := Element{XMLName: xml.Name{Local: rootName}, Attrs: []xml.Attr{rootAttr}}
	for k, v := range jsonMap {
		root.Content = append(root.Content, toXMLElement(k, v))
	}

	return xml.MarshalIndent(root, "", "  ")
}

func GetResultFromResponse(resultKey string, jsonData *bytes.Buffer) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(jsonData.Bytes(), &result)
	if err != nil {
		return nil, err
	}

	re, ok := findKey(result, resultKey)
	if ok {
		return re, nil
	} else {
		return nil, errors.New("result key not found")
	}
}

func findKey(data map[string]interface{}, key string) (map[string]interface{}, bool) {
	for k, v := range data {
		if k == key {
			return v.(map[string]interface{}), true
		}
		if nestedMap, ok := v.(map[string]interface{}); ok {
			if foundValue, found := findKey(nestedMap, key); found {
				return foundValue, true
			}
		}
	}
	return nil, false
}
