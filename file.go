package jcsv

import "os"

type file struct {
	// TODO: define variables if needed
}

func ParseJsonFile(path string) (file, error) {
	// TODO: open and read the given file into your `file` object
	return file{}, nil
}

func ParseOpenedJsonFile(f *os.File) (file, error) {
	// TODO: read the given file into your `file` object
	return file{}, nil
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
