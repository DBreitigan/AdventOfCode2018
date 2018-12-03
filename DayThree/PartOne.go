package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("DBreitigan/AdventOfCode2018/DayThree/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	cloth := [1200][1200]int{}

	overlappedSquares := 0

	//Create cloth and find overlapping squares
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		words := strings.Fields(fileScanner.Text())

		coords := strings.Split(words[2], ",")
		xcoord, _ := strconv.Atoi(coords[0])
		ycoord, _ := strconv.Atoi(coords[1][0:len(coords[1]) - 1])

		size := strings.Split(words[3], "x")
		rightSize, _ := strconv.Atoi(size[0])
		downSize, _ := strconv.Atoi(size[1])

		for i := ycoord; i < ycoord + downSize; i++ {
			for j := xcoord; j < xcoord + rightSize; j++ {
				cloth[j][i]++
				if cloth[j][i] == 2 {
					overlappedSquares++
				}

			}
		}
	}

	fmt.Println(overlappedSquares)
}
