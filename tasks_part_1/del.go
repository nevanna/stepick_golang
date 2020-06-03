package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	var d string
	fmt.Scan(&d)
	n = strings.Replace(n, d, "", -1)
	fmt.Print(n)
}
