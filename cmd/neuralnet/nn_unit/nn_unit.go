package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"
)

type Data struct {
	x []float64 // входные данных
	r []float64 // результат
}

// нейрон
type Unit struct {
	W     []float64 `json:"w"`     // входные веса нейрона
	Value float64   `json:"value"` // значение нейрона
}

type Net struct {
	NN  [][]*Unit `json:"nn"`  // данные нейросети
	MSE float64   `json:"mse"` // среднеквадратичная ошибка нейросети
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

func newNN(input, hidden, output int) *Net {
	nn := make([][]*Unit, 0)
	inputLayer := newLayer(input, 0)
	hiddenLayer := newLayer(hidden, input)
	outputLayer := newLayer(output, hidden)

	nn = append(nn, inputLayer, hiddenLayer, outputLayer)

	return &Net{nn, 1}
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
		{[]float64{1, 0, 0, 0, 0}, []float64{0.0, 0.0}},
		{[]float64{1, 1, 0, 0, 0}, []float64{0.1, 0.0}},
		{[]float64{1, 0, 1, 0, 0}, []float64{0.2, 0.0}},
		{[]float64{1, 1, 1, 0, 0}, []float64{0.3, 0.0}},
		{[]float64{1, 0, 0, 1, 0}, []float64{0.4, 0.0}},
		{[]float64{1, 1, 0, 1, 0}, []float64{0.5, 0.0}},
		{[]float64{1, 0, 1, 1, 0}, []float64{0.6, 0.0}},
		{[]float64{1, 1, 1, 1, 0}, []float64{0.7, 0.0}},
		{[]float64{1, 0, 0, 0, 1}, []float64{0.8, 0.0}},
		{[]float64{1, 1, 0, 0, 1}, []float64{0.9, 0.0}},
		{[]float64{1, 0, 1, 0, 1}, []float64{0.0, 0.1}},
		{[]float64{1, 1, 1, 0, 1}, []float64{0.1, 0.1}},
		{[]float64{1, 0, 0, 1, 1}, []float64{0.2, 0.1}},
		{[]float64{1, 1, 0, 1, 1}, []float64{0.3, 0.1}},
		{[]float64{1, 0, 1, 1, 1}, []float64{0.4, 0.1}},
		{[]float64{1, 1, 1, 1, 1}, []float64{0.5, 0.1}},
	}
	return data
}

func main() {
	rand.Seed(time.Now().UnixNano())
	data := getData()
	nn := newNN(5, 5, 1)

	trainingNN(nn, data)
	calculateAllErrRateNN(nn, data, true)
	//saveNN(nn, "cmd/neuralnet/nn_unit/file_nn.nn")

	//nn := openNN("cmd/neuralnet/nn_unit/file_nn.nn")
	//
	//idx := 9
	//d := data[idx].x
	//r := data[idx].r
	//res := directWorkNN(nn, d)
	//fmt.Printf("exp = %2.4f  res = %2.4f  MSE:%0.4f\n", r, res, MSE(res, r))
	//printNN(nn)
}

func trainingNN(net *Net, ds []Data) {
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
				//dw := errNN * thPrime(u)           // дельта весов для слоя outputLayer
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

			//fmt.Println("MSE:", errRate, era)
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
	net.MSE = mseMin
	fmt.Println("era :", era)
	//fmt.Println("rate:", rate)
}

