package jcsv

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
)

type file struct {
	data []map[string]interface{}
}

// Csv returns the data in the CSV format
func (f *file) Csv(addHeaders bool) []byte {
	if f.data == nil {
		panic("cannot convert nil")
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)

	// writing csv headers
	if addHeaders {
		var headers []string
		for key, _ := range f.data[0] {
			headers = append(headers, key)
		}

		w.Write(headers)
	}

	// writing csv values
	for _, rowData := range f.data {
		var row []string
		for _, value := range rowData {
			row = append(row, value.(string))
		}

		w.Write(row)
	}
	
	w.Flush()
	
	return buf.Bytes()
}

// Json returns the data in the JSON format
func (f *file) Json() []byte {
	if f.data == nil {
		panic("cannot convert nil")
	}

	jsonData, _ := json.Marshal(f.data)

	return jsonData
}
