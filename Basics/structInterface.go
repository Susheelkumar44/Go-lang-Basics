package main

import "fmt"

type lamborgini struct {
	name  string
	model string
	price string
}

type bmw struct {
	name  string
	model string
	price string
}

type carFunctions interface {
	Drive()
	Details()
}

func (l *lamborgini) Drive() {
	fmt.Println("Lamborgini driving")
}

func (l *lamborgini) Details() {
	fmt.Println(*l)
}

func (b *bmw) Drive() {
	fmt.Println("BMW Driving")
}

func (b *bmw) Details() {
	fmt.Println(*b)
}

func main() {
	l := lamborgini{
		name:  "Lamborgini46IClass",
		model: "2019-GC6",
		price: "Rs.800000",
	}
	b := bmw{
		name:  "BMW",
		model: "F-class",
		price: "Rs.650000",
	}

	l.Drive()
	l.Details()
	b.Details()
	b.Drive()

}
