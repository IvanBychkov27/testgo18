package main

import (
	"fmt"
	"math"
)

// нейросеть из одного нейрона
// x - входные данных
// w - входные веса нейрона (нейросеть)
func main() {
	//x, res := []float64{1, 0, 0}, 0.0 // res = 0
	x, res := []float64{1, 1, 0}, 0.1 // res = 1
	//x, res := []float64{1, 0, 1}, 0.2 // res = 2
	//x, res := []float64{1, 1, 1}, 0.3 // res = 3

	//w := []float64{-8.4, 6.203, 7.014} // входные веса нейрона
	w := []float64{-3.3798466596568337, 0.8422457924906098, 1.8060649570927565} // входные веса нейрона

	y, err := unitSum(x, w)
	if err != nil {
		fmt.Println("error unit sum:", err.Error())
		return
	}
	fmt.Printf("y  : %0.4f\n", y)
	mse := MSE(y, res)

	result := math.Round(y * 10)
	//if result > 3 {
	//	result = 3
	//}
	fmt.Printf("RES: %1.3f\n", result)
	fmt.Printf("MSE: %0.8f\n", mse)
}

func unitSum(x, w []float64) (float64, error) {
	if len(x) == 0 || len(x) != len(w) {
		return 0, fmt.Errorf("incorrect data entry")
	}
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += x[i] * w[i]
	}
	return sigmoid(sum), nil
}

// функция сигмойда
func sigmoid(x float64) float64 {
	return (1 / (1 + math.Exp(-x)))
}

// среднеквадратичная ошибка
func MSE(x, r float64) float64 {
	return math.Pow((r - x), 2.0)
}
