package main

import "fmt"

func main() {
	fmt.Println(Add(10, 10))
}

func Add(a int, b int) int {
	return a + b
}

func Substract(a int, b int) int {
	return a - b
}

func Times(a int, b int) int {
	return a * b
}

func AddX(a int, b int) int {
	return a + b + a
}

func SubstractX(a int, b int) int {
	return a + b - a
}
