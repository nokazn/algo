package main

import (
	"fmt"
	"os"
)

func scanf(f, message string) func(*int) int {
	return func(p *int) int {
		fmt.Print(message)
		_, err := fmt.Scanf(f, p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return *p
	}
}

func scanfIntList(p []int, length int, message string) []int {
	fmt.Print(message)
	for i := 0; i < length; i++ {
		_, err := fmt.Scan(&p[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return p
}

func linearSearch(list []int, x int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	var n, x int

	scanf("%d", "Please input the number of books: ")(&n)
	list := make([]int, n)
	scanfIntList(list, n, "Please input a list of turns of books: ")
	scanf("%d", "Please input the number you want to retrive: ")(&x)
	fmt.Printf("The index of the book with No.%d is %d.\n", x, linearSearch(list, x))
}
