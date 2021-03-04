package main

import (
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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

func GreatestCommonDivisor(a, b int) int {
	fmt.Println(a, b)
	n := 1
	for i := min(a, b); i > 1; i-- {
		if b%i == 0 && a%i == 0 {
			return i
		}
	}
	return n
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
		GreatestCommonDivisor(a, b),
	)
}
