package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	concatText()
}

func concatText() {
	fmt.Println(concat("Hello", " world"))
}
