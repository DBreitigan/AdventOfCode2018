package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("DBreitigan/AdventOfCode2018/Day_8/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var instructions []int64

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		words := strings.Fields(fileScanner.Text())

		for _, k := range words {
			val, _ := strconv.ParseInt(k, 10, 64)
			instructions = append(instructions, val)
		}
	}

	_, v := processChildren(instructions)

	fmt.Println(v)
}

func processChildren(input []int64) ([]int64, int64) {
	metadataVal := int64(0)
	numChildren := input[0]
	numMetadata := input[1]

	input = input[2:]

	for numChildren > 0 {
		inp, val := processChildren(input)
		input = inp
		metadataVal += val
		numChildren--
	}

	for i := int64(0); i < numMetadata; i++ {
		metadataVal += input[i]
	}

	input = input[numMetadata:]

	return input, metadataVal
}
