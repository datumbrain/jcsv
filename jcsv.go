package jcsv

import (
	"encoding/json"
	"strconv"
	"strings"
)

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	//create file object
	var JSON file

	//convert data in j to Data
	err := json.Unmarshal(j, &JSON.data)
	if err != nil {
		return nil, err
	}

	//return converted CSV in []byte form
	return JSON.Csv(addHeaders), nil
}

func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {

	//split the string on the basis of '\n'
	splittedStr := strings.Split(string(c), "\n")

	//create map
	JSON := make(map[string]map[string]string)

	//for storing header
	var header []string

	//for assigning row no as key
	rowInd := 0

	//if HasHeader is true then first row should be skipped
	//as it only contains heading of each column
	//initially skipHeaderRow is false
	skipHeaderRow := false

	//if hasHeaders is true then split the row 0 that contain headers
	if hasHeaders {
		header = strings.Split(splittedStr[0], ",")
		//now row 0 should be skipped
		skipHeaderRow = true
	}

	//for each comma separated string in splittedStr
	for _, row := range splittedStr {
		//if row is not blank and skipHeaderRow is false
		if row != "" && !skipHeaderRow {
			//create key for row
			key := "row" + strconv.Itoa(rowInd)

			//create map for row
			JSON[key] = make(map[string]string)

			//split string on the basis of comma
			splittedRow := strings.Split(row, ",")

			colInd := 0
			for _, attribute := range splittedRow {

				if hasHeaders {
					//if hasHeader then create that specific header as row key
					JSON[key][header[colInd]] = attribute
				} else {
					//if hasHeader then create that specific header as row key
					JSON[key]["key"+strconv.Itoa(colInd)] = attribute
				}

				colInd = colInd + 1
			}
		}
		skipHeaderRow = false
		rowInd = rowInd + 1
	}
	//convert map to JSON
	jsonData, _ := json.Marshal(JSON)

	return []byte(jsonData), nil
}
