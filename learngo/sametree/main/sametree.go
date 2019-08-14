package main

/*
https://tour.golang.org/concurrency/7
Exercise: Equivalent Binary Trees
There can be many different binary trees with the same sequence of values stored at the leaves. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.


A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's concurrency and channels to write a simple solution.

This example uses the tree package, which defines the type:

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.


*/

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walkRecurse(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}
	if t.Left != nil {
		walkRecurse(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkRecurse(t.Right, ch)
	}
	return
}

func Walk(t *tree.Tree, ch chan int) {
	walkRecurse(t, ch)
	close(ch)
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1ch, t2ch := make(chan int), make(chan int)
	go Walk(t1, t1ch)
	go Walk(t2, t2ch)

	for i := range t1ch {
		if i != <-t2ch {
			return false
		}
	}
	return true
}

func main() {
	t1, t2 := tree.New(1), tree.New(1)
	fmt.Printf("are same : %t", Same(t1, t2))

}
