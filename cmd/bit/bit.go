package main

import (
	"fmt"
	"github.com/yourbasic/bit"
)

func main() {
	//testBit()

	a := make([]int, 0, 2)
	a = append(a, 10)
	a = append(a, 20)
	fmt.Println("a =", a)
	a = nil
	fmt.Println("a =", a)
	a = append(a, 777)
	a = append(a, 999)
	fmt.Println("a =", a)

}

func testBit() {
	var r bit.Set
	r.Add(25)
	r.Add(5)
	r.Add(1)
	r.Add(0)
	r.Add(5)

	fmt.Println(r.String())

	fmt.Println(r.Contains(5))

	fmt.Println(r.Max())
	fmt.Println(r.Size())

	res := make([]int, 0, r.Size())
	for d := 0; d != -1; d = r.Next(d) {
		res = append(res, d)
	}
	fmt.Println(res)

}

//func bitSort(a[]int) []int {
//
//}
