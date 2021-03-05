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

func sum(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	var n int

	fmt.Println("Please input number.")
	scanf("%d")(&n)
	fmt.Printf("Sum of 1 to %d is %d.\n", n, sum(n))
}
