package main

import "fmt"

func main() {

	var i int = 0
	var max int = 0
	var nbMax int = 0
	for ; ; i++ {
		var n int = 0
		fmt.Scan(&n)
		if n == max {
			nbMax++
		}
		if n > max {
			nbMax = 1
			max = n
		}
		if n == 0 {
			break
		}
	}
	fmt.Println(nbMax)
}
