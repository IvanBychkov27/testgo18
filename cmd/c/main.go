package main

import "fmt"

func main() {
	d := make([]byte, 0, 1024)

	fmt.Println("len d =", len(d))
	fmt.Println("cap d =", cap(d))
}
