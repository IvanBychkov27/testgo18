/*
	Быстрая сортировка

	Сложность алгоритма: худшее время - квадратичное O(n2); лучшее - линейно-логарифмическое время O(n log n)
	Идея заключается:
    - взять первое значение как опорную точку
    - перебирать остальные элементы
    - разделить элементы на те, которые больше и те, которые меньше опорной точки
    - рекурсивно отсортировать меньшие и большие элементы
*/

package main

import (
	"fmt"
	"sync/atomic"
)

var count int64

func main() {
	data := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	res := quickSort(data)
	fmt.Println("res:", res)
	fmt.Println("count:", count)
}

func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	less := make([]int, 0)
	bigger := make([]int, 0)
	n := data[0]
	for _, v := range data[1:] {
		atomic.AddInt64(&count, 1)
		if v > n {
			bigger = append(bigger, v)
		} else {
			less = append(less, v)
		}
	}
	data = append(quickSort(less), n)
	data = append(data, quickSort(bigger)...)
	return data
}
