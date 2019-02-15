package main
import "fmt"
func main() {
	i:=1
	fmt.Println("Start")
	for i=1;i<=5;i++ {
		defer fmt.Println(i)
	}
	fmt.Println("finished")
}
