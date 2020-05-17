package main

import "fmt"

const delimetr int = 8

func main() {
	var nb int
	fmt.Scan(&nb)
	var sum int = 0
	var i int = 0
	for ; i < nb; i++ {
		var number int = 0
		fmt.Scan(&number)
		var isBad int = number % delimetr
		if isBad == 0 && number > 9 && number < 100 {
			sum = sum + number
		}
	}
	fmt.Println(sum)
}
