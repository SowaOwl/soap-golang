package constructor

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"soap-go/utils"
)

type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Xmlns   string   `xml:"xmlns:soap,attr"`
	Body    Body
}

type Body struct {
	XMLName xml.Name `xml:"soap:Body"`
	Request string   `xml:",innerxml"`
}

func NewRequestFromJson(data map[string]interface{}, methodName string, methodAttr xml.Attr, url string) (*http.Request, error) {
	xmlBody, err := utils.JsonMapToXml(data, methodName, methodAttr)
	if err != nil {
		return &http.Request{}, err
	}

	body := Envelope{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:  Body{Request: string(xmlBody)},
	}

	xmlData, err := xml.MarshalIndent(body, "", "")

	fmt.Println(string(xmlData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(xml.Header+string(xmlData))))
	if err != nil {
		return &http.Request{}, err
	}

	return req, nil
}

func NewRequestByEnvelope(envelope interface{}, url string) (*http.Request, error) {
	xmlData, err := xml.MarshalIndent(envelope, "", "")
	if err != nil {
		return &http.Request{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(xml.Header+string(xmlData))))
	if err != nil {
		return &http.Request{}, err
	}

	return req, nil
}
