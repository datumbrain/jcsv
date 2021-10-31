package main

import (
	"io/ioutil"
	"os"
)

type file struct {
	data []uint8
	fileName string
}

func ParseJsonFile(path string) (file, error) {
	var f file
	f.fileName=path
	f.data=nil
	fileStream,fileOpenError := os.Open(f.fileName)
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
