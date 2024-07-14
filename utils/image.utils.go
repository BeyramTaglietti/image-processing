package utils

import (
	"encoding/base64"
)

func ToBase64(src []byte, format string) string {
	base64Image := base64.StdEncoding.EncodeToString(src)
	imageSource := "data:image/" + format + ";base64," + base64Image
	return imageSource
}
