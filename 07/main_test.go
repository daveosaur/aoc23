package main

import (
	"testing"
)

func BenchmarkParseLine2(b *testing.B) {
	handStr := "QQQJA 483"

	for n := 0; n < b.N; n++ {
		parseLine2(handStr)
	}
	b.ReportAllocs()
}

func BenchmarkParseLine1(b *testing.B) {
	handStr := "QQQJA 483"

	for n := 0; n < b.N; n++ {
		parseLine(handStr)
	}
	b.ReportAllocs()
}
