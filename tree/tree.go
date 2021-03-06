//package tree has methods for initialization of random binary tree with specified length,
// walking the tree and displaying it with the help of the channel
package tree

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

//initialize new channel and start tree walk, return channel with sequence of tree symbols in it
func Walk(t *Tree) <-chan int {
	ch := make(chan int)
	go func() {
		walkTree(t, ch)
		close(ch)
	}()
	return ch
}

func walkTree(t *Tree, ch chan int) {
	if t.Left != nil {
		walkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = Insert(t, (1+v)*k)
	}
	return t
}

func Insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = Insert(t.Left, v)
		return t
	}
	t.Right = Insert(t.Right, v)
	return t
}

func PrintTree(t *Tree) {
	c := Walk(t)
	for i:= range c {
		fmt.Print(i, " ")
	}
	fmt.Println()
}