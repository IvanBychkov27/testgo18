package main

import (
	"fmt"
	"math"
)

func main() {
	//a := 19354838709
	a := 50000000000
	d := int(49e5)
	b := 1e7

	res := int(math.Round(float64(a+d)/b) * b)

	fmt.Println("res =", res)
}
