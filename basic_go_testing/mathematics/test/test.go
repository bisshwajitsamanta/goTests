package main

import (
	"fmt"
	"mathematics"
)

func main() {
	sum := mathematics.Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("FAIL: Wanted 11 but got %d", sum)
		panic(msg)
	}
	println("PASS")
}
