package jcsv

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type file struct {
	data []map[string]interface{}
}

func ParseJsonFile(path string) (file, error) {
	// open JSON file
	jsonFile, err := os.Open(path)

	if err != nil {
		return file{}, err
	}

	defer jsonFile.Close()

	return ParseOpenedJsonFile(jsonFile)
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	// read complete json file and store in byteValue
	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		return file{}, err
	}

	return ParseJson(byteValue)
}

func ParseCsvFile(path string, hasHeaders bool) (file, error) {
	// open CSV file
	csvFile, err := os.Open(path)

	if err != nil {
		return file{}, err
	}

	defer csvFile.Close()

	return ParseOpenedCsvFile(csvFile, hasHeaders)
}

func ParseOpenedCsvFile(f *os.File, hasHeaders bool) (file, error) {
	// parsing csv data
	//data, err := csv.NewReader(f).ReadAll()
	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		return file{}, err
	}

	if len(byteValue) == 0 {
		return file{}, nil
	}

	return ParseCsv(byteValue, hasHeaders)
}

func (f file) Csv(addHeaders bool) []byte {
	if f.data == nil {
		panic("cannot convert nil")
	}

	var myCSV [][]string
	var header []string
	// parse csv headers
	for key, _ := range f.data[0] {
		header = append(header, key)
	}
	if addHeaders {
		myCSV = append(myCSV, header)
	}

	// Parse csv values
	for _, rowData := range f.data {
		var row []string
		for _, value := range rowData {
			val := fmt.Sprintf("%v", value)

			val = strings.ReplaceAll(val, "map", "")
			val = strings.ReplaceAll(val, "[", "")
			val = strings.ReplaceAll(val, "]", "")

			row = append(row, val)
		}
		myCSV = append(myCSV, row)
	}

	// write csv data
	var buffer bytes.Buffer
	csvWriter := csv.NewWriter(&buffer)
	csvWriter.WriteAll(myCSV)
	fmt.Println(string(buffer.Bytes()))

	return buffer.Bytes()
}

func (f file) Json() []byte {
	if f.data == nil {
		panic("cannot convert nil")
	}

	// convert JSON data to []byte
	jsonData, _ := json.Marshal(f.data)

	return jsonData
}

func ParseJson(j []byte) (file, error) {
	var myFile file

	// convert JSON read from file to struct
	err := json.Unmarshal(j, &myFile.data)

	if err != nil {
		// If JSON data was not in the form of an array this will try to unmarshal it as an object
		// and append it in the array
		var data map[string]interface{}

		err = json.Unmarshal(j, &data)
		if err != nil {
			return file{}, nil
		}

		myFile.data = append(myFile.data, data)
	}

	return myFile, err
}

func ParseCsv(c []byte, hasHeaders bool) (file, error) {
	//data := toArrayOfArrayOfString(c)
	//buffer :=
	data, err := csv.NewReader(bytes.NewBuffer(c)).ReadAll()
	if err != nil {
		return file{}, err
	}

	var myFile file

	// parsing csv headers
	i := 0
	var header []string
	if hasHeaders {
		header = data[0]
		i = 1
	} else {
		for i := range data[0] {
			header = append(header, "key"+strconv.Itoa(i))
		}
	}

	// parsing csv values
	for ; i < len(data); i = i + 1 {
		row := make(map[string]interface{})

		for j := 0; j < len(data[i]); j++ {
			row[header[j]] = data[i][j]
		}

		myFile.data = append(myFile.data, row)
	}

	return myFile, nil
}
