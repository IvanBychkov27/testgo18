package main

import (
	"fmt"
	"time"
)

func main() {
	var res1, res2 int
	var res3 string

	chInt1 := make(chan int)
	chInt2 := make(chan int)
	chStr := make(chan string)

	go go01(chInt1)
	go go02(chInt2)
	go go03(chStr)

	res1 = <-chInt1
	res2 = <-chInt2
	res3 = <-chStr

	fmt.Println("res1:", res1)
	fmt.Println("res2:", res2)
	fmt.Println("res3:", res3)
}

func go01(ch chan int) {
	time.Sleep(time.Second * 2)
	ch <- 1
	return
}

func go02(ch chan int) {
	time.Sleep(time.Second * 1)
	ch <- 2
	return
}
func go03(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "3"
	return
}
