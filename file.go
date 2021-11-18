package jcsv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type file struct {
	data []map[string]interface{} //a 1D-array of maps having string keys and empty interface values
}

func ParseJsonFile(path string) (file, error) {
	jsonFile, openError := os.Open(path)
	if openError != nil {
		return file{}, openError
	}

	defer jsonFile.Close()
	return ParseOpenedJsonFile(jsonFile)
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	byteValue, readError := ioutil.ReadAll(f) //ReadAll gives read data in []byte

	if readError != nil {
		return file{}, readError
	}

	var myFile file
	readError = json.Unmarshal(byteValue, &myFile.data) //Unmarshal projects/casts []byte data on a variable wrt its datatype

	if readError != nil {
		var JSON map[string]interface{} //Unmarshall can not project []byte to []map[string] if JSON is limited to only one object
		readError = json.Unmarshal(byteValue, &JSON)

		if readError != nil {
			return file{}, nil
		}

		myFile.data = append(myFile.data, JSON) //myFile.data is of type []map[string]interface{} whereas JSON is of map[string]interface{}
	}

	return myFile, readError
}

func ParseCsvFile(path string, hasHeaders bool) (file, error) {
	csvFile, openErr := os.Open(path)

	if openErr != nil {
		return file{}, openErr
	}

	defer csvFile.Close()
	return ParseOpenedCsvFile(csvFile, hasHeaders)
}

func ParseOpenedCsvFile(f *os.File, hasHeaders bool) (file, error) {
	csvReader := csv.NewReader(f)
	CSVData, readError := csvReader.ReadAll() // csv ReadALl returns data in form of 2D-Array of String

	if readError != nil {
		return file{}, readError
	}

	var myFile file
	i := 0
	var header []string

	if hasHeaders {
		header = CSVData[0]
		i = 1
	} else {
		for i := 0; i < len(CSVData[0]); i = i + 1 {
			header = append(header, "key"+fmt.Sprint(i)) //making my own keys/attribute names
		}
	}

	for ; i < len(CSVData); i = i + 1 {
		row := make(map[string]interface{})    //var row map[string]interface{}
		for j := 0; j < len(CSVData[i]); j++ { //accessing every string object in a single row of csv data
			row[header[j]] = CSVData[i][j] //storing csv data in form of key-value pairs
		}
		myFile.data = append(myFile.data, row)
	}
	return myFile, readError
}

func (f file) Csv(addHeaders bool) []byte {
	if f.data == nil {
		return nil
	}

	var CSVFormat string
	isHeaderFormed := false
	var header string

	for key := range f.data {
		CSVRecord := fmt.Sprintf("%v", f.data[key])         //stringify data of given variable
		CSVRecord = strings.ReplaceAll(CSVRecord, "[", "")  //f.data holds data as 1D-array of maps, and
		CSVRecord = strings.ReplaceAll(CSVRecord, "]", "")  // here we are filtering out Json markups
		CSVRecord = strings.ReplaceAll(CSVRecord, " ", ",") //to convert it to csv
		CSVRecord = strings.ReplaceAll(CSVRecord, "map", "")

		for mapKey := range f.data[key] {
			CSVRecord = strings.ReplaceAll(CSVRecord, mapKey+":", "")
			if !isHeaderFormed {
				header = header + "," + mapKey
			}
		}

		if len(CSVRecord) != 0 {
			isHeaderFormed = true
		}

		CSVRecord = strings.ReplaceAll(CSVRecord, ",:", ",")
		CSVFormat = CSVFormat + CSVRecord + "\n"
	}

	if addHeaders {
		CSVFormat = header[1:] + "\n" + CSVFormat //header[0] contains ','
	}
	return []byte(CSVFormat)
}

func (f file) Json() []byte {
	if f.data == nil {
		return nil
	}
	jsonData, err := json.Marshal(f.data) //it projects data from a variable to []byte

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonData
}
