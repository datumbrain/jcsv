package main

import (
	"FYP/GO_ASSIGNMENT1/jcsv"
	"fmt"
)

func main() {

	myFile, _ := jcsv.ParseCsvFile("users.csv", true)
	myCsv, _ := jcsv.CsvToJson(myFile.Csv(true), true)
	fmt.Println(myCsv)
	//fmt.Println(toArrayOfArrayOfString(toByteArray(c)))
}

// func toByteArray(data [][]string) []byte {
// 	var convertedData string

// 	for _, strArray := range data {
// 		convertedData = convertedData + strings.Join(strArray, "~") + "\n"
// 	}

// 	return []byte(convertedData)
// }

// func toArrayOfArrayOfString(data []byte) [][]string {
// 	var convertedData [][]string
// 	strArray := strings.Split(string(data), "\n")
// 	for _, str := range strArray {
// 		convertedData = append(convertedData, strings.Split(str, "~"))
// 	}
// 	convertedData = convertedData[0 : len(convertedData)-1]
// 	return convertedData
// }
