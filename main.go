package main

import "fmt"

func Add(a, b int) int {
	if (b == 0) {
		return a
	}
	if (a == 0) {
		return b
	}
	return a + b
}

func main() {
	fmt.Println("Hello world from CI")
}