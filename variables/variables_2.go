package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	var r int = a*a + b*b
	fmt.Println(r)
}
