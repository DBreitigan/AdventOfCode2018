package main

import "fmt"

func main() {
	//input 793061
	input := 793061

	var first, last *node

	first = &node{3, nil, nil}
	last = &node{7, first, nil}
	first.next = last

	elfOnePos := first
	elfTwoPos := last

	numNodes := 2

	for numNodes < input+10 {
		elfOneValue := elfOnePos.val
		elfTwoValue := elfTwoPos.val

		total := elfOneValue + elfTwoValue

		//add nodes
		if total >= 10 {
			firstVal := total / 10
			secondVal := total % 10

			firstNode := &node{firstVal, last, nil}
			last.next = firstNode
			last = firstNode

			firstNode = &node{secondVal, last, nil}
			last.next = firstNode
			last = firstNode
			numNodes += 2
		} else {
			firstNode := &node{total, last, nil}
			last.next = firstNode
			last = firstNode
			numNodes++
		}

		//move elves
		for i := 0; i < elfOneValue + 1; i++ {
			if elfOnePos.next == nil {
				elfOnePos = first
			} else {
				elfOnePos = elfOnePos.next
			}
		}

		for i := 0; i < elfTwoValue + 1; i++ {
			if elfTwoPos.next == nil {
				elfTwoPos = first
			} else {
				elfTwoPos = elfTwoPos.next
			}
		}
	}
	currentNode := last

	for i := 0; i < numNodes - input -1; i++ {
		currentNode = currentNode.previous
	}

	count := 0
	for currentNode != nil && count < 10 {
		fmt.Print(currentNode.val, "")
		currentNode = currentNode.next
		count++
	}

}

type node struct {
	val            int
	previous, next *node
}
