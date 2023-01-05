package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return x, y
}

// func basic_add() {
// 	fmt.Println(add(62, 13))
// }

// func basic_swap() {
// 	a, b := swap("Hello", "world")
// 	fmt.Println(a+b, b)
// }

func main() {
	a, b := swap("Hello", "world")
	fmt.Println(a+b, b)

	fmt.Println(add(62, 13))
}
