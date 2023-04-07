package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmespath/go-jmespath"
)

func main() {
	jsondata := []byte(`{"foo": {"bar": {"baz": [0, 1, 2, 3, 4]}}}`) // your data
	expression := "foo.bar.baz[2]"

	res, err := jmespathExemple(jsondata, expression)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	fmt.Println("result:", res)
}

func jmespathExemple(jsondata []byte, expression string) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal(jsondata, &data)
	if err != nil {
		return nil, err
	}
	result, errSearch := jmespath.Search(expression, data)
	if errSearch != nil {
		return nil, errSearch
	}
	return result, nil
}
