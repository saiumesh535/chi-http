package utils

import (
	"os"
	"net/http"
)

// GetFileContentType file type
func GetFileContentType(out *os.File) (string, error) {

    // Only the first 512 bytes are used to sniff the content type.
    buffer := make([]byte, 512)

    _, err := out.Read(buffer)
    if err != nil {
        return "", err
    }

    // Use the net/http package's handy DectectContentType function. Always returns a valid
    // content-type by returning "application/octet-stream" if no others seemed to match.
    contentType := http.DetectContentType(buffer)

    return contentType, nil
}