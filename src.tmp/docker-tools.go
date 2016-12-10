package main

import (
    "fmt"
    "bytes"
    "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    rWalk(t, ch)
    close(ch)
}

func rWalk(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        rWalk(t.Left, ch)
    }
    if t.Right != nil {
        rWalk(t.Right, ch)
    }
    ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    var ret_val bool = false

    ch1 := make(chan int)
    ch2 := make(chan int)

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    var a_tour bytes.Buffer
    var b_tour bytes.Buffer
    select {
        case a := <- ch1:
            b_tour
        case a := <- ch2:
    }

    return ret_val
}

func main() {
    walk := make(chan int)
    go Walk(tree.New(1), walk)
    for a := range walk {
        fmt.Println(a)
    }

    fmt.PrintF

}
