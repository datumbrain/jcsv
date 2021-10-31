package jcsv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type file struct {
	// TODO: define variables if needed
	JSONFormat map[string]interface{}
}

func ParseJsonFile(path string) (file, error) {
	// TODO: open and read the given file into your `file` object
	jsonFile, err := os.Open(path)
	var myFile file
	if err != nil {
		fmt.Println(err)
		return myFile, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &myFile.JSONFormat)
	return myFile, nil
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	// TODO: read the given file into your `file` object
	var myFile file
	byteValue, _ := ioutil.ReadAll(f)
	json.Unmarshal([]byte(byteValue), &myFile.JSONFormat)
	return myFile, nil
}

func ParseCsvFile(path string, hasHeaders bool) (file, error) {
	// TODO: open and read the given file into your `file` object
	return file{}, nil
}

func ParseCsvJsonFile(f *os.File, hasHeaders bool) (file, error) {
	// TODO: read the given file into your `file` object
	return file{}, nil
}

func (f file) Csv(addHeaders bool) []byte {
	// TODO: return the file data in CSV format
	return nil
}

func (f file) Json() []byte {
	// TODO: return the file data in JSON format
	return nil
}
