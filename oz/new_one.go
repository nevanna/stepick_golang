package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var s string = strings.Replace(string(data), "\n", "", -1)
	var spllitted []string = strings.Split(s, " ")
	// fmt.Println(spllitted[0])
	var i int = 0
	var sum int = 0
	for ; i < len(spllitted); i++ {

		nb, err := strconv.Atoi(spllitted[i])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		sum += nb
		fmt.Println(sum)
	}
	newData := []byte(strconv.Itoa(sum))
	err = ioutil.WriteFile("output.txt", newData, 0777)
	if err != nil {
		// print it out
		fmt.Println(err)
	}

}
