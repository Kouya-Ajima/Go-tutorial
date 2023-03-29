package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, c chan int) {
	defer close(c)
	walk(t, c)
}

func walk(t *Tree, c chan int) {
	if t == nil {
		return
	}
	walk(t.Left, c)
	c <- t.Value
	walk(t.Right, c)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {

	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(2), New(2)))
	fmt.Println(Same(New(3), New(3)))

	fmt.Println(Same(New(1), New(2)))
	fmt.Println(Same(New(1), New(3)))
	fmt.Println(Same(New(1), New(4)))

	fmt.Println(Same(New(2), New(1)))
	fmt.Println(Same(New(3), New(1)))
	fmt.Println(Same(New(4), New(1)))
}
