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

func sum(a, d, n int) int {
	return (a + (a + (n-1)*d)) * n / 2
}

func main() {
	var n int

	fmt.Println("Please input number.")
	scanf("%d")(&n)
	fmt.Printf("Sum of 1 to %d is %d.\n", n, sum(1, 1, n))
}
