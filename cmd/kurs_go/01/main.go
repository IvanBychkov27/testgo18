// https://www.youtube.com/watch?v=9Pk7xAT_aCU&list=PLrCZzMib1e9q-X5V9pTM6J0AemRWseM7I&index=1

package main

import "fmt"

type Data struct {
	Sum  int
	Med  int
	Min  int
	Max  int
	Mult int
}

func main() {
	ds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//ds := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("sum =", sum(ds))
	fmt.Println("med =", med(ds))
	fmt.Println("min =", min(ds))
	fmt.Println("max =", max(ds))
	fmt.Println("mult=", mult(ds))

	fmt.Println(calculation(ds))
}

func calculation(ds []int) Data {
	l := len(ds)
	if l == 0 {
		return Data{}
	}

	data := Data{
		Sum:  0,
		Med:  0,
		Min:  ds[0],
		Max:  ds[0],
		Mult: 1,
	}

	for _, d := range ds {
		data.Sum += d
		data.Mult *= d
		if d < data.Min {
			data.Min = d
		}
		if d > data.Max {
			data.Max = d
		}
	}
	data.Med = data.Sum / l

	return data
}

func sum(ds []int) int {
	res := 0
	for _, d := range ds {
		res += d
	}
	return res
}

func mult(ds []int) int {
	res := 1
	for _, d := range ds {
		res *= d
	}
	return res
}

func med(ds []int) int {
	l := len(ds)
	if l == 0 {
		return 0
	}
	res := 0
	for _, d := range ds {
		res += d
	}
	return res / l
}

func min(ds []int) int {
	l := len(ds)
	if l == 0 {
		return 0
	}
	res := ds[0]
	for _, d := range ds {
		if d < res {
			res = d
		}
	}
	return res
}

func max(ds []int) int {
	l := len(ds)
	if l == 0 {
		return 0
	}
	res := ds[0]
	for _, d := range ds {
		if d > res {
			res = d
		}
	}
	return res
}
