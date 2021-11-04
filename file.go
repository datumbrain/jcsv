package jcsv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type file struct {
	Data [][]string
	//if "headers==true", data[0] will contain headers
	HasHeaders bool
}

func ParseJsonFile(path string) (file, error) {
	// TODO: open and read the given file into your `file` object
	jsonFile, err := os.Open(path)
	if err != nil {
		return file{}, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	var myFile file
	if err != nil {
		fmt.Println(err)
		return file{}, err
	}
	json.Unmarshal(byteValue, &myFile.Data)
	return myFile, err
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	// TODO: read the given file into your `file` object
	var myFile file
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return file{}, nil

	}
	json.Unmarshal(byteValue, &myFile)
	return myFile, err
}

func ParseCsvFile(path string, hasHeaders bool) (file, error) {
	// TODO: open and read the given file into your `file` object
	csvFile, err := os.Open(path)
	if err != nil {
		return file{}, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	var myFile file
	myFile.Data, err = csvReader.ReadAll()
	if err != nil {
		return file{}, nil
	}
	myFile.HasHeaders = hasHeaders
	return myFile, err
}

func ParseOpenedCsvFile(f *os.File, hasHeaders bool) (file, error) {
	// TODO: read the given file into your `file` object
	csvReader := csv.NewReader(f)
	var myFile file
	var err error
	myFile.Data, err = csvReader.ReadAll()
	if err != nil {
		return file{}, nil
	}
	myFile.HasHeaders = hasHeaders
	return myFile, err
}

func (f file) Csv(addHeaders bool) []byte {
	// TODO: return the file data in CSV format
	if f.Data == nil {
		return nil
	}
	var CSVFormat string
	var i int
	if addHeaders == true {
		i = 0
	} else {
		i = 1
	}
	for ; i < len(f.Data); i = i + 1 {
		for j := 0; j < len(f.Data[i]); j = j + 1 {
			CSVFormat = CSVFormat + (f.Data[i][j])
			if j != len(f.Data[i])-1 {
				CSVFormat = CSVFormat + ","
			}
		}
		CSVFormat = CSVFormat + "\n"
	}
	return []byte(CSVFormat)
}

func (f file) Json() []byte {
	// TODO: return the file data in JSON format
	if f.Data == nil {
		return nil
	}
	jsonData, err := json.Marshal(f.Data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonData
}
