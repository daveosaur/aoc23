package main

import (
	"os"
	"strings"
	"testing"
)

func TestInput(t *testing.T) {
	input, _ := os.ReadFile("testinput.txt")
	p1 := part1(string(input))
	if p1 != 13 {
		t.Errorf("got %d. Should be: 13\n", p1)
	}
	p2 := part2(string(input))
	if p2 != 30 {
		t.Errorf("got %d. Should be: 30\n", p2)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, _ := os.ReadFile("input.txt")

	for n := 0; n < b.N; n++ {
		part2(string(input))
	}
	b.ReportAllocs()
}

func BenchmarkParse(b *testing.B) {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	start := strings.Index(lines[0], ":") + 1
	data := strings.Split(lines[0][start:], "|")

	for n := 0; n < b.N; n++ {
		parseCard(data[1])
	}
	b.ReportAllocs()
}
