package main

import "fmt"

func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return x, y
}

// func main() {
// 	fmt.Println(add(62, 13))
// }

func main() {
	a, b := swap("Hello", "world")
	fmt.Println(a+b, b)

	fmt.Println(add(62, 13))
}
