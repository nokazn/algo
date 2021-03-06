package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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

func randomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3) == 1
}

func makeList(l int) []int {
	list := make([]int, l)
	for i := 0; i < l; i++ {
		if i == 0 {
			list[i] = 0
		} else if randomBool() {
			list[i] = list[i-1]
		} else {
			list[i] = list[i-1] + 1
		}
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

	for start < end {
		mid := (start + end) / 2
		fmt.Println(start, mid, end)
		if list[mid] < x {
			start = mid + 1
		} else {
			end = mid
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
	fmt.Println(list)
}
