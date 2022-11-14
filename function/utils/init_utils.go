package utils

import "fmt"

var Name string
var Age int

func init() {
	fmt.Println("utils init()...")
	Name = "Mary"
	Age = 18
}