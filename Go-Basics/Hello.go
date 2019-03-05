package main
import "fmt" //used for printing and scanning to standard output and from standard input
func main() {
	i:=1
	fmt.Println("Start")
	for i=1;i<=5;i++ {
		defer fmt.Println(i)
	}
	fmt.Println("finished")
}
