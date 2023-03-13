package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0, 100)
	a = append(a, 1, 2, 3)
	fmt.Println(a)
	//process(a)
	process1(a)
	fmt.Println(a)
	fmt.Println("----------------------------")

	var m map[string]string
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
	d, _ := m["k"]
	fmt.Println("d:", d)

	fmt.Println("ok")

}

func process(a []int) {
	for idx, _ := range a {
		a[idx] = 0
	}
}

func process1(a []int) {
	a = append(a, 4, 5, 6)
	fmt.Println(a)
}
