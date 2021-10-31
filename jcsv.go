package jcsv

import (
	"encoding/json"
)

type Json struct {
	Data map[string]string
}

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	// TODO: convert JSON data in `j` into CSV format and return
	var JSON Json
	json.Unmarshal(j, &JSON.Data)
	var CSVData string
	var header string
	for key, element := range JSON.Data {
		header = header + "," + key
		CSVData = CSVData + "," + element
	}
	header = header[1:]
	CSVData = CSVData[1:]
	if addHeaders == true {
		return []byte(header + "\n" + CSVData), nil
	}
	return []byte(CSVData), nil
}

func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {
	// TODO: convert CSV data in `c` into JSON format and return
	return nil, nil
}
