package main
import "fmt"
func main(){
	msg := "Hello"
	//var greet *string = &msg
	greet := &msg
	fmt.Println(msg,greet)
	fmt.Println(msg,*greet)
	//fmt.Println(msg,&greet)
}