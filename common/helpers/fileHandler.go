package helpers

import (
	"encoding/csv"
	"os"
	"strings"
)

// Parse json to struct type
func ParseJSON() int {
	return 42
}

// Parse yaml to struct type
func ParseYAML() int {
	return 42
}

// Write json to CSV file
func JSONToCSV() int {
	return 42
}

func WriteToCsv(data [][]string, filePath string) {
	tokens := strings.Split(filePath, "/")
	path := strings.Join(tokens[:len(tokens)-1], "/")

	err := os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		writer.Write(value)
	}
}
