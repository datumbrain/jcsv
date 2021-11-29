package jcsv

import (
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
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return file{}, err
	}

	if len(data) == 0 {
		return file{}, nil
	}

	return ParseCsv(toByteArray(data), hasHeaders)
}

func (f file) Csv(addHeaders bool) []byte {
	if f.data == nil {
		panic("cannot convert nil")
	}

	var CSVFormat string
	isHeaderFormed := false
	var header string

	for key := range f.data {

		CSVRecord := fmt.Sprintf("%v", f.data[key])

		// remove square brackets
		CSVRecord = strings.ReplaceAll(CSVRecord, "[", "")
		CSVRecord = strings.ReplaceAll(CSVRecord, "]", "")
		// replace spaces with commas
		CSVRecord = strings.ReplaceAll(CSVRecord, " ", ",")
		// remove map keyword from string
		CSVRecord = strings.ReplaceAll(CSVRecord, "map", "")
		// replace all keys
		// user1 , user2 and so on
		for mapKey := range f.data[key] {
			CSVRecord = strings.ReplaceAll(CSVRecord, mapKey+":", "")
			if !isHeaderFormed {
				header = header + "," + mapKey
			}
		}
		if len(CSVRecord) != 0 {
			isHeaderFormed = true
		}
		// CSVFormat = CSVFormat + CSVRecord + "\n"
		// wherever newline is meant to be inserted that index contains ,:
		// ......csv.....,key:.......csv
		// when key is removed ,: remains
		CSVRecord = strings.ReplaceAll(CSVRecord, ",:", ",")
		CSVFormat = CSVFormat + CSVRecord + "\n"
	}

	if addHeaders {
		CSVFormat = header[1:] + "\n" + CSVFormat
	}
	return []byte(CSVFormat)
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

	data := toArrayOfArrayOfString(c)

	var myFile file

	// store csv data into file structure
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

	//store data in structure
	for ; i < len(data); i = i + 1 {
		row := make(map[string]interface{})

		for j := 0; j < len(data[i]); j++ {
			row[header[j]] = data[i][j]
		}

		myFile.data = append(myFile.data, row)
	}

	return myFile, nil

}
