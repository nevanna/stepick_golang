package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)

	if n <= 0 {
		// fmt.Printf("nb %2.2f is wrong", n)
		fmt.Printf("число %2.2f не подходит", n)
	} else if n > 10000 {
		fmt.Printf("%e", n)
	} else {
		n = n * n
		fmt.Printf("%.4f", n)
	}

}
