package jcsv

import (
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
	myFile, err := ParseCsv(c, hasHeaders)
	if err != nil {
		return nil, err
	}

	return myFile.Json(), err
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
