package main

import "fmt"

func main() {

	m := make(map[string]string)

	key := "a"
	m[key] = "text"

	m = nil

	_, ok := m[key]
	if !ok {
		fmt.Printf("key [%s] no found\n", key)
	} else {
		fmt.Printf("key [%s] ok\n", key)
	}

}
