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
		fmt.Println(err)
		return file{}, err
	}

	defer jsonFile.Close()

	// read complete json file and store in byteValue
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
		return file{}, err
	}

	var myFile file
	// convert JSON read from file in byteValue to "Data" inside file stucture
	err = json.Unmarshal(byteValue, &myFile.data)

	return myFile, err
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	var myFile file

	// read complete json file and store in byteValue
	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
		return file{}, nil
	}

	// convert JSON read from file in byteValue to "Data" inside file stucture
	json.Unmarshal(byteValue, &myFile.data)

	return myFile, err
}

func ParseCsvFile(path string, hasHeaders bool) (file, error) {
	// open CSV file
	csvFile, err := os.Open(path)

	if err != nil {
		return file{}, err
	}

	defer csvFile.Close()

	// get CSV Reader instance
	csvReader := csv.NewReader(csvFile)
	// create an object to hold csv data
	var CSVData [][]string
	// read all data into CSVData
	CSVData, err = csvReader.ReadAll()
	// if an error occurred
	if err != nil {
		return file{}, nil
	}

	var myFile file
	myFile.data = make([]map[string]interface{}, len(CSVData))

	// store csv data into file structure
	i := 0
	var header []string
	if hasHeaders {
		i = 1
		header = CSVData[0]
	}
	for ; i < len(CSVData); i = i + 1 {
		if hasHeaders {
			myFile.data[i-1] = make(map[string]interface{})
			for j := 0; j < len(CSVData[i]); j++ {
				myFile.data[i-1][header[j]] = CSVData[i][j]
			}
		} else {
			myFile.data[i] = make(map[string]interface{})
			for j := 0; j < len(CSVData[i]); j++ {
				myFile.data[i]["key"+fmt.Sprint(j)] = CSVData[i][j]
			}
		}
	}
	return myFile, err
}

func ParseOpenedCsvFile(f *os.File, hasHeaders bool) (file, error) {
	// creat instance of csv reader
	csvReader := csv.NewReader(f)
	var err error

	// create var to hold CSV Data
	var CSVData [][]string
	// read the file and store data in CSVData
	CSVData, err = csvReader.ReadAll()
	// if an error occurred
	if err != nil {
		return file{}, nil
	}

	var myFile file
	myFile.data = make([]map[string]interface{}, len(CSVData))

	// store csv data into file structure
	i := 0
	var header []string
	if hasHeaders {
		i = 1
		header = CSVData[0]
	}
	for ; i < len(CSVData); i = i + 1 {
		if hasHeaders {
			myFile.data[i-1] = make(map[string]interface{})
			for j := 0; j < len(CSVData[i]); j++ {
				myFile.data[i-1][header[j]] = CSVData[i][j]
			}
		} else {
			myFile.data[i] = make(map[string]interface{})
			for j := 0; j < len(CSVData[i]); j++ {
				myFile.data[i]["key"+fmt.Sprint(j)] = CSVData[i][j]
			}
		}
	}
	return myFile, err
}

func (f file) Csv(addHeaders bool) []byte {
	if f.data == nil {
		return nil
	}
	var CSVFormat string
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

		if addHeaders {
			// replace all keys
			// user1 , user2 and so on
			for mapKey := range f.data[key] {
				if key == 0 {
					header = header + "," + mapKey
				}
				CSVRecord = strings.ReplaceAll(CSVRecord, mapKey+":", "")
			}
			// wherever newline is meant to be inserted that index contains ,:
			// ......csv.....,key:.......csv
			// when key is removed ,: remains
			CSVRecord = strings.ReplaceAll(CSVRecord, ",:", ",")
		} else {
			for mapKey := range f.data[key] {
				CSVRecord = strings.ReplaceAll(CSVRecord, mapKey+":", "")
			}
		}
		CSVFormat = CSVFormat + CSVRecord + "\n"
	}
	if addHeaders {
		CSVFormat = header[1:] + "\n" + CSVFormat
	}
	return []byte(CSVFormat)
}

func (f file) Json() []byte {
	if f.data == nil {
		return nil
	}

	// convert JSON data in f.Data to []byte
	jsonData, err := json.Marshal(f.data)
	// if error occurred
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return jsonData
}
