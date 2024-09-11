package greet

import "fmt"

func ExampleHello() {
	greeting := Hello("Bisshwajit")
	fmt.Println(greeting)

	// Output: Hello Bisshwajit
}

func ExamplePage() {
	checkIns := map[string]bool{
		"bisshwajit": false,
		"payelh":     false,
		"sinchan":    true,
		"Sipra":      false,
	}
	Page(checkIns)
	// Unordered Output:
	// Paging bisshwajit; please see the front desk to check in.
	// Paging payelh; please see the front desk to check in.
	// Paging Sipra; please see the front desk to check in.
}
