package main

import (
	"fmt"
	"./tree"
	)

func main() {
	fmt.Println("Printing the tree:")
	t1 := tree.New(10, 3)
	tree.PrintTree(t1)
}