package custom

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io"
	"net/http"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &Client{
		http: &http.Client{Transport: transport},
	}
}

type Envelope struct {
	XMLName  xml.Name `xml:"soapenv:Envelope"`
	Xmlns    string   `xml:"xmlns:soapenv,attr"`
	XmlnsGbd string   `xml:"xmlns:gbd,attr"`
	Header   Header
	Body     Body
}

type Header struct {
	XMLName xml.Name `xml:"soapenv:Header"`
	UserID  string   `xml:"userId"`
}

type Body struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Request string   `xml:",innerxml"`
}

func (c *Client) SendRequest(url string, xmlStr string) ([]byte, error) {
	soapRequest := Envelope{
		Xmlns:    "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsGbd: "http://data.gbd.chdb.scb.kz/",
		Header: Header{
			UserID: url,
		},
		Body: Body{
			Request: xmlStr,
		},
	}

	xmlData, err := xml.MarshalIndent(soapRequest, "", "  ")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(xml.Header+string(xmlData))))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

	//jsonData, err := xj.Convert(bytes.NewReader(body))
	//if err != nil {
	//	return "", err
	//}
	//
	//jsonResult, err := GetResultFromResponse(resultStr, jsonData)
	//if err != nil {
	//	return "", err
	//}
	//
	//response, err := json.MarshalIndent(jsonResult, "", "")
	//if err != nil {
	//	return "", err
	//}
	//
	//return string(response), nil
}
