package main

import "fmt"

func main() {
	var (
		a   int
		b   int
		rez int
	)
	fmt.Scan(&a)
	fmt.Scan(&b)
	rez = a + b
	fmt.Println(rez)

}
