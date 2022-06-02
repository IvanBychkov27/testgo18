package main

import (
	"encoding/json"
	"fmt"
)

type ConsulData struct {
	SSDisable int64
	QPSMax    int64
	Throttle  int64
	Disable   int64
}

type Unit struct {
	W     []float64 `json:"w"`
	Value float64   `json:"value"`
}

type Net struct {
	NN  [][]*Unit `json:"nn"`  // данные нейросети
	MSE float64   `json:"mse"` // среднеквадратичная ошибка нейросети
}

func main() {

	array()

}

func array() {
	nn := make([][]*Unit, 0)
	input := &Unit{[]float64{0.1, 0.2, 0.3}, 0.1}
	hidden := &Unit{[]float64{0.4, 0.5}, 0.2}
	output := &Unit{[]float64{0.6}, 0.3}

	d := []*Unit{input, hidden, output}
	nn = append(nn, d)

	n := Net{
		NN:  nn,
		MSE: 0.001,
	}

	data, err := json.Marshal(n)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	fmt.Println("data:", string(data))

}

func networks() {
	d := []byte(`
    	{
  "10002":1111,
  "10003":1111,
  "10004":11000,
  "10005":120000,
  "10006":120000
}		`)

	res := map[int]int64{}
	err := json.Unmarshal(d, &res)
	if err != nil {
		fmt.Println("error json marshal:", err.Error())
	}
	fmt.Println("res:", res)

	networks := map[int]*ConsulData{}
	for networkID, qps := range res {
		d, ok := networks[networkID]
		if !ok {
			d = &ConsulData{}
			networks[networkID] = d
		}
		d.QPSMax = qps
	}

	fmt.Println("networks:")
	for id, d := range networks {
		fmt.Printf("%d:%d\n", id, *d)
	}
}
