package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
    	    return
    	}
        walker(t.Left)
	    ch <- t.Value
    	walker(t.Right)
	}
	walker(t)
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1 := make(chan int)
    c2 := make(chan int)
    go Walk(t1, c1)
    go Walk(t2, c2)
    for n1 := range c1 {
        n2 := <-c2
        if n1 != n2 {
            return false
        }
    }
    return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		if i != 9 {
			fmt.Print(<-ch,",")
		} else {
			fmt.Print(<-ch)
		}
	}
	fmt.Println()
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
