package main

import (
	"fmt"
	"os"
)

func scanf(f string) func(string) int {
	return func(message string) int {
		var i int
		fmt.Print(message)
		_, err := fmt.Scanf(f, &i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return i
	}
}

func scanfIntList(length int, message string) []int {
	list := make([]int, length)
	fmt.Print(message)
	for i := 0; i < length; i++ {
		_, err := fmt.Scan(&list[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return list
}

func makeList(l int) []int {
	list := make([]int, l)
	for i := 0; i < l; i++ {
		list[i] = i
	}
	return list
}

func binarySearch(list []int, x int) int {
	start, end := 0, len(list)
	if end == 0 {
		return -1
	}
	if x < list[start] || list[end-1] < x {
		return -1
	}

	for start <= end-1 {
		i := (start + end) / 2
		center := list[i]
		if x == center {
			return i
		}
		if x < center {
			end = i
		} else {
			start = i
		}
	}

	if list[start] == x {
		return start
	}
	return -1
}

func main() {
	scanInt := scanf("%d")
	n := scanInt("Please input the number of books: ")
	// list := scanfIntList(n, "Please input a list of turns of books: ")
	list := makeList(n)
	x := scanInt("Please input the number of books: ")
	fmt.Printf("The index of the book with No.%d is %d.\n", x, binarySearch(list, x))
}
