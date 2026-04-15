package main

func adder() func(int) int {
	sum := 0
	return func(val1 int) int {
		sum += val1
		return sum
	}
}
