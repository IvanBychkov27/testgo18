package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	dataMX sync.RWMutex
	data   []int
	n      int
	m      map[int]int
	ch     chan int
}

func New(n int) *Queue {
	return &Queue{
		data: []int{},
		n:    n,
	}
}

func (q *Queue) Add(d int) error {
	q.dataMX.Lock()
	defer q.dataMX.Unlock()

	if len(q.data) >= q.n {
		return fmt.Errorf("exceeding the limit, n = %d", q.n)
	}

	q.data = append(q.data, d)

	return nil
}

func (q *Queue) Read() int {
	q.dataMX.RLock()
	defer q.dataMX.RUnlock()

	if len(q.data) == 0 {
		return 0
	}

	d := q.data[0]

	q.Delete()

	return d
}

func (q *Queue) Delete() {
	if len(q.data) == 0 {
		return
	}
	q.data = q.data[1:]
}

func main() {
	q := New(3)
	q.m = make(map[int]int)
	q.ch = make(chan int)

	//for i := 1; i < 5; i++ {
	//	err := q.Add(i)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//}
	//
	//time.Sleep(time.Second * 1)
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Println(q.Read())
	//}
	//
	//if q.m != nil {
	//	q.m[1] = 10
	//}
	//
	//d, ok := q.m[1]
	//fmt.Println(d, ok)

	if q.ch == nil {
		fmt.Println("channel close")
	}

	fmt.Println("m count", len(q.m))

	fmt.Println("Done...")
}
