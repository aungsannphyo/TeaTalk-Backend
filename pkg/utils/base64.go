package utils

import "encoding/base64"

func DecodeBase64(request string) ([]byte, error) {
	decodeByte, err := base64.StdEncoding.DecodeString(request)

	if err != nil {
		return nil, err
	}

	return decodeByte, nil
}

func EncodeBase64(request []byte) string {
	return base64.StdEncoding.EncodeToString(request)
}
