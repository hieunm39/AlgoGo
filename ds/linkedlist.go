package main

import "fmt"


type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}


func main() {
	
	linkedList := LinkedList{nil, 0}
	fmt.Println(linkedList)
	node1 := &Node{10, nil}
	linkedList.insert(node1.data)

	fmt.Println(linkedList.head)
}	


func (l *LinkedList) insert(data int) {
	temp1 := &Node{data, nil}

	if l.head == nil {
		l.head = temp1
	} else {
		
		temp2 := l.head
		l.head = temp1
		temp1.next = temp2

	}
	l.length++
}