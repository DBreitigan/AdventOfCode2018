package main

import "fmt"

const gridSize = 300
const input = 5791

func main() {
	var powerGrid = buildGrid(input)

	max := -1000
	coord := "-1,-1,-1"

	for currentSize := 1; currentSize <= gridSize; currentSize++ {
		fmt.Println(currentSize)
		for x := 0; x < gridSize-currentSize+1; x++ {
			for y := 0; y < gridSize-currentSize+1; y++ {
				currentVal := 0

				for i := 0; i < currentSize; i++ {
					for j := 0; j < currentSize; j++ {
						currentVal += powerGrid[x + i][y + j]
					}
				}

				if currentVal > max {
					max = currentVal
					coord = coordToString(x, y, currentSize)
				}
			}
		}
	}

	fmt.Println("max: ", max)
	fmt.Println("coord: ",coord)
}

func coordToString(x, y, size int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, size)
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
