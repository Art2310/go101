package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
/*
func fibonacci() func() int {
	numbers := make([]int, 0)

	return func() int {
		var number int
		switch len(numbers) {
		case 0:
			number = 0
			numbers = append(numbers, number)
			break
		case 1:
			number = 1
			numbers = append(numbers, number)
		default:
			number = numbers[0] + numbers[1]
			numbers = append(numbers, number)
			newNumbers := numbers[1:]
			numbers = newNumbers
		}
		return number
	}
}*/

func fibonacci() func() int {
	var x, y int = 0, 1

	return func() int {
		number := x
		x, y = x+y, x

		return number
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", f())
	}
}
