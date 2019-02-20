package GoTestingBasics2
import(
	"testing"
//	"time"
//	"os"
)
var addTable = [] struct{
	in [] int
	out int
}{
	{[]int{1,2},3},
	{[]int{1,2,3,4},10},
}

//Testcase 1 (Table Driven approach)
func TestCanAddNumber(t *testing.T){

	for _,entry := range addTable{
		result := Add(entry.in...)
		if result!=entry.out{
			t.Error("failed to add")
		}
	}
}
//Testcase 2 
func TestCanSubtractNumber(t *testing.T){

	result := Subtract(10, 5)
	if result != 5 {
		t.Error("failed to subtract two numbers")
	}	
}