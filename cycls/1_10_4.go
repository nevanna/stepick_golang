package main

import "fmt"

func main() {
	var (
		x int
		p int
		y int
		i int = 1
	)
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	for ; ; i++ {
		x = (x * (100 + p)) / 100
		if x >= y {
			fmt.Println(i)
			break
		}
	}
}
