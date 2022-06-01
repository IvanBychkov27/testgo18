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
	nn := newNN(5)
	ds := getData()
	//ds := getData2()

	trainingNN(nn, ds)

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
func trainingNN(nn *NN, ds []Data) {
	if len(ds) == 0 || len(nn.w) == 0 || len(nn.w) != len(ds[0].x) {
		fmt.Printf("error training nn - incorrect data entry: len data %d, len nn %d", len(ds[0].x), len(nn.w))
		return
	}

	rate := 0.5
	mseMin := 1.0
	era := 0
	for { // количество обучений
		era++
		for _, d := range ds { // перебор всех входных данных
			u := unitSum(d.x, nn.w)             // текущий результат нейросети
			errSum := u - d.r                   // отклонение от ожидаемого результата
			dw := errSum * sigmoidDerivative(u) // дельта весов
			for i := 0; i < len(d.x); i++ {
				nn.w[i] -= d.x[i] * dw * rate // корректировка весов нейросети
			}
		}

		if era%5 == 0 {
			dmse := calculateAllErrRateNN(nn, ds, false)
			if mseMin > dmse {
				mseMin = dmse
			} else {
				break
			}

			//fmt.Println("mse:", dmse, t)
			if rate > 0.01 {
				rate -= 0.01
			} else {
				rate -= 0.001
			}

			if rate <= 0 {
				rate = 0.0001
			}
		}
	}
	nn.mse = calculateAllErrRateNN(nn, ds, false) // сохранение данных среднеквадратичной ошибки для нейросети
	fmt.Println("era :", era)
	//fmt.Println("rate:", rate)
}

func getData() []Data {
	data := []Data{
		{[]float64{1, 0, 0, 0, 0}, 0.0},
		{[]float64{1, 1, 0, 0, 0}, 0.1},
		{[]float64{1, 0, 1, 0, 0}, 0.2},
		{[]float64{1, 1, 1, 0, 0}, 0.3},
		{[]float64{1, 0, 0, 1, 0}, 0.4},
		{[]float64{1, 1, 0, 1, 0}, 0.5},
		{[]float64{1, 0, 1, 1, 0}, 0.6},
		{[]float64{1, 1, 1, 1, 0}, 0.7},
		{[]float64{1, 0, 0, 0, 1}, 0.8},
		{[]float64{1, 1, 0, 0, 1}, 0.9},
		{[]float64{1, 0, 1, 0, 1}, 1.0},
	}
	return data
}

func getData2() []Data {
	data := []Data{
		{[]float64{1, 0, 0, 0, 0}, 0.00},
		{[]float64{1, 1, 0, 0, 0}, 0.08},
		{[]float64{1, 0, 1, 0, 0}, 0.16},
		{[]float64{1, 1, 1, 0, 0}, 0.24},
		{[]float64{1, 0, 0, 1, 0}, 0.32},
		{[]float64{1, 1, 0, 1, 0}, 0.40},
		{[]float64{1, 0, 1, 1, 0}, 0.48},
		{[]float64{1, 1, 1, 1, 0}, 0.56},
		{[]float64{1, 0, 0, 0, 1}, 0.64},
		{[]float64{1, 1, 0, 0, 1}, 0.72},
		{[]float64{1, 0, 1, 0, 1}, 0.80},
		{[]float64{1, 1, 1, 0, 1}, 0.88},
		{[]float64{1, 0, 0, 1, 1}, 0.96},
		{[]float64{1, 1, 0, 1, 1}, 1.0},
	}
	return data
}

// загружает нейросеть
func getNN() *NN {
	return &NN{[]float64{-2.796601189839126, 0.5634434994953129, 1.338581057367806}, 0.0008850383718151193}
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
