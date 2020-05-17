package main

import (
	"fmt"
)

const (
	a = iota + 1
	_
	b
	c
	d = c + 2
	t
	i
	k
	i2 = iota + 2
	i3
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(t)
	fmt.Println(i)
	fmt.Println(k)
	fmt.Println(i2)
	fmt.Println(i3)
}
