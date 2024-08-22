package mathematics

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{10, -2, 3})
	if sum != 11 {
		t.Errorf("Wanted 11 but got %d", sum)
	}
}

func TestMul(t *testing.T) {
	mul := Mul(4, 5)
	if mul != 20 {
		t.Errorf("Wanted 20 but got %d", mul)
	}
}

func TestDiv(t *testing.T) {
	div := Div(10, 2)
	if div != 5 {
		t.Errorf("Wanted 5 but got %d", div)
	}
}
