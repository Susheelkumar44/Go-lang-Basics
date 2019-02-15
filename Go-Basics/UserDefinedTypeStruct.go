package main
import "fmt"
type Salutation struct{
	name string
	greet string
}
func main(){
	//s := Salutation{"Susheel","Hello"}
	//var s = Salutation{greet : "Hello Hi!", name : "Susheel"}
	s := Salutation{greet : "Hello! Hi", name : "Susheel"}
	fmt.Println(s.name,s.greet)
}