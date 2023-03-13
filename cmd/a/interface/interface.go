// https://habr.com/ru/company/vk/blog/463063/
package main

import (
	"fmt"
	"log"
	"strconv"
)

// Объявляем тип Book, который удовлетворяет интерфейсу fmt.Stringer.
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Объявляем тип Count, который удовлетворяет интерфейсу fmt.Stringer.
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Объявляем функцию WriteLog(), которая берёт любой объект,
// удовлетворяющий интерфейсу fmt.Stringer в виде параметра.
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	// Инициализируем объект Book и передаём в WriteLog().
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)

	// Инициализируем объект Count и передаём в WriteLog().
	count := Count(3)
	WriteLog(count)
}

/*
Главное, что нужно запомнить:
используя интерфейсный тип в объявлении функции WriteLog(),
мы сделали функцию безразличной (или гибкой) к типу принимаемого объекта.
Важно лишь то, какие у него методы.
*/
