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
	Data map[string]interface{}
}

func ParseJsonFile(path string) (file, error) {
	//open JSON file
	jsonFile, err := os.Open(path)

	//if an error occurred
	if err != nil {
		fmt.Println(err)
		return file{}, err
	}

	//defer the closing of file till the function ends
	defer jsonFile.Close()

	//read complete json file and store in byteValue
	byteValue, err := ioutil.ReadAll(jsonFile)
	//if an error occurred while reading JSON file

	if err != nil {
		fmt.Println(err)
		return file{}, err
	}

	//create file object
	var myFile file
	//convert JSON read from file in byteValue to "Data" inside file stucture
	json.Unmarshal(byteValue, &myFile.Data)

	return myFile, err
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	//create file object
	var myFile file
	//read complete json file and store in byteValue
	byteValue, err := ioutil.ReadAll(f)

	//if an error occurred while reading JSON file
	if err != nil {
		fmt.Println(err)
		return file{}, nil
	}
	//convert JSON read from file in byteValue to "Data" inside file stucture

	json.Unmarshal(byteValue, &myFile)

	return myFile, err
}

func ParseCsvFile(path string, hasHeaders bool) (file, error) {
	//open CSV file
	csvFile, err := os.Open(path)

	//if an error occurred
	if err != nil {
		return file{}, err
	}

	//defer closing of file
	defer csvFile.Close()

	//get CSV Reader instance
	csvReader := csv.NewReader(csvFile)
	//create an object to hold csv data
	var CSVData [][]string
	//read all data into CSVData
	CSVData, err = csvReader.ReadAll()

	//if an error occurred
	if err != nil {
		return file{}, nil
	}

	//create file object
	var myFile file
	//make map
	myFile.Data = make(map[string]interface{})
	//store csv data into file structure
	for i := 0; i < len(CSVData); i = i + 1 {
		myFile.Data["Record "+fmt.Sprint(i+1+'0')] = CSVData[i]
	}

	return myFile, err
}

func ParseOpenedCsvFile(f *os.File, hasHeaders bool) (file, error) {
	//creat instance of csv reader
	csvReader := csv.NewReader(f)
	var err error

	//create var to hold CSV Data
	var CSVData [][]string
	//read the file and store data in CSVData
	CSVData, err = csvReader.ReadAll()
	//if an error occurred
	if err != nil {
		return file{}, nil
	}

	var myFile file
	//make map
	myFile.Data = make(map[string]interface{})
	//store csv data into file structure
	for i := 0; i < len(CSVData); i = i + 1 {
		myFile.Data["Record "+fmt.Sprint(i+1+'0')] = CSVData[i]
	}
	return myFile, err
}

func (f file) Csv(addHeaders bool) []byte {
	if f.Data == nil {
		return nil
	}

	var CSVFormat string
	for key := range f.Data {
		//if addHeaders is true then add first row else skip first row
		if addHeaders {
			//convert data against key to string
			//most probably it will be a map
			CSVRecord := fmt.Sprintf("%v", f.Data[key])
			//remove square brackets
			CSVRecord = strings.ReplaceAll(CSVRecord, "[", "")
			CSVRecord = strings.ReplaceAll(CSVRecord, "]", "")
			//replace spaces with commas
			CSVRecord = strings.ReplaceAll(CSVRecord, " ", ",")
			//remove map keyword from string
			CSVRecord = strings.ReplaceAll(CSVRecord, "map", "")
			//add new line
			CSVFormat = CSVFormat + CSVRecord + "\n"
		} else {
			addHeaders = true
		}
	}
	return []byte(CSVFormat)
}

func (f file) Json() []byte {
	if f.Data == nil {
		return nil
	}

	//convert JSON data in f.Data to []byte
	jsonData, err := json.Marshal(f.Data)
	//if error occurred
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return jsonData
}
