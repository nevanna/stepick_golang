package main

import "fmt"

func main() {
	var i int = 0
	var n int
	for ; ; i++ {
		fmt.Scan(&n)
		if n < 10 {
			continue
		} else if n > 100 {
			break
		} else {
			fmt.Println(n)
		}
	}
}
