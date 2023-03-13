package main

import "fmt"

func main() {

	a := make(map[int]int, 10)
	a[1] = 10
	a[2] = 20
	fmt.Println(a, len(a))

	fmt.Printf("%%")
}