// метод обратного распределение ошибки - корректировка весов нейросети
func reverseDistribution(net *Net, dw, rate float64) {
	const (
		inputLayer  = 0
		hiddenLayer = 1
		outputLayer = 2
	)
	var nn *Unit

	// корректировка весов слоя outputLayer
	for l := 0; l < len(net.NN[outputLayer]); l++ { // перебор всех эл-тов слоя outputLayer
		nn = net.NN[outputLayer][l]
		for i := 0; i < len(net.NN[hiddenLayer]); i++ {
			nn.W[i] -= net.NN[hiddenLayer][i].Value * dw * rate // корректировка весов слоя outputLayer
		}
	}

	// корректировка весов слоя hiddenLayer
	for l := 0; l < len(net.NN[hiddenLayer]); l++ { // перебор всех эл-тов слоя hiddenLayer
		nn = net.NN[hiddenLayer][l]

		var errHL float64                               // средняее отклонение от ожидаемого результата для слоя hiddenLayer
		for t := 0; t < len(net.NN[outputLayer]); t++ { // перебор всех эл-тов слоя outputLayer
			errHL += net.NN[outputLayer][t].W[l] * dw //
		}
		errHL = errHL / float64(len(net.NN[outputLayer]))

		dwHL := errHL * sigmoidDerivative(nn.Value) // дельта весов для слоя hiddenLayer
		//dwHL := errHL * thPrime(nn.Value) // дельта весов для слоя hiddenLayer
		for i := 0; i < len(net.NN[inputLayer]); i++ {
			nn.W[i] -= net.NN[inputLayer][i].Value * dwHL * rate // корректировка весов слоя hiddenLayer
		}
	}
}

// расчет общей среднеквадратичной погрешности нейросети
// print = true - выводит погрешности всех входных данных на экран
func calculateAllErrRateNN(net *Net, ds []Data, print bool) float64 {
	var errRateNN float64 // погрешность нейросети

	if print {
		fmt.Println("exp      res      MSE")
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
		fmt.Printf("total MSE: %0.6f\n", errRateNN/float64(len(ds)))
	}

	return errRateNN / float64(len(ds))
}

// прямая работа нейросети:
// data - входные данные,
// NN - нейросеть
func directWorkNN(net *Net, data []float64) []float64 {
	var dataLayer []float64
	submitEntryNN(net, data)           // подаем данные на входной слой нейросети
	for i := 1; i < len(net.NN); i++ { // начинаем со скрытого слоя
		dataLayer = make([]float64, 0, len(net.NN[i]))
		for _, d := range net.NN[i] {
			v := unitCalc(data, d.W)
			dataLayer = append(dataLayer, v)
			d.Value = v
		}
		data = dataLayer
	}
	return dataLayer
}

// расчет значения нейрона  (d - входные данные, W - веса нейрона)
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
	//return th(sum)
}

// подать данные на вход нейросети
func submitEntryNN(net *Net, data []float64) {
	const inputLayer = 0
	for i, d := range data {
		if i > len(net.NN[inputLayer])-1 {
			break
		}
		net.NN[inputLayer][i].Value = d
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

// th - гиперболический тангенс
func th(x float64) float64 {
	return (math.Exp(2*x) - 1) / (math.Exp(2*x) + 1)
}

// thPrime - производная гиперболического тангенса
func thPrime(x float64) float64 {
	ch := (math.Exp(x) + math.Exp(-x)) / 2 // гиперболический косинус
	return 1.0 / ch * ch
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
	return math.Sqrt(sum / float64(len(actuals)))
}

// распечатать данные нейросети
func printNN(d *Net) {
	s := []string{"Input", "Hidden", "Output"}
	for i, ls := range d.NN {
		fmt.Printf("%s layer:\n", s[i])
		for _, l := range ls {
			fmt.Println(l)
		}
		fmt.Println()
	}
}

func saveNN(net *Net, fileNameNN string) {
	jsonDataNN, errMarshal := json.Marshal(net)
	if errMarshal != nil {
		fmt.Println("error json marshal:", errMarshal.Error())
		return
	}

	file, errCreate := os.Create(fileNameNN)
	if errCreate != nil {
		fmt.Println("error create file:", errCreate.Error())
		return
	}
	defer file.Close()

	_, errWrite := file.Write(jsonDataNN)
	if errWrite != nil {
		fmt.Println("error write file:", errWrite.Error())
		return
	}

	fmt.Println("neural network in a saved file:", fileNameNN)
}

func openNN(fileNameNN string) *Net {
	nn := &Net{}
	data, err := ioutil.ReadFile(fileNameNN)
	if err != nil {
		fmt.Println("error open file", err.Error())
		return nil
	}
	err = json.Unmarshal(data, nn)
	if err != nil {
		fmt.Println("error json unmarshal", err.Error())
		return nil
	}
	return nn
}
