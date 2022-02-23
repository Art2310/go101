package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	y := float64(0)
	for i := 1; i <= 10; i++ {
		y = z
		z -= (z*z - x) / (2 * z)
		if y == z {
			break
		}
		fmt.Println(z)
	}
	return z

}

func main() {
	Sqrt(4)
}
