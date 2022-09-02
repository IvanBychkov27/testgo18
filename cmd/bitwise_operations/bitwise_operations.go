/* https://metanit.com/go/tutorial/2.7.php
Поразрядные операции

*/
package main

import "fmt"

func main() {
	//shiftOperations()
	//fmt.Println()
	//andBitwiseOperations()
	//fmt.Println()
	//orBitwiseOperations()
	//fmt.Println()
	//exclusiveBitwiseOperations()
	//fmt.Println()
	//andNoBitwiseOperations()

	binarySystem()
}

// &: поразрядная конъюнкция (операция И или поразрядное умножение).
// Возвращает 1, если оба из соответствующих разрядов обоих чисел равны 1.
// Возвращает 0, если разряд хотя бы одного числа равен 0
/*
	0 & 0 = 0
	0 & 1 = 0
	1 & 0 = 0
	1 & 1 = 1
*/
func andBitwiseOperations() {
	var a, b int
	a, b = 0, 0
	fmt.Printf("%d & %d = %d\n", a, b, a&b)
	a, b = 0, 1
	fmt.Printf("%d & %d = %d\n", a, b, a&b)
	a, b = 1, 0
	fmt.Printf("%d & %d = %d\n", a, b, a&b)
	a, b = 1, 1
	fmt.Printf("%d & %d = %d\n", a, b, a&b)
}

// |: поразрядная дизъюнкция (операция ИЛИ или поразрядное сложение).
// Возвращает 1, если хотя бы один из соответствующих разрядов обоих чисел равен 1
/*
	0 | 0 = 0
	0 | 1 = 1
	1 | 0 = 1
	1 | 1 = 1
*/
func orBitwiseOperations() {
	var a, b int
	a, b = 0, 0
	fmt.Printf("%d | %d = %d\n", a, b, a|b)
	a, b = 0, 1
	fmt.Printf("%d | %d = %d\n", a, b, a|b)
	a, b = 1, 0
	fmt.Printf("%d | %d = %d\n", a, b, a|b)
	a, b = 1, 1
	fmt.Printf("%d | %d = %d\n", a, b, a|b)
}

// ^: поразрядное исключающее ИЛИ.
// Возвращает 1, если только один из соответствующих разрядов обоих чисел равен 1
/*
	0 ^ 0 = 0
	0 ^ 1 = 1
	1 ^ 0 = 1
	1 ^ 1 = 0
*/
func exclusiveBitwiseOperations() {
	var a, b int
	a, b = 0, 0
	fmt.Printf("%d ^ %d = %d\n", a, b, a^b)
	a, b = 0, 1
	fmt.Printf("%d ^ %d = %d\n", a, b, a^b)
	a, b = 1, 0
	fmt.Printf("%d ^ %d = %d\n", a, b, a^b)
	a, b = 1, 1
	fmt.Printf("%d ^ %d = %d\n", a, b, a^b)
}

// &^: сброс бита (И НЕ). В выражении z = x &^ y каждый бит z равен 0, если соответствующий бит y равен 1.
// Если бит в y равен 0, то берется значение соответствующего бита из x.
/*
	0 &^ 0 = 0
	0 &^ 1 = 0
	1 &^ 0 = 1
	1 &^ 1 = 0
*/
func andNoBitwiseOperations() {
	var a, b int
	a, b = 0, 0
	fmt.Printf("%d &^ %d = %d\n", a, b, a&^b)
	a, b = 0, 1
	fmt.Printf("%d &^ %d = %d\n", a, b, a&^b)
	a, b = 1, 0
	fmt.Printf("%d &^ %d = %d\n", a, b, a&^b)
	a, b = 1, 1
	fmt.Printf("%d &^ %d = %d\n", a, b, a&^b)
}

// Операции сдвига
func shiftOperations() {
	var d int
	d = 2                               // 10 (двоичная)
	fmt.Printf("d = %d -> %b \n", d, d) // d = 2 -> 10 (двоичная)
	d = d >> 1                          // сдвиг впаво на 1 разряд
	fmt.Printf("d = %d -> %b \n", d, d) // d = 1 -> 1 (двоичная)
	d = d << 3                          // сдвиг влево на 3 разряда
	fmt.Printf("d = %d -> %b \n", d, d) // d = 8 -> 1000 (двоичная)
}

// перевод числа в двоичную систему
func binarySystem() {
	var d int
	d = 10 // 1010
	res := fmt.Sprintf("%b", d)
	fmt.Println(res)
}
