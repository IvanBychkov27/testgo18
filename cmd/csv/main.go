package main

import (
	"bytes"
	"fmt"
	"github.com/dnlo/struct2csv"
)

type DataSCV struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	UA   string `json:"ua"`
}

func main() {
	ds := []DataSCV{
		{1, "One", "1 11", "aa,a"},
		{2, "Two", "2023-03-07 10:16:41", "bb,b"},
		{3, "Three", "3 33", "cc,c"},
	}

	var buf bytes.Buffer

	w := struct2csv.NewWriter(&buf)
	err := w.WriteColNames(ds[0])
	if err != nil {
		fmt.Printf("error write col name, %w", err)
	}
	w.ColNames()

	for _, d := range ds {
		err := w.WriteStruct(d)
		if err != nil {
			fmt.Printf("error write struct, %w", err)
		}
		w.Flush()
	}

	fmt.Println(buf.String())
}
