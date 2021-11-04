package jcsv


import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type file struct {
	jsonData []uint8
	csvData [][]string
	csvHeader []string
}

func ParseJsonFile(path string) (file, error) {
	var fileObj file
	fileStream,fileOpenError := os.Open(path)
	if fileOpenError!=nil{
		return file{},fileOpenError
	}
	var fileReadError error
	fileObj.jsonData,fileReadError=ioutil.ReadAll(fileStream)
	fileStream.Close()
	if fileReadError!=nil {
		return file{},fileReadError
	}
	return fileObj,nil
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	var fileRead file
	var fileReadError error
	fileRead.jsonData,fileReadError=ioutil.ReadAll(f)
	if fileReadError!=nil {
		return file{},fileReadError
	}
	return fileRead,nil
}

func ParseCsvFile(path string, hasHeaders bool) (file, error) {
	var returnFile file
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
	if hasHeaders {
		for j := 0; j < len(readBuffer[0]); j++ {
			returnFile.csvHeader = append(returnFile.csvHeader, readBuffer[0][j])
		}
		i=1
	}else {
		for j:=0;j<len(readBuffer[0]);j++{
			returnFile.csvHeader=append(returnFile.csvHeader,"key" + strconv.Itoa(j))
		}
		i=0
	}
	for ;i<len(readBuffer);i++{
		returnFile.csvData = append(returnFile.csvData, readBuffer[i])
	}
	csvFile.Close()
	return  returnFile,nil
}

func ParseOpenedCsvFile(f *os.File, hasHeaders bool) (file, error) {
	var fileObj file
	readBuffer,readError := csv.NewReader(f).ReadAll()
	if readError!=nil{
		return file{},readError
	}
	var i int
	if hasHeaders {
		for j := 0; j < len(readBuffer[0]); j++ {
			fileObj.csvHeader = append(fileObj.csvHeader, readBuffer[0][j])
		}
		i=1
	}else {
		for j:=0;j<len(readBuffer[0]);j++{
			fileObj.csvHeader=append(fileObj.csvHeader,"key" + strconv.Itoa(j))
		}
		i=0
	}
	for ;i<len(readBuffer);i++{
		fileObj.csvData = append(fileObj.csvData, readBuffer[i])
	}
	return  fileObj,nil
}

func (f file) Csv(addHeaders bool) []byte {
	var dataToBeReturned []byte
	if addHeaders == true{
		for j:=0;j<len(f.csvHeader);j++{
			for k:=0;k<len(f.csvHeader[j]);k++{
				dataToBeReturned = append(dataToBeReturned , f.csvHeader[j][k])
			}
			if j<len(f.csvHeader)-1 {
				dataToBeReturned = append(dataToBeReturned, ',')
			}
		}
	}
	dataToBeReturned= append(dataToBeReturned, '\n')
	for i:=0;i< len(f.csvData);i++{
		for j:=0;j<len(f.csvData[i]);j++{
			for k:=0;k<len(f.csvData[i][j]);k++{
				dataToBeReturned = append(dataToBeReturned , f.csvData[i][j][k])
			}
			if j<len(f.csvData[i]) - 1 {
				dataToBeReturned = append(dataToBeReturned, ',')
			}
		}
		dataToBeReturned=append(dataToBeReturned,'\n')
	}
	return dataToBeReturned

}

func (f file) Json() []byte {
	return f.jsonData
}
