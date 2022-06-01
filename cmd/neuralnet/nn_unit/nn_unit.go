package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Data struct {
	x []float64 // входные данных
	r float64   // результат
}

// нейрон
type Unit struct {
	w     []float64 // входные веса нейрона
	value float64   // значение нейрона
}

type NN struct {
	nn [][]*Unit
}

// создание нейрона:
// n - кол-во связей (весов) нейрона с предыдущем слоем,
// n = 0 для входного (первого слоя) нейрона,
// v - значение нейрона
func newUnit(n int) *Unit {
	w := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		w = append(w, rand.Float64())
	}
	return &Unit{w, 1}
}

// создание слоя нейронной сети:
// n - кол-во нейронов в слое,
// m - кол-во связей у нейрона (кол-во нейронов в предыдущем слое),
// m = 0 для входного слоя
func newLayer(n, m int) []*Unit {
	layer := make([]*Unit, 0, n)
	for i := 0; i < n; i++ {
		layer = append(layer, newUnit(m))
	}
	return layer
}

func newNN(input, hidden, output int) *NN {
	nn := make([][]*Unit, 0)
	inputLayer := newLayer(input, 0)
	hiddenLayer := newLayer(hidden, input)
	outputLayer := newLayer(output, hidden)

	nn = append(nn, inputLayer, hiddenLayer, outputLayer)

	return &NN{nn}
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

func main() {
	//rand.Seed(time.Now().UnixNano())
	nn := newNN(3, 2, 1)
	data := getData()

	//submitEntryNN(nn, data[0].x) // подаем данные на входной слой нейросети
	res := directWorkNN(nn, data[0].x)

	fmt.Printf("res = %2.4f\n", res)
	printNN(nn)
}

// прямая работа нейросети:
// data - входные данные,
// nn - нейросеть
func directWorkNN(net *NN, data []float64) []float64 {
	const (
		inputLayer = iota
		hiddenLayer
		outputLayer
	)
	var dataLayer []float64
	for i := hiddenLayer; i < len(net.nn); i++ {
		dataLayer = make([]float64, 0, len(net.nn[i]))
		for _, d := range net.nn[i] {
			v := unitCalc(data, d.w)
			dataLayer = append(dataLayer, v)
			d.value = v
		}
		data = dataLayer
	}
	return dataLayer
}

// расчет значения нейрона  (d - входные данные, w - веса нейрона)
func unitCalc(d, w []float64) float64 {
	if len(d) < len(w) {
		for i := len(d); i < len(w); i++ {
			d = append(d, 1)
		}
	}
	var sum float64
	for i := 0; i < len(w); i++ {
		sum += d[i] * w[i]
	}
	return sigmoid(sum)
}

// подать данные на вход нейросети
func submitEntryNN(net *NN, data []float64) {
	const inputLayer = 0
	for i, d := range data {
		if i > len(net.nn[inputLayer])-1 {
			break
		}
		net.nn[inputLayer][i].value = d
	}
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

// распечатать данные нейросети
func printNN(d *NN) {
	s := []string{"Input", "Hidden", "Output"}
	for i, ls := range d.nn {
		fmt.Printf("%s layer:\n", s[i])
		for _, l := range ls {
			fmt.Println(l)
		}
		fmt.Println()
	}
}
