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

	var myFile file
	// convert JSON read from file in byteValue to "Data" inside file stucture
	err = json.Unmarshal(byteValue, &myFile.data)
	if err != nil {
		//if JSON data was not in the form of array
		//the Unmartial it in map[string]interface{}
		//and append it in the array in file
		//so first index of array contains the JSON object
		var JSON map[string]interface{}
		err = json.Unmarshal(byteValue, &JSON)
		if err != nil {
			return file{}, nil
		}
		myFile.data = append(myFile.data, JSON)
	}
	return myFile, err
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

	// get CSV Reader instance
	csvReader := csv.NewReader(f)
	// create an object to hold csv data
	var CSVData [][]string
	var err error
	// read all data into CSVData
	CSVData, err = csvReader.ReadAll()
	// if an error occurred
	if err != nil {
		return file{}, nil
	}

	var myFile file

	// store csv data into file structure
	i := 0
	var header []string
	if hasHeaders {
		header = CSVData[0]
		i = 1
	} else {
		for i := 0; i < len(CSVData[0]); i = i + 1 {
			header = append(header, "key"+fmt.Sprint(i))
		}
	}

	for ; i < len(CSVData); i = i + 1 {
		row := make(map[string]interface{})
		for j := 0; j < len(CSVData[i]); j++ {
			row[header[j]] = CSVData[i][j]
		}
		myFile.data = append(myFile.data, row)
	}
	return myFile, err
}

func (f file) Csv(addHeaders bool) []byte {
	if f.data == nil {
		return nil
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
	fmt.Println(CSVFormat)
	return []byte(CSVFormat)
}

func (f file) Json() []byte {
	if f.data == nil {
		return nil
	}

	// convert JSON data in f.Data to []byte
	jsonData, err := json.Marshal(f.data)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonData
}
