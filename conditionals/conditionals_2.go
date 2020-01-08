package main

import "fmt"

func main() {

	var a int
	var first int
	var second int
	var third int
	fmt.Scan(&a)
	first = a / 100
	a %= 100
	second = a / 10
	third = a % 10
	//fmt.Println(first, second, third)
	if first != second && first != third && second != third {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
