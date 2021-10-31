package main

import "fmt"

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	// TODO: convert JSON data in `j` into CSV format and return
	return nil, nil
}

func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {
	// TODO: convert CSV data in `c` into JSON format and return
	return nil, nil
}

func main() {
	f,err:=ParseJsonFile("data.json")
	if err==nil{
		fmt.Println(string(f.data))
	}
}