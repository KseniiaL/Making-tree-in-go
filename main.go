//package main uses methods from tree and printing the result values
package main

import (
	"fmt"
	"github.com/KseniiaL/Making-tree-in-go/tree"
)

func main() {
	fmt.Println("Printing the tree:")
	t1 := tree.New(10, 3)
	tree.PrintTree(t1)
}
