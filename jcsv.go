package jcsv

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	// convert []byte to file struct
	file, err := ParseJson(j)
	if err != nil {
		return nil, err
	}

	// return converted CSV in []byte form
	return file.Csv(addHeaders), nil
}

func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {
	myFile, err := ParseCsv(c, hasHeaders)
	if err != nil {
		return nil, err
	}

	return myFile.Json(), err
}
