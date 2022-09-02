package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(random(30))
	}
}

// генератор квазислучайных чисел
func random(n int64) int64 {
	var a, b int64
	a = 7
	b = 3
	x := (-1) * time.Now().UnixNano()
	return (a*x + b) % n
}
