package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var min int
	var q int
	for i := 0; i < n; i++ {
		var tmp int

		fmt.Scan(&tmp)
		if i == 0 {
			min = tmp
			q = 1
		} else {
			if tmp < min {
				min = tmp
				q = 1
			} else if tmp == min {
				q++
			}
		}
	}
	fmt.Print(q)
}

// 5
// 14
// 85
// 14
// 14
// 2
