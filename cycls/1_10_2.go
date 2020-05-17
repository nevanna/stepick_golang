package main

import "fmt"

func main() {
	var (
		n int
		c int
		d int
	)
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)
	var i int = 1
	for ; i <= n; i++ {
		if i%d != 0 && i%c == 0 {
			fmt.Println(i)
			break
		}
	}
}
