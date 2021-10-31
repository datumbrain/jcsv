package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
)

type file struct {
	data []uint8
	csvData [][]string
}

func ParseJsonFile(path string) (file, error) {
	var f file
	f.data=nil
	fileStream,fileOpenError := os.Open(path)
	if fileOpenError!=nil{
		return file{},fileOpenError
	}
	var fileReadError error
	f.data,fileReadError=ioutil.ReadAll(fileStream)
	if fileReadError!=nil {
		return file{},fileReadError
	}
	return f,nil
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	var fileRead file
	var fileReadError error
	fileRead.data,fileReadError=ioutil.ReadAll(f)
	if fileReadError!=nil {
		return file{},fileReadError
	}
	return fileRead,nil
}

func ParseCsvFile(path string, hasHeaders bool) (file, error) {
	var f file
	f.data=nil
	f.csvData=nil
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return file{},err
	}
	readBuffer,readError := csv.NewReader(csvFile).ReadAll()
	if readError!=nil{
		return file{},readError
	}
	var i int
	if hasHeaders{
		i=1
	}else{
		i=0
	}
	for ;i<len(readBuffer);i++{
		f.csvData = append(f.csvData, readBuffer[i])
	}
	return  f,nil
}

func ParseOpenedCsvFile(f *os.File, hasHeaders bool) (file, error) {
	var fileObj file
	fileObj.data=nil
	fileObj.csvData=nil
	readBuffer,readError := csv.NewReader(f).ReadAll()
	if readError!=nil{
		return file{},readError
	}
	var i int
	if hasHeaders{
		i=1
	}else{
		i=0
	}
	for ;i<len(readBuffer);i++{
		fileObj.csvData = append(fileObj.csvData, readBuffer[i])
	}
	return  fileObj,nil
}

func (f file) Csv(addHeaders bool) []byte {
	// TODO: return the file data in CSV format
	return nil
}

func (f file) Json() []byte {
	// TODO: return the file data in JSON format
	return nil
}
