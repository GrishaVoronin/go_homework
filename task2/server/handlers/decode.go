package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type inputString struct {
	InputString string `json:"inputString"`
}

type outputString struct {
	OutputString string `json:"outputString"`
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	var data inputString
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	decoded, err := base64.StdEncoding.DecodeString(data.InputString)
	if err != nil {
		http.Error(w, "Invalid base64 input", http.StatusBadRequest)
	}

	decodedString := outputString{
		OutputString: string(decoded),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(decodedString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
