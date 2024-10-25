package encoders

import (
	"encoding/base64"
)

// Decode64 converts an incoming base64 string into a regular string.
func Decode64(encoded64Data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded64Data)
}

// Encoder64 converts string to base64.
func Encoder64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
