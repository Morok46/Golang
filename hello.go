package main

import "fmt"

// Арифметические и логические операции
// && (конъюнкция, логическое умножение)
//|| (дизъюнкция, логическое сложение)

func main() {
	var (
		a1      = 5
		a2      = 6
		a3 bool = a1 >= 5 || a2 < 3
	)
	// a1++
	// a1++
	// a1 = a1 + a1
	//var a2 = 5
	//fmt.Println(a1 + a2 + a3)
	fmt.Println(a3)
}

/*
// Константы
func main() {
	//const Pi float32 = 3.14

	//const Pi = 3.14

	const (
		a = 1
		b
		c
		d = 3
		e
	)

	fmt.Println(a, b, c, d, e)
}
*/

/*
// Пременные
func main() {
	//var Hello string
	//Hello = "Hi all"

	//var Hello string = "Hi all"

	var (
		Name string = "Serg"
		Age  int    = 22
	)
	//Hello := "Hay u"
	fmt.Println(Name)
	fmt.Println(Age)
}
*/

/*
// Приветствие
func main() {
	fmt.Println("Hello Go! Test Git 4")
}
*/
