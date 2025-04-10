package fileutils

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadFile(file *os.File) ([]string, error) {
	reader := csv.NewReader(file)

	record, err := reader.Read()

	if err == io.EOF {
		return record, err
	}

	return record, nil
}
