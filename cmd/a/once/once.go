package main

import (
	"fmt"
	"sync"
)

/*
Do вызывает функцию f тогда и только тогда, когда Do вызывается впервые для данного экземпляра Once. Другими словами,
если Once.Do(f) вызывается несколько раз, только первый!!! вызов вызовет f, даже если f имеет другое значение при каждом вызове.
Для выполнения каждой функции требуется новый экземпляр Once.
*/
func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
