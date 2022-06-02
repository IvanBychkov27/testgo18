package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Data struct {
	x []float64 // входные данных
	r []float64 // результат
}

// нейрон
type Unit struct {
	w     []float64 // входные веса нейрона
	value float64   // значение нейрона
}

type NN struct {
	nn  [][]*Unit // данные нейросети
	mse float64   // среднеквадратичная ошибка нейросети
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

	return &NN{nn, 1}
}

func getData0() []Data {
	data := []Data{
		{[]float64{1, 0, 0}, []float64{0.0}},
		{[]float64{1, 1, 0}, []float64{0.3}},
		{[]float64{1, 0, 1}, []float64{0.6}},
		{[]float64{1, 1, 1}, []float64{1.0}},
	}
	return data
}

func getData() []Data {
	data := []Data{
		{[]float64{1, 0, 0, 0, 0}, []float64{0.0}},
		{[]float64{1, 1, 0, 0, 0}, []float64{0.1}},
		{[]float64{1, 0, 1, 0, 0}, []float64{0.2}},
		{[]float64{1, 1, 1, 0, 0}, []float64{0.3}},
		{[]float64{1, 0, 0, 1, 0}, []float64{0.4}},
		{[]float64{1, 1, 0, 1, 0}, []float64{0.5}},
		{[]float64{1, 0, 1, 1, 0}, []float64{0.6}},
		{[]float64{1, 1, 1, 1, 0}, []float64{0.7}},
		{[]float64{1, 0, 0, 0, 1}, []float64{0.8}},
		{[]float64{1, 1, 0, 0, 1}, []float64{0.9}},
		{[]float64{1, 0, 1, 0, 1}, []float64{1.0}},
	}
	return data
}

func getData2() []Data {
	data := []Data{
		{[]float64{1, 0, 0, 0, 0}, []float64{0.00}},
		{[]float64{1, 1, 0, 0, 0}, []float64{0.08}},
		{[]float64{1, 0, 1, 0, 0}, []float64{0.16}},
		{[]float64{1, 1, 1, 0, 0}, []float64{0.24}},
		{[]float64{1, 0, 0, 1, 0}, []float64{0.32}},
		{[]float64{1, 1, 0, 1, 0}, []float64{0.40}},
		{[]float64{1, 0, 1, 1, 0}, []float64{0.48}},
		{[]float64{1, 1, 1, 1, 0}, []float64{0.56}},
		{[]float64{1, 0, 0, 0, 1}, []float64{0.64}},
		{[]float64{1, 1, 0, 0, 1}, []float64{0.72}},
		{[]float64{1, 0, 1, 0, 1}, []float64{0.80}},
		{[]float64{1, 1, 1, 0, 1}, []float64{0.88}},
		{[]float64{1, 0, 0, 1, 1}, []float64{0.96}},
		{[]float64{1, 1, 0, 1, 1}, []float64{1.0}},
	}
	return data
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nn := newNN(5, 4, 1)
	data := getData()

	trainingNN(nn, data)
	calculateAllErrRateNN(nn, data, true)

	//idx := 1
	//d := data[idx].x
	//r := data[idx].r
	//res := directWorkNN(nn, d)
	//fmt.Printf("exp = %2.4f  res = %2.4f  mse:%0.4f\n", r, res, MSE(res, r))
	//printNN(nn)
}

func trainingNN(net *NN, ds []Data) {
	rate := 0.5
	mseMin := 1.0
	era := 0
	for {
		era++
		for _, d := range ds { // перебор всех входных данных
			submitEntryNN(net, d.x)      // подаем данные на входной слой нейросети
			us := directWorkNN(net, d.x) // текущий результат нейросети
			for i, u := range us {       // перебор всех резултатов нейросети
				errNN := u - d.r[i]                // отклонение от ожидаемого результата
				dw := errNN * sigmoidDerivative(u) // дельта весов для слоя outputLayer
				reverseDistribution(net, dw, rate) // корректировка весов нейросети
			}
		}

		if era%5 == 0 {
			errRate := calculateAllErrRateNN(net, ds, false)
			if mseMin > errRate {
				mseMin = errRate
			} else {
				break
			}

			//fmt.Println("mse:", errRate, era)
			if era > 1000000 {
				break
			}

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
	net.mse = mseMin
	fmt.Println("era :", era)
	//fmt.Println("rate:", rate)
}

// метод обратного распределение ошибки - корректировка весов нейросети
func reverseDistribution(net *NN, dw, rate float64) {
	const (
		inputLayer  = 0
		hiddenLayer = 1
		outputLayer = 2
	)
	var nn *Unit

	// корректировка весов слоя outputLayer
	for l := 0; l < len(net.nn[outputLayer]); l++ { // перебор всех эл-тов слоя outputLayer
		nn = net.nn[outputLayer][l]
		for i := 0; i < len(net.nn[hiddenLayer]); i++ {
			nn.w[i] -= net.nn[hiddenLayer][i].value * dw * rate // корректировка весов слоя outputLayer
		}
	}

	// корректировка весов слоя hiddenLayer
	for l := 0; l < len(net.nn[hiddenLayer]); l++ { // перебор всех эл-тов слоя hiddenLayer
		nn = net.nn[hiddenLayer][l]

		var errHL float64                               // средняее отклонение от ожидаемого результата для слоя hiddenLayer
		for t := 0; t < len(net.nn[outputLayer]); t++ { // перебор всех эл-тов слоя outputLayer
			errHL += net.nn[outputLayer][t].w[l] * dw //
		}
		errHL = errHL / float64(len(net.nn[outputLayer]))

		dwHL := errHL * sigmoidDerivative(nn.value) // дельта весов для слоя hiddenLayer
		for i := 0; i < len(net.nn[inputLayer]); i++ {
			nn.w[i] -= net.nn[inputLayer][i].value * dwHL * rate // корректировка весов слоя hiddenLayer
		}
	}
}

// расчет общей среднеквадратичной погрешности нейросети
// print = true - выводит погрешности всех входных данных на экран
func calculateAllErrRateNN(net *NN, ds []Data, print bool) float64 {
	var errRateNN float64 // погрешность нейросети

	if print {
		fmt.Println("exp      res      mse")
	}

	for _, d := range ds {
		y := directWorkNN(net, d.x)
		mse := MSE(y, d.r)
		errRateNN += mse
		if print {
			fmt.Printf("%0.2f   %0.2f   %0.6f\n", d.r, y, mse)
		}
	}

	if print {
		fmt.Printf("total mse: %0.6f\n", errRateNN/float64(len(ds)))
	}

	return errRateNN / float64(len(ds))
}

// прямая работа нейросети:
// data - входные данные,
// nn - нейросеть
func directWorkNN(net *NN, data []float64) []float64 {
	var dataLayer []float64
	submitEntryNN(net, data)           // подаем данные на входной слой нейросети
	for i := 1; i < len(net.nn); i++ { // начинаем со скрытого слоя
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
func MSE(actuals, expected []float64) float64 {
	if len(actuals) != len(expected) {
		return -1
	}
	var sum float64
	for i, actual := range actuals {
		sum += math.Pow((expected[i] - actual), 2.0)
	}
	return sum / float64(len(actuals))
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
