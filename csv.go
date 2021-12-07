package jcsv

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"os"
	"strconv"
)

// ParseCsvFile opens, reads and parses a CSV file and returns a file object
func ParseCsvFile(path string, hasHeaders bool) (*file, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ParseOpenedCsvFile(f, hasHeaders)
}

// ParseOpenedCsvFile reads and parses an opened CSV file and returns a file object
func ParseOpenedCsvFile(f *os.File, hasHeaders bool) (*file, error) {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return ParseCsv(b, hasHeaders)
}

// ParseCsv parses CSV data and returns a file object
func ParseCsv(c []byte, hasHeaders bool) (*file, error) {
	data, err := csv.NewReader(bytes.NewBuffer(c)).ReadAll()
	if err != nil {
		return nil, err
	}

	// parsing csv headers
	var headers []string
	if hasHeaders {
		headers = data[0]
	} else {
		for i := range data[0] {
			headers = append(headers, "column"+strconv.Itoa(i))
		}
	}

	// parsing csv values
	var f file
	for _, rowData := range data[1:] {
		row := make(map[string]interface{})

		for i, value := range rowData {
			row[headers[i]] = value
		}

		f.data = append(f.data, row)
	}

	return &f, nil
}
