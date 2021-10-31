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
	//fil,err := os.Open("C:\\Users\\sulem\\OneDrive\\Desktop\\data.json")
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//data,readError := ParseOpenedJsonFile(fil)
	//if readError!=nil{
	//	fmt.Println(readError)
	//}
	//fmt.Println(string(data.data))



	//f,err:=ParseJsonFile("data.json")
	//if err==nil{
	//	fmt.Println(string(f.data))
	//}

	//f,_ := os.Open("data.json")
	//var data []uint8
	//data,_ = ioutil.ReadAll(f)
	//for i,v:= range data{
	//	fmt.Println(i,string(v))
	//}
	//fmt.Printf("%T--%s",data,string(data))


	//csvFile, err := os.Open("C:\\Users\\sulem\\OneDrive\\Desktop\\data.csv")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//csvLines,_ := csv.NewReader(csvFile).ReadAll()
	//fmt.Println(csvLines)
	////data := make(map[string] []uint8)
	//var data [][]string
	//for i,lin:=range csvLines{
	//	if i==0{
	//		continue
	//	}
	//	data = append(data, lin)
	//}
	//fmt.Println(data)
	//f,err:=ParseCsvFile("C:\\Users\\sulem\\OneDrive\\Desktop\\data.csv",true)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(f.csvData)
	f,err:=os.Open("C:\\Users\\sulem\\OneDrive\\Desktop\\data.csv")
	if err!=nil{
		fmt.Println(err)
	}
	fil,csvErr:=ParseOpenedCsvFile(f,true)
	if csvErr!=nil{
		fmt.Println(csvErr)
	}
	fmt.Println(fil.csvData)
}