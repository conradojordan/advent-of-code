package utils

import (
	"os"
	"strings"
)

func ParseInput(path string, delimiter string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	stringData := strings.TrimRight(string(data), "\n")

	sliceData := strings.Split(stringData, delimiter)
	return sliceData, nil
}
