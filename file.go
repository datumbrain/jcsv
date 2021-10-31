package jcsv

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
)

type file struct {
	// TODO: define variables if needed
	JSONFormat map[string]interface{}
	CSVFormat  struct {
		Data [][]string
		//if "headers==true", data[0] will contain headers
		HasHeaders bool
	}
}

func (f *file) initialize() {
	f.JSONFormat = nil
	f.CSVFormat.Data = nil
	f.CSVFormat.HasHeaders = false
}
func ParseJsonFile(path string) (file, error) {
	// TODO: open and read the given file into your `file` object
	jsonFile, err := os.Open(path)
	var myFile file
	myFile.initialize()
	if err != nil {
		return myFile, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err == nil {
		json.Unmarshal([]byte(byteValue), &myFile.JSONFormat)
	}
	return myFile, err
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	// TODO: read the given file into your `file` object
	var myFile file
	myFile.initialize()
	byteValue, err := ioutil.ReadAll(f)
	if err == nil {
		json.Unmarshal([]byte(byteValue), &myFile.JSONFormat)
	}
	return myFile, err
}

func ParseCsvFile(path string, hasHeaders bool) (file, error) {
	// TODO: open and read the given file into your `file` object
	csvFile, err := os.Open((path))
	var myFile file
	myFile.initialize()
	if err != nil {
		return myFile, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	myFile.CSVFormat.Data, err = csvReader.ReadAll()
	if err == nil {
		myFile.CSVFormat.HasHeaders = hasHeaders
		myFile.JSONFormat = nil
	}
	return myFile, err
}

func ParseOpenedCsvFile(f *os.File, hasHeaders bool) (file, error) {
	// TODO: read the given file into your `file` object
	var myFile file
	myFile.initialize()
	var err error
	csvReader := csv.NewReader(f)
	myFile.CSVFormat.Data, err = csvReader.ReadAll()
	if err == nil {
		myFile.CSVFormat.HasHeaders = hasHeaders
		myFile.JSONFormat = nil
	}
	return myFile, err
}

func (f file) Csv(addHeaders bool) []byte {
	// TODO: return the file data in CSV format
	return nil
}

func (f file) Json() []byte {
	// TODO: return the file data in JSON format
	return nil
}
