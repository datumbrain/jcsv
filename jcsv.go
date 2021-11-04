package jcsv

import (
	"fmt"
	"strconv"
	"strings"
)

func JsonToCsv(j []byte, addHeaders bool) ([]byte, error) {
	if j==nil{
		return nil,nil
	}
	var header string
	i:=0
	begin:=0
	if addHeaders{
		for ;j[i]!='}';i++{
			if j[i]==':'{
				header+=strings.Trim(string(j[begin:i]),"\"  ") +","
				begin=i+1
			}
			if j[i]=='\n'{
				begin=i+1
			}
		}
		header = strings.TrimRight(header,",")
	}
	var objectFinish bool=false
	fmt.Println(header)
	var result string
	result+=header+"\n"
	i=0
	begin=0
	for ;i<len(j);i++{
		if j[i]==':'{
			objectFinish=false
			begin=i+1
		}
		if j[i]==',' && objectFinish==false{
			result+=strings.Trim(string(j[begin:i]),"\"  ")+","
		}
		if j[i]=='}'{
			result+=strings.Trim(string(j[begin:i]),"\" \n  ")+"\n"
			objectFinish=true
		}
	}


	return []byte(result), nil
}

func CsvToJson(c []byte, hasHeaders bool) ([]byte, error) {
	if c==nil{
		return nil, nil
	}
	var attributes[]string
	begin:=0
	keysCount:=1
	i:=0
	if hasHeaders{
		for ;c[i]!='\n';i++{
			if c[i]==','{
				attributes=append(attributes,string(c[begin:i]))
				begin=i+1
			}else if c[i+1]=='\n'{
				attributes=append(attributes,string(c[begin:i+1]))
				begin=i+1
			}
		}
	}else{
		for ;c[i]!='\n';i++{
			if c[i]==','{
				attributes=append(attributes,"key"+strconv.Itoa(keysCount))
				keysCount++
			}else if c[i+1]=='\n'{
				attributes=append(attributes,"key"+strconv.Itoa(keysCount))
			}
		}
	}
	fmt.Println(attributes)

	var result string
	i++
	count:=0
	for ;i<len(c);i++{
		for ;c[i]!='\n';i++{
			if c[i]==','{
				if count==0{
					result+="{\n"
				}
				result+=string("\t\"")+attributes[count]+string("\":\"")+strings.TrimLeft(string(c[begin:i]),"\n")+string("\",\n")
				begin=i+1
				count++
			}else if c[i+1]=='\n'{
				result+=string("\t\"")+attributes[count]+string("\":\"")+string(c[begin:i+1])+string("\"\n}\n")
				begin=i+1
				count=0
			}
		}
	}

	return []byte(result), nil
}
