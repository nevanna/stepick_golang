package main

import "fmt"

func main() {
	var workArray [10]uint8
	var i int = 0
	var tmp [2]int
	for ; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}
	var j int = 0
	for ; j < 3; j++ {
		for i = 0; i < len(tmp); i++ {
			fmt.Scan(&tmp[i])
		}
		bubble := workArray[tmp[0]]
		workArray[tmp[0]] = workArray[tmp[1]]
		workArray[tmp[1]] = bubble
	}
	for i = 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}
