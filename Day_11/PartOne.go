package main

import "fmt"

const gridSize = 300
const input = 5791

func main() {
	var powerGrid = buildGrid(input)

	max := -1000
	coord := "-1,-1"
	for x := 0; x < gridSize - 2; x++ {
		for y := 0; y < gridSize - 2; y++ {
			currenVal := powerGrid[x][y] + powerGrid[x + 1][y] + powerGrid[x + 2][y] + powerGrid[x][y + 1] + powerGrid[x + 1][y + 1] + powerGrid[x + 2][y + 1] + powerGrid[x][y + 2] + powerGrid[x + 1][y + 2] + powerGrid[x + 2][y + 2]

			if currenVal > max {
				max = currenVal
				coord = coordToString(x,y)
			}
		}
	}

	fmt.Println(max)
	fmt.Println(coord)
}

func coordToString(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func buildGrid(input int) [gridSize][gridSize]int {
	var powerGrid = [gridSize][gridSize]int{}

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			rackId := x + 10
			powerLevel := rackId
			powerLevel *= y
			powerLevel += input
			powerLevel *= rackId

			if powerLevel < 100 {
				powerLevel = 0
			} else {
				powerLevel %= 1000
				powerLevel /= 100
			}

			powerLevel -= 5

			powerGrid[x][y] = powerLevel
		}
	}

	return powerGrid
}
