package mathematics

// Sum - Will add all the numbers in the slice and return the sum
func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Mul(a, b int) int {
	return a * b
}

func Div(x, y int) int {
	return x / y
}
