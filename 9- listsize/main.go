package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	currentNode := l.Head
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}
	return count
}
