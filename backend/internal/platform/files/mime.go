package files

import "strings"

const (
	MIMEImagePNG  = "image/png"
	MIMEImageJPEG = "image/jpeg"
)

var _validMIMETypes = [...]string{
	MIMEImagePNG,
	MIMEImageJPEG,
}

func GetMIMEType(base64Data string) (string, bool) {
	for _, validImageType := range _validMIMETypes {
		if strings.Contains(base64Data, validImageType) {
			return validImageType, true
		}
	}

	return "", false
}
