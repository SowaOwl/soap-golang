package soap

import (
	"encoding/xml"
)

type SoapEnvelope struct {
	XMLName  xml.Name `xml:"soapenv:Envelope"`
	Xmlns    string   `xml:"xmlns:soapenv,attr"`
	XmlnsGbd string   `xml:"xmlns:gbd,attr"`
	Header   SoapHeader
	Body     SoapBody
}

type SoapHeader struct {
	XMLName xml.Name `xml:"soapenv:Header"`
	UserID  string   `xml:"userId"`
}

type SoapBody struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Request string   `xml:",innerxml"`
}
