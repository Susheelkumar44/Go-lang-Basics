package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* the difference holds here if you are not sending reciever type as a pointer
If we use below method instead of (v *Vertex) output will be:
Before scaling: v = {X:3 Y:4}, Abs: 5
After scaling: v = {X:3 Y:4}, Abs: 5
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
*/
func main() {
	v := Vertex{3, 4}
	fmt.Printf("Before scaling: v = %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: v = %+v, Abs: %v\n", v, v.Abs())
}
