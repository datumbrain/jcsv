package jcsv

import (
	"encoding/json"
	"strconv"
	"strings"
)

type JSOBObject struct {
	Data map[string]map[string]string
}

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	// TODO: convert JSON data in `j` into CSV format and return
	var JSON JSOBObject
	json.Unmarshal(j, &JSON.Data)
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
	CSVStr := string(c)
	var JSON JSOBObject
	splittedStr := strings.Split(CSVStr, "\n")
	var header []string
	if hasHeaders {
		header = strings.Split(splittedStr[0], ",")
	}
	JSON.Data = make(map[string]map[string]string)
	outerInd := 1
	skipped := false
	for _, row := range splittedStr {
		if row != "" {
			key := "row" + strconv.Itoa(outerInd)
			if hasHeaders == false || skipped == true {
				JSON.Data[key] = make(map[string]string)
			}
			splittedRow := strings.Split(row, ",")
			innerInd := 1
			for _, attribute := range splittedRow {
				if hasHeaders == true {
					if skipped == true {
						JSON.Data[key][header[innerInd-1]] = attribute
					}
				} else {
					JSON.Data[key]["column"+strconv.Itoa(innerInd)] = attribute
				}
				innerInd = innerInd + 1
			}
			skipped = true
		}
		outerInd = outerInd + 1
	}
	jsonData, _ := json.Marshal(JSON.Data)
	return []byte(jsonData), nil
}
