package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 10
	a := 17
	b := 9
	d := 10
	for x := 0; x < n; x++ {
		r := (a*x + b) % d
		fmt.Println(r)
	}

}

func process(a []int) ([]int, []int) {
	n := len(a)
	//idx := n / 2
	idx := rand.Intn(n)
	d := a[idx]
	i := 0
	j := n - 1
	count := 0
	replay := 0

	for i < j {
		count++
		if count > 1000 || replay > n {
			fmt.Println("error: count=10000: replay =", replay)
			break
		}

		if i == idx {
			if a[j] < d {
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

		if a[i] == d {
			replay++
		}

		if a[j] <= d {
			a[i], a[j] = a[j], a[i]
			if j == idx {
				idx = i
			}
		}

		j--

	}
	fmt.Println("count:", count)
	return a[:idx], a[idx:]
}

// проверка, что все элементы > d
func checkB(d int, b []int) bool {
	for _, v := range b {
		if v < d {
			return false
		}
	}
	return true
}

//  проверка, что все элементы <= d
func checkL(d int, c []int) bool {
	for _, v := range c {
		if v > d {
			return false
		}
	}
	return true
}

func setDataRandom(n int) []int {
	rand.Seed(time.Now().UnixMicro())
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, rand.Intn(n*2)-n)
	}
	return a
}

// min
func minEl(a []int) int {
	m := 1000
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

//  max
func maxEl(a []int) int {
	m := -1000
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}
