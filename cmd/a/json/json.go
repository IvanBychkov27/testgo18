package main

import (
	"bytes"
	"fmt"
)

// бодавить экранирование кавычек в поле source
func main() {
	a := []byte(`{"source_id":"`)
	source_id := []byte(`12"3""4`)
	//source_id := []byte(`1234`)
	b := []byte(`"}`)

	source_id = escapeQuotationMarks(source_id)

	res := make([]byte, 0, 192)
	res = append(res, a...)
	res = append(res, source_id...)
	res = append(res, b...)

	fmt.Println("res:", string(res))
}

// экранирование кавычек
func escapeQuotationMarks(data []byte) []byte {
	if !bytes.Contains(data, []byte(`"`)) {
		return data
	}

	escapeQuotationMark := []byte(`\"`)
	res := make([]byte, 0, 2*len(data))
	for _, d := range data {
		if d == '"' {
			res = append(res, escapeQuotationMark...)
			continue
		}
		res = append(res, d)
	}
	return res
}
