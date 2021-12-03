package main

import (
	"FYP/GO_ASSIGNMENT1/jcsv"
	"fmt"
)

func main() {

	myFile, err := jcsv.ParseCsvFile("users.csv", true)
	myFile.Csv(true)
	fmt.Println(err, "=======================")

	// var a []map[string]interface{}
	// var b map[string]interface{}
	// b = make(map[string]interface{})
	// b["name"] = "Ahsan"
	// b["roll-no"] = "BCSF18M007"
	// c := make(map[string]interface{})
	// a = append(a, b)
	// c["name"] = "Suleman"
	// c["roll-no"] = "BCSF18M002"
	// a = append(a, c)

	//w := csv.NewWriter(os.Stdout)
	// for i := range a {
	// 	//var row []string
	// 	for _, value := range a[i] {
	// 		fmt.Println(fmt.Sprintf("%v", value))

	// 		//row = append(row, value)
	// 	}

	// 	//err = w.Write(strings.Join(row, ","))
	// }

	// records := [][]string{
	// 	{"Name", "Roll No", "ID"},
	// 	{"ahsan", "BCSF18M007", "7"},
	// 	{"suleman", "BCSF18M002", "2"}}

	// var b []byte
	// buf := bytes.NewBuffer(b)
	// w := csv.NewWriter(buf)
	// w.WriteAll(records)

	// fmt.Println(w, string(buf.Bytes()))

	// var buffer bytes.Buffer
	// w := csv.NewWriter(&buffer)
	// err := w.WriteAll(records)
	// w.Flush()
	//fmt.Println(buffer.Bytes(), err)
}
