package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	// input := "HASH"
	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

	answer := solve(input)

	if answer != 1320 {
		t.Errorf("answer should be 1320. you got %d", answer)
	}
	fmt.Println(answer)
}

func TestP2(t *testing.T) {
	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

	answer := solveP2(input)
	if answer != 145 {
		t.Errorf("answer should be 145. you got %d", answer)
	}
	fmt.Println(answer)

}
