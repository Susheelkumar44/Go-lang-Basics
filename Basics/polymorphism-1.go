package main

import (
	"fmt"
	"reflect"
	"errors"
)

func add(var1, var2 interface{}) (interface{}, error) {

	typeOfVar1 := reflect.TypeOf(var1).Name()
	typeOfVar2 := reflect.TypeOf(var2).Name()
	
	if typeOfVar1 == "int" && typeOfVar2 == "int" {
		var3 := var1.(int)
		var4 := var2.(int)
		return var3 + var4,  nil
	} else if typeOfVar1 == "string" && typeOfVar2 == "string" {
		var5 := var1.(string)
		var6 := var2.(string) 
		return var5 + var6, nil
	}
		
	return nil, errors.New("Data are of different types")
}

func main() {

	result, err := add("susheel", "kumar")
	if err!= nil {
		fmt.Println(err)
	}
	fmt.Println("Result is", result)
	result, err = add(4, 6)
	if err!= nil {
		fmt.Println(err)
	}
	fmt.Println("Result is", result)
	result, err =add(1, "string")
	if err!= nil {
		fmt.Println("Error!!!! ",err)
	}
}