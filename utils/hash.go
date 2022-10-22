package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GeneratePassHash(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
