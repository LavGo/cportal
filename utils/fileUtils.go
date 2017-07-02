package utils

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func LoadFile(filename string)[]byte{
	bytes,err:=ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println("[ File Open Error ] : ",err)
		return nil
	}
	return bytes
}
func ParseJson(filename string,jsons  interface{}){
	bytes:=LoadFile(filename)
	if err := json.Unmarshal(bytes, jsons);err!=nil {
		fmt.Println("[ Json Parse Error ] : ",err)
	}
}


