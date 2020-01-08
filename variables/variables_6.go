package main

import "fmt"

func main() {
	var a int
	var h int
	var m int
	fmt.Scan(&a)
	h = a / 30
	m = (a % 30) * 2
	fmt.Println("It is", h, "hours", m, "minutes.")

}
