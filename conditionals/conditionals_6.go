package main

import "fmt"

//напомним, что год является високосным, если его номер кратен 4, но не кратен 100,
//а также если он кратен 400

func main() {
	var a int
	fmt.Scan(&a)
	if (a%4 == 0 && a%100 != 0) || (a%400 == 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
