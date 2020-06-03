package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sli := make([]int, 2)
	sli[1] = 1
	for i := 2; ; i++ {
		el := sli[i-1] + sli[i-2]
		sli = append(sli, el)
		if el == n {
			fmt.Print(i)
			break
		}
		if el > n {
			fmt.Print("-1")
			break
		}
	}
}
