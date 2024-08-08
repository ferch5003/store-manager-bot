package files

import (
	"errors"
	"fmt"
	"strings"
)

// GetBase64Image returns the string on the base64 encoded string.
// It could be base64Data = String after data:{mimeType};base64.
func GetBase64Image(data, mimeType string) (string, error) {
	if !strings.Contains(data, "data:") {
		return "", errors.New("data is not encoded correctly")
	}

	// If data is encoded in correctly MIME type the strings.SplitAfterN may return at least 2 elements. The second
	// string is the base64 encoded image.
	base64Parts := strings.SplitAfterN(data, fmt.Sprintf("data:%s;base64,", mimeType), 2)
	if len(base64Parts) < 2 {
		return "", errors.New("data is not encoded correctly")
	}

	return base64Parts[1], nil
}
