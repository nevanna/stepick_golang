package main

import "fmt"

func sum_(n int) int {
	if n < 10 {
		return n
	}
	var sum int = 0
	for ; n > 0; n /= 10 {
		sum += n % 10
	}
	v := sum_(sum)
	return v
}

func main() {
	var nb int
	fmt.Scan(&nb)
	rez := sum_(nb)
	fmt.Print(rez)
}
