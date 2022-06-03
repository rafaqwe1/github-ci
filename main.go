package main

import "fmt"

func main() {
	result := Sum(15, 10)
	fmt.Println(result)
}

func Sum(num1 int, num2 int) int {
	return num1 + num2
}
