package GoTestingBasics2

func Add(nums... int) int{
	var result int

	if len(nums)==0{
		println("No arguments passed")
	}

	for _,i:= range nums{
		result+=i
	}
	return result
}

func Subtract(initial int, nums ... int) int{

	for _, i := range nums{
		initial -= i
	}
	return initial
}

