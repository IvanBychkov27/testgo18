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

func main() {
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
