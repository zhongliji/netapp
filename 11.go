package main

import "fmt"

func changeNum(num1 int, num2 int) (int, int) {
	var t int
	t = num1
	num1 = num2
	num2 = t
	return num1, num2
}

func main() {
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("num1的值为: %v   num2的值为: %v\n", num1, num2)
	s1, s2 := changeNum(num1, num2)
	fmt.Printf("num1的值为: %v   num2的值为: %v", s1, s2)
}
