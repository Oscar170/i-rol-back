package base64

import b64 "encoding/base64"

// Decode microwrapper to base64 lib to avoid string cast.
func Decode(data string) (string, error) {
	dec, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	return string(dec), err
}

// Encode microwrapper to base64 lib to avoid byte cast.
func Encode(data string) string {
	return b64.StdEncoding.EncodeToString([]byte(data))
}
