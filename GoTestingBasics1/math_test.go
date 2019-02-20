package GoTestingBasics1
import(
	"testing"
	//"time"
	"os"
)
//Creating a custom test runner
func TestMain(m* testing.M){
	println("Tests are about to run")
	result:=m.Run()
	println("Tests executing done")
	os.Exit(result)
}
//Testcase 1
func TestCanAddNumber(t *testing.T){
	/*if testing.Short(){   //(This Implementation skips long Testcases)
		t.Skip("Skipping long tests!!")
	}*/
	/*t.Parallel()
	time.Sleep(1*time.Second)*/ //This will introduce delay in the test
	result := Add(1,2)
	
	if result!=3 {
		t.Log("Failed to add")
		t.Fail()
		/*t.Log("Abandoned")
		t.FailNow()    
			//OR
		t.Fatal("Abandoned") // this will abandon the tests that are making no sense*/

	}

	result = Add(1,2,3,4)
	if result!=10 {
		t.Error("failed to add more numbers or result may not be correct")
	}
}

//Testcase 2
func TestCanSubtractNumber(t *testing.T){
	/*t.Parallel()
	time.Sleep(1*time.Second)*/ 
	result := Subtract(10, 5)
	if result != 5 {
		t.Error("failed to subtract two numbers")
	}	
}

//Testcase 3(Skipping the test)
/*func TestCanMultiplyNumber(t *testing.T){
	t.Skip("Coming shortly !!!")
}*/
