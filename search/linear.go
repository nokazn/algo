package main

import (
	"fmt"
	"os"
)

func scanf(f string) func(*int, string) int {
	return func(p *int, message string) int {
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

func makeList(l int) []int {
	list := make([]int, l)
	for i := 0; i < l; i++ {
		list[i] = i
	}
	return list
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
	scanInt := scanf("%d")
	scanInt(&n, "Please input the number of books: ")
	// list := make([]int, n)
	// scanfIntList(list, n, "Please input a list of turns of books: ")
	list := makeList(n)
	scanInt(&x, "Please input the number you want to retrive: ")
	fmt.Printf("The index of the book with No.%d is %d.\n", x, linearSearch(list, x))
}
