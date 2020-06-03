package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	for i := b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Print(i)
			return
		}

	}
	fmt.Print("NO")
}
