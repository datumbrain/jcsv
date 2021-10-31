package main

import (
	"fmt"
	"os"
)

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	// TODO: convert JSON data in `j` into CSV format and return
	return nil, nil
}

func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {
	// TODO: convert CSV data in `c` into JSON format and return
	return nil, nil
}

func main() {
	fil,err := os.Open("C:\\Users\\sulem\\OneDrive\\Desktop\\data.json")
	if err!=nil{
		fmt.Println(err)
	}
	data,readError := ParseOpenedJsonFile(fil)
	if readError!=nil{
		fmt.Println(readError)
	}
	fmt.Println(string(data.data))
}