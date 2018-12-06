package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("DBreitigan/AdventOfCode2018/Day_5/input.txt")
	var fileInput string
	var AValue uint8 = 65

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fileInput = fileScanner.Text()
	}

	smallestCount := 500000

	for charVal := AValue; charVal < AValue+26; charVal++ {
		var input = fileInput
		for i := 0; i < len(input); i++ {
			if input[i] == charVal || input[i] == charVal+32 {
				input = input[0:i] + input[i+1:]
				i--
			}
		}

		count := 1
		for count > 0 {
			count = 0
			for i := 0; i < len(input)-1; i++ {
				if input[i] == input[i+1]+32 || input[i]+32 == input[i+1] {
					count++
					input = input[0:i] + input[i+2:]
					i--
				}
			}

		}

		if len(input) < smallestCount {
			smallestCount = len(input)
		}
	}

	fmt.Println(smallestCount)
}
