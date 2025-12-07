package utils

import (
	"os"
	"strings"
)

func ParseInput(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	stringData := string(data)

	sliceData := strings.Split(stringData, "\n")
	if sliceData[len(sliceData)-1] == "" {
		sliceData = sliceData[:len(sliceData)-1]
	}
	return sliceData, nil
}
