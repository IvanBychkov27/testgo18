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
	"math/rand"
	"sync/atomic"
)

var count int64

func main() {
	data := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//res := quickSort(data)
	res := quickSortV4(data, false)
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

// O(N-1) !!!
func quickSortV4(a []int, random bool) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	isSort := true
	for i := 1; i < n; i++ {
		count++
		if a[i-1] > a[i] {
			isSort = false
			break
		}
	}
	if isSort {
		return a
	}

	idx := n / 2
	if random {
		idx = rand.Intn(n)
	}

	d := a[idx]
	i := 0
	j := n - 1

	for i < j {
		count++
		if i == idx {
			if a[j] <= d {
				a[i], a[j] = a[j], a[i]
				idx = j
				i++
			} else {
				j--
			}
			continue
		}

		if a[i] < d {
			i++
			continue
		}

		if a[j] <= d {
			a[i], a[j] = a[j], a[i]
			if j == idx {
				idx = i
			}
		}
		j--

	}

	random = false
	if idx == 0 || idx == n-1 {
		random = true
	}

	return append(quickSortV4(a[:idx], random), quickSortV4(a[idx:], random)...)
}
