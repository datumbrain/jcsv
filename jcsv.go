package jcsv

import (
	"encoding/json"
	"strconv"
	"strings"
)

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	// convert []byte to file struct
	file, err := ParseJson(j)
	if err != nil {
		return nil, err
	}

	// return converted CSV in []byte form
	return file.Csv(addHeaders), nil
}

func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {

	// split the string on the basis of '\n'
	splittedStr := strings.Split(string(c), "\n")

	// create map
	JSON := make(map[string]map[string]string)

	// for storing header
	var header []string

	// for assigning row no as key
	rowInd := 0

	// if HasHeader is true then first row should be skipped
	// as it only contains heading of each column
	// initially skipHeaderRow is false
	skipHeaderRow := false

	// if hasHeaders is true then split the row 0 that contain headers
	if hasHeaders {
		header = strings.Split(splittedStr[0], ",")
		// now row 0 should be skipped
		skipHeaderRow = true
	}

	// for each comma separated string in splittedStr
	for _, row := range splittedStr {
		// if row is not blank and skipHeaderRow is false
		if row != "" && !skipHeaderRow {
			// create key for row
			key := "row" + strconv.Itoa(rowInd)

			// create map for row
			JSON[key] = make(map[string]string)

			// split string on the basis of comma
			splittedRow := strings.Split(row, ",")

			colInd := 0
			for _, attribute := range splittedRow {

				if hasHeaders {
					// if hasHeader then create that specific header as row key
					JSON[key][header[colInd]] = attribute
				} else {
					// if hasHeader then create that specific header as row key
					JSON[key]["key"+strconv.Itoa(colInd)] = attribute
				}

				colInd = colInd + 1
			}
		}
		skipHeaderRow = false
		rowInd = rowInd + 1
	}
	// convert map to JSON
	jsonData, _ := json.Marshal(JSON)

	return []byte(jsonData), nil
}

func toByteArray(data [][]string) []byte {
	var convertedData string

	for _, strArray := range data {
		// separate []string at every index of [][]string by "\n"
		// separate string at every index of []string by "~"
		// This formating will be considered while converted [][] string again to []byte
		convertedData = convertedData + strings.Join(strArray, "~") + "\n"
	}

	return []byte(convertedData)
}

func toArrayOfArrayOfString(data []byte) [][]string {
	var convertedData [][]string

	// spliting on the basis of "\n" that will give the actual []string at every index of [][]string
	strArray := strings.Split(string(data), "\n")
	for _, str := range strArray {
		// spliting on the basis of "~" that will give the actual string at every index of []string
		convertedData = append(convertedData, strings.Split(str, "~"))
	}
	convertedData = convertedData[0 : len(convertedData)-1]

	return convertedData
}
