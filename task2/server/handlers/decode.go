package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type userInput struct {
	Value string `json:"inputString"`
}

type processedOutput struct {
	Value string `json:"outputString"`
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	var data userInput
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	decoded, err := base64.StdEncoding.DecodeString(data.Value)
	if err != nil {
		http.Error(w, "Invalid base64 input", http.StatusBadRequest)
	}

	decodedString := processedOutput{
		Value: string(decoded),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(decodedString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
