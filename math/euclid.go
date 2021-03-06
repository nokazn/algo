package main

import (
	"fmt"
	"os"
)

func scanf(f string) func(*int) int {
	return func(p *int) int {
		input, err := fmt.Scanf(f, p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return input
	}
}

func euclid(a, b int) int {
	r := a % b
	if r == 0 {
		return b
	}
	return euclid(b, r)
}

func main() {
	var a, b int

	fmt.Println("Please input 2 numbers.")
	scan := scanf("%d")
	scan(&a)
	scan(&b)
	fmt.Printf(
		"The greatest common divisor of %d and %d is %d.\n",
		a,
		b,
		euclid(a, b),
	)
}
