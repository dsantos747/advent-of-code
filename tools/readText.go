package tools

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadInput returns a slice of all lines in r. oc: ashishjh-bst
func ReadInput(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s, error: %s", path, err.Error())
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s, error: %s", path, err.Error())
	}

	// Below line hopefully accounts for usage of CRLF
	dataString := strings.ReplaceAll(string(data), "\r\n", "\n")
	return dataString, nil
}