package main

import "fmt"

func main() {

	var numbers [5]int

	numbers = [5]int{1, 2, 3}
	numbers = [5]int{4: 7}
	numbers = [5]int{1, 2, 3, 4: 7}

	fmt.Println(numbers)
}
