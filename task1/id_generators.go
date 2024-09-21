package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateId1(name string) string {
	hash := md5.New()
	hash.Write([]byte(name))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateId2(name string) string {
	hash := sha256.New()
	hash.Write([]byte(name))
	return hex.EncodeToString(hash.Sum(nil))
}
