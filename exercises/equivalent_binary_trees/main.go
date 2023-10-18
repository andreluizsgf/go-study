package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	// using defer to close channel after all nodes are read
	defer close(ch)

	// using closure func to walk through all tree nodes recursively
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}

	walk(t)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// compare values pushed into both channels.
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 {
			return false
		}

		if !ok1 || !ok2 {
			break
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(5), tree.New(5)))
}
