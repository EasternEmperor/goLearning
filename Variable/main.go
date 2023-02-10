package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b
		c, d = iota, iota
	)
	fmt.Println(a, b, c, d)

	const (
		e = iota
		f
		g = iota
		h 
	)
	fmt.Println(e, f, g, h)
}