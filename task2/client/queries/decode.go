package queries

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type userInput struct {
	Value string `json:"inputString"`
}

type processedOutput struct {
	Value string `json:"outputString"`
}

func DecodeString(url string, data string) (string, error) {
	stringToDecode := userInput{Value: data}
	jsonData, err := json.Marshal(stringToDecode)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var outputString processedOutput
	err = json.Unmarshal(body, &outputString)
	if err != nil {
		return "", err
	}
	return outputString.Value, nil
}
