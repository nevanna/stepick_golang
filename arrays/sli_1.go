package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	sli := make([]int, n)
	var s string
	fmt.Scanln(&s)
	for i := 0; i < n; i++ {
		fmt.Scanln(&sli)
	}
	fmt.Println(sli[3])
}
