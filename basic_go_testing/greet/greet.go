package greet

import "fmt"

// Hello - prints out Hello to the Person Name provided
func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

// Page - Print out a message asking each person who hasn't checked in to do so
func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; please see the front desk to check in.\n", name)
		}
	}
}
