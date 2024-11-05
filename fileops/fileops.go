package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)

	os.WriteFile(fileName, []byte(valueText), 0644)
}
