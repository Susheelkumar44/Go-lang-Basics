package main
import "fmt"
type Salutation struct{
	name string
	greet string
}
func Greet(salutation Salutation){
	message,alternative := CreateMessage(salutation.name,salutation.greet,"yo!")
	fmt.Println(message)
	fmt.Println(alternative)
	//_,alternative := CreateMessage(salutation.name,salutation.greet,"yo!")
}
func CreateMessage(name string,greet ...string) (message string,alternate string){
	//message=greet[0]+" "+name
	message=greet[1]+" "+name
	//message=greet[2]+" "+name -> Index out of range
	alternate="HEY! "+ name
	return
}
func main(){
	var s=Salutation{"Susheel","Hello"}
	Greet(s)
}