package queries

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type inputString struct {
	InputString string `json:"inputString"`
}

type outputString struct {
	OutputString string `json:"outputString"`
}

func DecodeString(url string, input_string string) (string, error) {
	string_to_decode := inputString{InputString: input_string}
	jsonData, err := json.Marshal(string_to_decode)
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

	var output_string outputString
	err = json.Unmarshal(body, &output_string)
	if err != nil {
		return "", err
	}
	return output_string.OutputString, nil
}
