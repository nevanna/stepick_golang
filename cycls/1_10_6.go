package main

import "fmt"

func reverse_int(value int) int {
	var new_int int = 0
	for value > 0 {
		var tmp = value % 10
		new_int = new_int * 10
		new_int += tmp
		value /= 10
	}
	return new_int
}

func main() {
	var (
		a int
		b int
	)
	fmt.Scan(&a)
	fmt.Scan(&b)

	a = reverse_int(a)
	for a > 0 {
		var n int = a % 10
		var tmp = b
		for tmp > 0 {
			var t int = tmp % 10
			if t == n {
				fmt.Print(t, " ")
				break
			}
			tmp = tmp / 10
		}
		a = a / 10
	}
	fmt.Print("\n")
}
