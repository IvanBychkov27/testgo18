// http://прохоренок.рф/pdf/go/ch15-go-funktsii-dlya-raboty-s-katalogami.html
package main

import "fmt"

type data struct {
	d int
}

func main() {
	a := map[int]*data{1: &data{10}}

	idx := 1
	_, ok := a[idx]
	if ok && a[idx].d == 10 {
		fmt.Println("len(a):", len(a))
		fmt.Println("data:", a[idx].d)
	}

}
