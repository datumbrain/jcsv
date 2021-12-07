package jcsv

// JsonToCsv convert the given JSON data to CSV data
func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	f, err := ParseJson(j)
	if err != nil {
		return nil, err
	}

	return f.Csv(addHeaders), nil
}

// CsvToJson convert the given CSV data to JSON data
func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {
	f, err := ParseCsv(c, hasHeaders)
	if err != nil {
		return nil, err
	}

	return f.Json(), err
}
