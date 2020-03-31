package main

import (
	"errors"
	"fmt"
	"reflect"
)

type data struct {
	var1, var2 interface{}
}

type Addition interface {
	add() (interface{}, error)
}

func (d *data) add() (interface{}, error) {

	if (reflect.TypeOf(d.var1).Name()) == (reflect.TypeOf(d.var2).Name()) {
		switch v := d.var1.(type) {
			case int : 		return v + (d.var2.(int)), nil
			case string: 	return v + (d.var2.(string)), nil
			}
	}
	return nil, errors.New("Data are of different types")
}

func main() {

	integerAdd := data{var1: 98, var2: 42}
	stringAdd := data{var1: "susheel", var2: "kumar"}
	invalidAdd := data{var1: 1, var2: "susheel"}

	var addIntData Addition
	var addStringData Addition
	var addInvalidData Addition

	addIntData = &integerAdd 
	addStringData = &stringAdd
	addInvalidData = &invalidAdd

	result, err := addIntData.add()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Result is", result)
	
	result, err = addStringData.add()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Result is", result)
	
	result, err = addInvalidData.add()
	if err != nil {
		fmt.Println("Error!!!! ", err)
	}
}
