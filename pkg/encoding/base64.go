package encoding

import "encoding/base64"

func Base64RawStdEncodeToString(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
}

func Base64RawStdDecodeString(src string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(src)
}
