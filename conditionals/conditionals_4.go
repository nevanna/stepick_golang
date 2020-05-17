package main

import "fmt"

//10000

func main() {
	var a int
	var b int = 10000
	fmt.Scan(&a)
	//fmt.Scan(&a)
	if a/b == 0 {
		b /= 10
	} else {
		fmt.Println(a / b)
		return
	}
	//1000
	if a/b == 0 {
		b /= 10
	} else {
		fmt.Println(a / b)
		return
	}
	//100
	if a/b == 0 {
		b /= 10
	} else {
		fmt.Println(a / b)
		return
	}
	//10
	if a/b == 0 {
		b /= 10
	} else {
		fmt.Println(a / b)
		return
	}
	fmt.Println(a / b)
}
