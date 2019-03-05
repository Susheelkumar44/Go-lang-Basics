package main
import "fmt" //used for printing and scanning
func main() {
	i:=1
	fmt.Println("Start")
	for i=1;i<=5;i++ {
		defer fmt.Println(i)
	}
	fmt.Println("finished")
}
