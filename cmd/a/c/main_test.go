package main

import "testing"

func BenchmarkQuote(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strQuote(`abc"def\h`)
	}
}

func BenchmarkAppendQuote(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strAppendQuote([]byte{}, `abc"def\h`)
	}
}
