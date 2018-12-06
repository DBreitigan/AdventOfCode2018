package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("DBreitigan/AdventOfCode2018/Day_5/input.txt")
	var input string

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		input = fileScanner.Text()
	}

	count := 1

	for count > 0 {
		count = 0

		for i := 0; i < len(input) - 1; i++ {
			//if unicode.ToUpper(input[i]) == input[i + 1] {

			if input[i] == input[i + 1] + 32 || input[i] + 32 == input[i + 1] {
				count++
				input = input[0:i] + input[i+2:]
				i--
			}
		}

	}

	fmt.Println(len(input))
	fmt.Println(input)
}
