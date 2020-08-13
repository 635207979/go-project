package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	Visit(t, ch)
	close(ch)
}
func Visit(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Visit(t.Left, ch)
	}
	if t.Right != nil {
		Visit(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	i := 0
	for n := range ch1 {
		fmt.Println("t1:", n)
		i ^= n
	}
	for n := range ch2 {
		fmt.Println("t2:", n)
		i ^= n
	}

	return i == 0
}

func main() {

	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	//for i := range ch {
	//	fmt.Println(i)
	//}
	fmt.Println(Same(tree.New(2), tree.New(2)))
}
