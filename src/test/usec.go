package main

import "fmt"

func cal1(a int, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func cal2(a int, b int) (add int, mul int, sub int) {
	add = a + b
	mul = a * b
	sub = a - b
	return
}
func main() {
	x1, x2, x3 := cal1(3, 4)
	fmt.Printf("%d,%d,%d", x1, x2, x3)
}
