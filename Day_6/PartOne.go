package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("DBreitigan/AdventOfCode2018/Day_6/input.txt")
	//var input string

	var maxX, maxY float64 = 0, 0



	coordList := [50]coord{}

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	counter := 0
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), ", ")
		x, _ := strconv.ParseFloat(line[0], 64)

		if x > maxX {
			maxX = x
		}
		y, _ := strconv.ParseFloat(line[1], 64)

		if y > maxY {
			maxY = y
		}

		coordList[counter] = coord{x , y}

		counter++
	}

	infiniteCoords := make(map[coord]bool)
	coordsMap := make(map[coord]int)

	for x := float64(0); x < maxX; x++ {
		for y := float64(0); y < maxY; y++ {
			minCoord := coord{0,0}
			minDist := float64(1000000)

			for _, co := range coordList {
				dist := math.Abs(x - co.X) + math.Abs(y - co.Y)

				if dist < minDist {
					minDist = dist
					minCoord = co
				} else if dist == minDist {
					minCoord = coord{-1, -1}
				}

			}

			if x == 0 || y == 0 || x == maxY || y == maxX {
				infiniteCoords[minCoord] = true
			}

			coordsMap[minCoord]++
		}
	}

	max := 0
	for k, v := range coordsMap {
		if _, found := infiniteCoords[k]; v > max && !found {
			max = v
		}
	}


	fmt.Println(max)
}

type coord struct {
	X float64
	Y float64
}
