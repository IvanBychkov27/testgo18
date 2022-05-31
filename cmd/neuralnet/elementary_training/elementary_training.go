package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Data struct {
	x []float64 // входные данных
	r float64   // результат
}

type NN struct {
	w   []float64 // входные веса нейрона
	mse float64   // среднеквадратичная ошибка нейросети
}

func main() {
	nn := newNN(3)
	ds := getData()

	trainingNN(nn, ds, 0.9, 100)

	fmt.Println("nn :", nn.w)
	fmt.Println("mse:", nn.mse)

	calculateAllErrRateNN(nn, ds, true)
}

// обучение нейросети:
// nn - обучаемая нейросеть;
// ds - данные для обучения нейросети;
// rate - скорость обучения нейросети (от 10 до 0.001 - чем меньше значение тем медленее обучение, но более точное);
// era - кол-во проходов по всем обучаемым данным;
// rate и era - подбираются экспериментально
func trainingNN(nn *NN, ds []Data, rate float64, era int) {
	if len(ds) == 0 || len(nn.w) == 0 || len(nn.w) != len(ds[0].x) {
		fmt.Printf("error training nn - incorrect data entry: len data %d, len nn %d", len(ds[0].x), len(nn.w))
		return
	}

	for t := 0; t < era; t++ { // количество обучений
		for _, d := range ds { // перебор всех входных данных
			u := unitSum(d.x, nn.w)
			errSum := u - d.r
			dw := errSum * sigmoidDerivative(u)
			for i := 0; i < len(d.x); i++ {
				nn.w[i] -= d.x[i] * dw * rate
			}
		}
	}
	nn.mse = calculateAllErrRateNN(nn, ds, false)
}

func getData() []Data {
	data := []Data{
		{[]float64{1, 0, 0}, 0.0},
		{[]float64{1, 1, 0}, 0.1},
		{[]float64{1, 0, 1}, 0.2},
		{[]float64{1, 1, 1}, 0.3},
	}
	return data
}

// загружает нейросеть
func getNN() *NN {
	return &NN{[]float64{-8.4, 6.203, 7.014}, 0}
}

// создает новую необученную нейросеть с n входами
func newNN(n int) *NN {
	rand.Seed(time.Now().UnixNano())
	w := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		w = append(w, rand.Float64())
	}
	return &NN{w, 1}
}

// расчет общей среднеквадратичной погрешности нейросети
func calculateAllErrRateNN(nn *NN, ds []Data, print bool) float64 {
	var errRateNN float64 // погрешность нейросети
	if print {
		fmt.Println("exp    res    mse")
	}

	for _, d := range ds {
		y := unitSum(d.x, nn.w)
		mse := MSE(y, d.r)
		errRateNN += mse
		if print {
			fmt.Printf("%0.2f   %0.2f   %0.4f\n", d.r, y, mse)
		}
	}
	return errRateNN / float64(len(ds))
}

// прямая работа нейросети (x - входные данные, w - веса нейросети)
func unitSum(x, w []float64) float64 {
	if len(x) == 0 || len(x) != len(w) {
		fmt.Printf("incorrect data entry: len data %d, len nn %d", len(x), len(w))
		return 0
	}
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += x[i] * w[i]
	}
	return sigmoid(sum)
}

// функция сигмойда
func sigmoid(x float64) float64 {
	return (1 / (1 + math.Exp(-x)))
}

// производная функции сигмойда
func sigmoidDerivative(x float64) float64 {
	return sigmoid(x) * (1 - sigmoid(x))
}

// среднеквадратичная ошибка
func MSE(x, r float64) float64 {
	return math.Pow((r - x), 2.0)
}
