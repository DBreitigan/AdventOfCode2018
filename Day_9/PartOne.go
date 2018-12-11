package main

import "fmt"

func main() {
	//435 players; last marble is worth 71184 points
	players := 435
	numMarbles := 71184

	var playerScores []int
	var unplayedMarbles []int
	var playedmarbles []int

	for i := 0; i < players; i++ {
		playerScores = append(playerScores, 0)
	}

	//create list of marbles
	for i := 2; i < numMarbles; i++ {
		unplayedMarbles = append(unplayedMarbles, i)
	}

	playedmarbles = append(playedmarbles, 0, 1)

	currentPlayer := 2
	currentPos := 1

	for len(unplayedMarbles) > 0 {
		lowestMarble := unplayedMarbles[0]
		unplayedMarbles = unplayedMarbles[1:]
		fmt.Println(lowestMarble)

		if lowestMarble % 23 == 0 {
			playerScores[currentPlayer] += lowestMarble

			for i := 0; i < 7; i++ {
				currentPos--
				if currentPos == -1 {
					currentPos = len(playedmarbles) -1
				}
			}

			playerScores[currentPlayer] += playedmarbles[currentPos]
			playedmarbles = append(playedmarbles[0: currentPos], playedmarbles[currentPos + 1:]...)

		} else {
			for i:= 0; i < 2; i++ {
				currentPos++
				if currentPos == len(playedmarbles) {
					currentPos = 0
				}
			}

			playedmarbles  = append(playedmarbles , 0)
			copy(playedmarbles[currentPos+1:], playedmarbles[currentPos:])
			playedmarbles[currentPos] = lowestMarble
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
