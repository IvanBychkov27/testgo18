package main

import "fmt"

func main() {
	a := make([]int, 0, 100)
	a = append(a, 1, 2, 3)
	fmt.Println(a)
	//process(a)
	//process1(a)
	a = process2(a)
	fmt.Println(a)
}

func process(a []int) {
	for idx := range a {
		a[idx] = 0
	}
}

func process1(a []int) {
	a = append(a, 4, 5, 6)
	fmt.Println(a)
}

func process2(a []int) []int {
	a = append(a, 4, 5, 6)
	return a
}
