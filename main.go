package main

import (
	"fmt"
	)

func main() {
	fmt.Println("Start")
	t1 := New(10, 3)
	PrintTree(t1)
	fmt.Println("end")
}
