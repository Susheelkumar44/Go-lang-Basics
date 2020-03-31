package main

import "fmt"

//Function anything accepts parameter of any data type
func Anything(anything interface{}) {
	fmt.Println(anything)
}

type testEmptyInterface struct {
	name string
	age  int
}

func main() {
	Anything(2.44)
	Anything("Any Data type")
	Anything(struct{}{}) //Printing empty structure

	test := testEmptyInterface{
		name: "Name",
		age:  24,
	}
	Anything(test) //Printing structure data

	//This means Key is of type string and value is of any type
	mymap := make(map[string]interface{})
	mymap["name"] = "Name1"
	mymap["age"] = 26

	Anything(mymap)
}
