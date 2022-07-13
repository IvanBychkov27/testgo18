// https://vavilen84.com/ru/algorithms/03/algorithms_in_go/

/*
	Бинарный поиск

	Сложность алгоритма: худшее время - логарифмическое O(log n); лучшее - константное O(1)
	Алгоритм будет работать только с отсортированными массивами.
	Идея заключается:
    - разделить набор данных на две части
    - сравнить поиск с первым (последним) элементом блока данных
    - выбрать правильный блок
    - разделить выбранный блок на две части
    - и т. д. до тех пор, пока мы не получим правильное значение
*/

package main

import (
	"fmt"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 22, 33, 40, 50, 60}
	res, idx, iterationsCount := searchInIntSlice(data, 40)
	fmt.Printf("res: %t, idx: %d, iterationsCount: %d \n", res, idx, iterationsCount)
}

// binary_search
func searchInIntSlice(data []int, n int) (res bool, idx, iterationsCount int) {
	//sort.Ints(data)
	lowKey := 0              // первый индекс
	highKey := len(data) - 1 // последний индекс
	if (data[lowKey] > n) || (data[highKey] < n) {
		return // нужное значение не в диапазоне данных
	}
	for lowKey <= highKey {
		// уменьшаем список рекурсивно
		iterationsCount++
		idx = (lowKey + highKey) / 2 // середина
		if data[idx] == n {
			res = true // мы нашли значение
			return
		}
		if data[idx] < n {
			// если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
			lowKey = idx + 1
			continue
		}
		if data[idx] > n {
			// если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
			highKey = idx - 1
		}
	}
	return
}
