package jcsv

import (
	"encoding/json"
	"fmt"
)

type JSOBObject struct {
	Data map[string]map[string]string
}

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	// TODO: convert JSON data in `j` into CSV format and return
	var JSON JSOBObject
	json.Unmarshal(j, &JSON.Data)
	fmt.Println(JSON.Data)
	var CSVData string
	var header string
	isHeaderFormed := false
	for key, _ := range JSON.Data {
		var rowData string
		for innerKey, value := range JSON.Data[key] {
			if !isHeaderFormed {
				header = header + "," + innerKey
			}
			rowData = rowData + "," + string(value)
		}
		isHeaderFormed = true
		CSVData = CSVData + rowData[1:] + "\n"
	}
	header = header[1:]
	if addHeaders == true {
		return []byte(header + "\n" + CSVData), nil
	}
	return []byte(CSVData), nil
}

func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {
	// TODO: convert CSV data in `c` into JSON format and return
	//if(has)
	return nil, nil
}
