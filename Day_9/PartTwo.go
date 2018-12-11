package main

import "fmt"

func main() {
	//435 players; last marble is worth 7118400 points
	players := 435
	numMarbles := 7118400

	var firstNode *node
	var lastNode *node
	var currentNode *node

	var playerScores []int
	var unplayedMarbles []int

	for i := 0; i < players; i++ {
		playerScores = append(playerScores, 0)
	}

	//create list of marbles
	for i := 2; i < numMarbles; i++ {
		unplayedMarbles = append(unplayedMarbles, i)
	}

	firstNode = &node{nil, nil, 0}
	lastNode = &node{nil, firstNode, 1}
	firstNode.next = lastNode

	currentNode = lastNode
	currentPlayer := 2

	for len(unplayedMarbles) > 0 {
		lowestMarble := unplayedMarbles[0]
		unplayedMarbles = unplayedMarbles[1:]

		if lowestMarble%23 == 0 {

			playerScores[currentPlayer] += lowestMarble

			for i := 0; i < 7; i++ {
				if currentNode == firstNode {
					currentNode = lastNode
				} else {
					currentNode = currentNode.previous
				}
			}

			playerScores[currentPlayer] += currentNode.val

			if currentNode == firstNode {
				currentNode = currentNode.next
				firstNode.next = nil
				firstNode = currentNode
			} else if currentNode == lastNode {
				lastNode = currentNode.previous
				lastNode.next = nil
				currentNode = firstNode
			} else {
				prev := currentNode.previous
				next := currentNode.next
				prev.next = next
				next.previous = prev
				currentNode = next
			}

		} else {
			if currentNode == lastNode {
				currentNode = firstNode
			} else {
				currentNode = currentNode.next
			}

			if currentNode == lastNode {
				n := node{nil, currentNode, lowestMarble}
				lastNode = &n
				currentNode.next = lastNode
				currentNode = lastNode
			} else {
				nextNode := currentNode.next
				n := node{nextNode, nextNode.previous, lowestMarble}
				currentNode.next = &n
				nextNode.previous = &n
				currentNode = currentNode.next
			}
		}

		currentPlayer++
		if currentPlayer == players {
			currentPlayer = 0
		}
	}

	max := 0
	for _, val := range playerScores {
		if val > max {
			max = val
		}
	}

	fmt.Println(max)
}

type node struct {
	next     *node
	previous *node
	val      int
}
