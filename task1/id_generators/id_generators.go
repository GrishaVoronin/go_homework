package id_generators

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateIdMD5(name string) string {
	hash := md5.New()
	hash.Write([]byte(name))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateIdSHA256(name string) string {
	hash := sha256.New()
	hash.Write([]byte(name))
	return hex.EncodeToString(hash.Sum(nil))
}
