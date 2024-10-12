package main

import (
	"client/queries"
	"fmt"
)

func main() {
	version_url := "http://localhost:8080/version"
	version, err := queries.GetVersion(version_url)
	if err != nil {
		fmt.Printf("Error getting version: %s\n", err)
	}
	fmt.Println(version)

	decode_url := "http://localhost:8080/decode"
	encodedString := "SGVsbG8gV29ybGQ="
	decoded, err := queries.DecodeString(decode_url, encodedString)
	if err != nil {
		fmt.Printf("Error decoding: %s\n", err)
	}
	fmt.Println(decoded)

	hard_op_url := "http://localhost:8080/hard-op"
	is_done, resp_code, err := queries.GetHardOpQuery(hard_op_url)
	if err != nil {
		fmt.Printf("Error getting hard op: %s\n", err)
	}
	fmt.Println(is_done, resp_code)
}
