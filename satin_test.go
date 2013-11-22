package main

import "testing"

func BenchmarkSingleThread(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(false)
	}
}

func BenchmarkGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(true)
	}
}