package main

import (
	"testing"
)

func BenchmarkChain(b *testing.B) {
	data, conv := loadInput("input.txt")

	for i := 0; i < b.N; i++ {
		part1(data, conv)
	}
	b.ReportAllocs()
}
