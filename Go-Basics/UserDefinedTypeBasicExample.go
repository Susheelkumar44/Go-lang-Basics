package main
import "fmt"
type Salutation string
func main(){
	var msg Salutation = "Hello World"
	//msg Salutation :="Hello World" -> will not work
	//msg := "Hello World"
	//msg string := "Hello" ->Will not work
	fmt.Println(msg)
}
