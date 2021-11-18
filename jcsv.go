package jcsv

import (
	"encoding/json"
	"strconv"
	"strings"
)

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	var JSON file
	parseError := json.Unmarshal(j, &JSON.data) //projecting []byte to json.data which is []map[string]interface{}

	if parseError != nil {
		return nil, parseError
	}

	return JSON.Csv(addHeaders), nil
}

func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {
	splittedStr := strings.Split(string(c), "\n") //typecasting 'c' to string and splitting it on basis of new line
	JSON := make(map[string]map[string]string)
	var header []string
	rowInd := 0
	skipHeaderRow := false
	if hasHeaders { //if hasHeader is true, first row of given csv data must be header which will be keys in json
		header = strings.Split(splittedStr[0], ",")
		skipHeaderRow = true
	}

	for _, row := range splittedStr {
		if row != "" && !skipHeaderRow { //skipHeaderRow==false means we don't have keys, we will make them
			key := strconv.Itoa(rowInd) //key := "row" + strconv.Itoa(rowInd)  //
			JSON[key] = make(map[string]string)
			splittedRow := strings.Split(row, ",") //dividing complete row on basis of ',' these
			colInd := 0                            //splitted strings will be values for json

			for _, attribute := range splittedRow {

				if hasHeaders { //using given headers as keys
					JSON[key][header[colInd]] = attribute
				} else { //making our own keys
					JSON[key]["key"+strconv.Itoa(colInd)] = attribute
				}

				colInd = colInd + 1
			}
		}

		skipHeaderRow = false
		rowInd = rowInd + 1
	}

	jsonData, _ := json.Marshal(JSON) //coverting the json data we built to []byte

	return []byte(jsonData), nil
}
