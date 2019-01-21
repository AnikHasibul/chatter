package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateID generates a random hashed string based on the current unixnano timestamp.
func GenerateID() string {
	hasher := sha256.New()
	if _, err := hasher.Write([]byte(fmt.Sprint(time.Now().UnixNano()))); err != nil {
		panic(err.Error())
	}
	hash := hasher.Sum([]byte(fmt.Sprint(time.Now().UnixNano())))
	return hex.EncodeToString(hash)
}
