package main

import (
	"fmt"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

	page := 4
	perPage := 5
	total := len(data)

	n, k := currentData(page, perPage, total)

	fmt.Println("d:", data[n:k])

}

func currentData(page, perPage, total int) (int, int) {
	k := page * perPage
	n := k - perPage
	if n >= total {
		n = 0
		k = perPage
	}
	if k > total {
		k = total
	}
	return n, k
}
