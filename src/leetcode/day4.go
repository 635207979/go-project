package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	listarray := []*ListNode{}
	list1 := head.Next
	for list1 != nil {
		listarray = append(listarray, list1)
		list1 = list1.Next
	}
	j := len(listarray) - 1
	fmt.Printf("%d", j)
	i := 0
	p := head
	for i <= j {
		if i == j {
			p.Next = listarray[j]
			break
		}
		fmt.Printf("%v", p.Val)
		p.Next = listarray[j]
		p = p.Next
		j--
		fmt.Printf("%d", j)
		p.Next = listarray[i]
		p = p.Next
		i++
	}
}
