package main

import "testing"

func Benchmark_bytesKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesKey([]byte("asdfghjklzxcvbnmqwertyu9012qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2123456789012345678901234567890123456789"))
	}
}

func Benchmark_bytesKey2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesKey2([]byte("asdfghjklzxcvbnmqwertyu9012qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2123456789012345678901234567890123456789"))
	}
}

func Benchmark_compare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compare(
			[]byte("asdfghjklzxcvbnmqwertyu9012qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2123456789012345678901234567890123456789"),
			[]byte("asdfghjklzxcvbnmqwertyu9012qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2123456789012345678901234567890123456789"),
		)
	}
}

func Benchmark_bytesInChecksum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesInChecksum([]byte("asdfghjklzxcvbnmqwertyu9012qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2123456789012345678901234567890123456789"))
	}
}

func Benchmark_bytesInInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesInInt([]byte("asdfghjklzxcvbnmqwertyu9012qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2123456789012345678901234567890123456789"))
	}
}

func Benchmark_bytesInIntCRC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesInIntCRC([]byte("asdfghjklzxcvbnmqwertyu9012qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2123456789012345678901234567890123456789"))
	}
}
