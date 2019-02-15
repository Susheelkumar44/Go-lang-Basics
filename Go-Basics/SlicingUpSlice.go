package main
import "fmt"
func main(){
	slice := []int {1,2,3}
	fmt.Println(slice)
	slice=append(slice[0:2])
	//slice=append(slice[:1],slice[2:]...)
	fmt.Println(slice)
}
	