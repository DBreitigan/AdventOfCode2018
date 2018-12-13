package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("DBreitigan/AdventOfCode2018/Day_13/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var board [150][150]rune

	var carts []cart

	fileScanner := bufio.NewScanner(file)
	boardCount := 0
	id := 0

	for fileScanner.Scan() {
		runeArray := []rune(fileScanner.Text())

		for pos, val := range runeArray {
			switch val {
			case 'v':
				board[boardCount][pos] = '|'
				carts = append(carts, cart{id, boardCount, pos, val, 0})
				id++
			case '^':
				board[boardCount][pos] = '|'
				carts = append(carts, cart{id, boardCount, pos, val, 0})
				id++
			case '<':
				board[boardCount][pos] = '-'
				carts = append(carts, cart{id, boardCount, pos, val, 0})
				id++
			case '>':
				board[boardCount][pos] = '-'
				carts = append(carts, cart{id, boardCount, pos, val, 0})
				id++
			default:
				board[boardCount][pos] = val
			}

		}

		boardCount++
	}

	//crash := false
	for len(carts) > 1 {
		var cartsToRemove []int

		for pos, val := range carts {
			if val.direction == 'v' {
				val.yPos++
				boardSpot := board[val.yPos][val.xPos]

				//Update direction if necessary
				switch boardSpot {
				case '/':
					val.direction = '<'
				case '\\':
					val.direction = '>'
				case '+':
					if val.numTurns%3 == 0 {
						val.direction = '>'
					} else if val.numTurns%3 == 2 {
						val.direction = '<'
					}
					val.numTurns++
				}

				if c, badCartId := containsCrash(carts, val.yPos, val.xPos); c {
					cartsToRemove = append(cartsToRemove, val.id)
					cartsToRemove = append(cartsToRemove, badCartId)

					fmt.Println("crash:", val.xPos, val.yPos)
				}
					carts[pos] = val

			} else if val.direction == '^' {
				val.yPos--
				boardSpot := board[val.yPos][val.xPos]

				//Update direction if necessary
				switch boardSpot {
				case '/':
					val.direction = '>'
				case '\\':
					val.direction = '<'
				case '+':
					if val.numTurns%3 == 0 {
						val.direction = '<'
					} else if val.numTurns%3 == 2 {
						val.direction = '>'
					}
					val.numTurns++
				}

				if c, badCartId := containsCrash(carts, val.yPos, val.xPos); c {
					cartsToRemove = append(cartsToRemove, val.id)
					cartsToRemove = append(cartsToRemove, badCartId)

					fmt.Println("crash:", val.xPos, val.yPos)
				}
					carts[pos] = val


			} else if val.direction == '>' {
				val.xPos++
				boardSpot := board[val.yPos][val.xPos]

				//Update direction if necessary
				switch boardSpot {
				case '/':
					val.direction = '^'
				case '\\':
					val.direction = 'v'
				case '+':
					if val.numTurns%3 == 0 {
						val.direction = '^'
					} else if val.numTurns%3 == 2 {
						val.direction = 'v'
					}
					val.numTurns++
				}

				if c, badCartId := containsCrash(carts, val.yPos, val.xPos); c {
					cartsToRemove = append(cartsToRemove, val.id)
					cartsToRemove = append(cartsToRemove, badCartId)

					fmt.Println("crash:", val.xPos, val.yPos)
				}
					carts[pos] = val


			} else if val.direction == '<' {
				val.xPos--
				boardSpot := board[val.yPos][val.xPos]

				//Update direction if necessary
				switch boardSpot {
				case '\\':
					val.direction = '^'
				case '/':
					val.direction = 'v'
				case '+':
					if val.numTurns%3 == 0 {
						val.direction = 'v'
					} else if val.numTurns%3 == 2 {
						val.direction = '^'
					}
					val.numTurns++
				}

				if c, badCartId := containsCrash(carts, val.yPos, val.xPos); c {
					cartsToRemove = append(cartsToRemove, val.id)
					cartsToRemove = append(cartsToRemove, badCartId)

					fmt.Println("crash:", val.xPos, val.yPos)
				}
					carts[pos] = val

			}

		}
		for _, val := range cartsToRemove {
			for pos, cart := range carts {
				if cart.id == val {
					carts = append(carts[:pos], carts[pos+1:]...)
					break
				}
			}
		}

	}
	fmt.Println(carts[0], carts[0].xPos, carts[0].yPos)
}

func containsCrash(carts []cart, xpos, ypos int) (bool, int) {
	for _, val := range carts {
		if val.yPos == xpos && val.xPos == ypos {
			return true, val.id
		}
	}
	return false, -1
}

type cart struct {
	id         int
	yPos, xPos int
	direction  rune
	numTurns   int
}
