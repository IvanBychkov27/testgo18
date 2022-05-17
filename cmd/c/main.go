// http://прохоренок.рф/pdf/go/ch15-go-funktsii-dlya-raboty-s-katalogami.html
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getwd()) // Getwd() — позволяет получить строковое представление текущего рабочего каталога
}
