package util

import "encoding/base64"

func Base64EncodeString(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
