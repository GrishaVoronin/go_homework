package main

import (
	"client/queries"
	"fmt"
)

const host = "http://localhost:8080"

func main() {
	version, err := queries.GetVersion(host + "/version")
	if err != nil {
		fmt.Printf("Error getting version: %s\n", err)
	}
	fmt.Println(version)

	encodedString := "SGVsbG8gV29ybGQ="
	decoded, err := queries.DecodeString(host+"/decode", encodedString)
	if err != nil {
		fmt.Printf("Error decoding: %s\n", err)
	}
	fmt.Println(decoded)

	isDone, respCode, err := queries.GetHardOpQuery(host + "/hard-op")
	if err != nil {
		fmt.Printf("Error getting hard op: %s\n", err)
	}
	fmt.Println(isDone, respCode)
}
