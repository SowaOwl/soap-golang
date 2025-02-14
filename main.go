package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	xj "github.com/basgys/goxml2json"
	"github.com/joho/godotenv"
	"log"
	"os"
	"soap-go/custom"
	"soap-go/library"
	"soap-go/utils"
)

func main() {
	err := godotenv.Load(".env")

	jsonInputCustom, err := os.ReadFile("custom.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonInputLibrary, err := os.ReadFile("library.json")
	if err != nil {
		log.Fatal(err)
	}

	useCustom, err := sendUseCustom(string(jsonInputCustom), os.Getenv("CUSTOM_METHOD"), os.Getenv("CUSTOM_URL"))
	if err != nil {
		log.Fatal(err)
	}

	useLibrary, err := sendUseLibrary(string(jsonInputLibrary), os.Getenv("LIBRARY_METHOD"), os.Getenv("LIBRARY_URL"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CUSTOM -")
	fmt.Println(useCustom)
	fmt.Println("LIBRARY -")
	fmt.Println(useLibrary)
}

func sendUseCustom(jsonString string, methodName string, url string) (string, error) {
	resultStr := methodName + "Response"

	xmlOutput, err := utils.JsonToXML([]byte(jsonString), "gbd:"+methodName)
	if err != nil {
		return "", err
	}

	login := os.Getenv("CUSTOM_LOGIN")
	password := os.Getenv("CUSTOM_PASSWORD")

	client := custom.NewClient()

	response, err := client.SendRequest(
		url+methodName+"?wsdl",
		string(xmlOutput),
		login,
		password,
	)
	if err != nil {
		return "", err
	}

	jsonData, err := xj.Convert(bytes.NewReader(response))
	if err != nil {
		return "", err
	}

	jsonResult, err := utils.GetResultFromResponse(resultStr, jsonData)
	if err != nil {
		return "", err
	}

	responseStr, err := json.MarshalIndent(jsonResult, "", "")
	if err != nil {
		return "", err
	}

	return string(responseStr), nil
}

func sendUseLibrary(jsonString string, methodName string, url string) (string, error) {
	resultStr := methodName + "Result"

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return "", err
	}

	client := library.NewClient()

	response, err := client.SendRequest(url, methodName, data)

	if err != nil {
		return "", err
	}

	jsonData, err := xj.Convert(bytes.NewReader(response))
	if err != nil {
		return "", err
	}

	jsonResult, err := utils.GetResultFromResponse(resultStr, jsonData)
	if err != nil {
		return "", err
	}

	responseStr, err := json.MarshalIndent(jsonResult, "", "")
	if err != nil {
		return "", err
	}

	return string(responseStr), nil
}
