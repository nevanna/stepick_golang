package main

import "fmt"

//100000
//613244
func main() {
	var a int
	fmt.Scan(&a)
	var sum_left int = 0
	var sum_right int = 0
	sum_right += a % 10
	a /= 10
	sum_right += a % 10
	a /= 10
	sum_right += a % 10
	a /= 10

	sum_left += a % 10
	a /= 10
	sum_left += a % 10
	a /= 10
	sum_left += a % 10
	if sum_left == sum_right {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
