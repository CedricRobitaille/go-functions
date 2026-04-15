package main

import "fmt"

func concat(s1, s2 string) string {
	return s1 + s2
}

func main() {
	// concatText()
}

func concatText() {
	fmt.Println(concat("Hello", " world"))
}

// Paired with the unit test for monthly prices
func getMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 10000
	case "premium":
		return 15000
	case "enterprise":
		return 50000
	default:
		return 0
	}
}
