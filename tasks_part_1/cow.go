package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if (n >= 5 && n <= 20) || (n%10 >= 5 && n%10 <= 9) || (n%10 == 0) {
		fmt.Printf("%d korov", n)
	} else if n%10 >= 2 && n%10 <= 4 {
		fmt.Printf("%d korovy", n)
	} else if n%10 == 1 {
		fmt.Printf("%d korova", n)
	}
}
