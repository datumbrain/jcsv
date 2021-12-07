package jcsv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ParseJsonFile opens, reads and parses a JSON file and returns a file object
func ParseJsonFile(path string) (*file, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ParseOpenedJsonFile(f)
}

// ParseJsonFile reads and parses an opened JSON file and returns a file object
func ParseOpenedJsonFile(f *os.File) (*file, error) {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return ParseJson(b)
}

// ParseJsonFile parses JSON data and returns a file object
func ParseJson(j []byte) (*file, error) {
	j = bytes.TrimSpace(j)
	if len(j) == 0 {
		return nil, fmt.Errorf("the given input is empty")
	}

	var f file
	if j[0] == '[' {
		err := json.Unmarshal(j, &f.data)
		if err != nil {
			return nil, err
		}
	} else if j[0] == '{' {
		data := make(map[string]interface{})

		err := json.Unmarshal(j, &data)
		if err != nil {
			return nil, err
		}

		f.data = append(f.data, data)
	} else {
		return nil, fmt.Errorf("invalid character '%c' looking for beginning of value", j[0])
	}

	return &f, nil
}
