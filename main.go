package main

import (
	"fmt"
	"log"
	"soap-go/soap"
	"time"
)

func main() {
	jsonInput := `{
		"uin": "930421400841",
		"company": "АО «СК«Amanat»",
		"company_bin": "000940001446",
		"respon-uin": "001119550524",
		"expiresIn": "157680000"
	}`

	startTime := time.Now()

	xmlOutput, err := soap.JsonToXML([]byte(jsonInput), "gbd:getPersonDataAccessControlForESBD")
	if err != nil {
		log.Fatal(err)
	}

	executionTime := time.Since(startTime)

	fmt.Println(string(xmlOutput))

	fmt.Println(executionTime)
}
