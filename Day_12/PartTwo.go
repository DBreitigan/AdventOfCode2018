package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	initialByteArray := []byte("#.##.##.##.##.......###..####..#....#...#.##...##.#.####...#..##..###...##.#..#.##.#.#.#.#..####..#")
	generations := 20

	file, err := os.Open("DBreitigan/AdventOfCode2018/Day_12/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var spreads = make(map[string]string)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		words := strings.Fields(fileScanner.Text())

		spreads[words[0]] = words[2]
	}

	fmt.Println("initial:", string(initialByteArray))
	neg := 0
	for gen := 0; gen < generations; gen++ {
		var changes []change

		if string(initialByteArray[len(initialByteArray)-5:]) != "....." {
			initialByteArray = append(initialByteArray, []byte("..............................")...)
		}

		if string(initialByteArray[0:5]) != "....." {
			initialByteArray = append([]byte(".............................."), initialByteArray...)
			neg += 30
		}

		for i := 0; i < len(initialByteArray) - 5; i++ {
			thisString := string(initialByteArray[i:i+5])
			val := spreads[thisString]
			if val == "#" {
				changes = append(changes, change{i+2, '#'})
				//initialByteArray[i+2] = '#'
			} else if val == "."{
				changes = append(changes, change{i+2, '.'})
				//initialByteArray[i+2] = '.'
			}
		}

		for _, val := range changes {
			initialByteArray[val.pos] = val.changeTo
		}

		if gen % 10 == 0 {
			fmt.Println("gen:", gen + 1, string(initialByteArray))
			fmt.Println(neg)
		}
	}

	count := 0
	for pos, val := range initialByteArray {
		if val == '#' {
			count += pos - neg
		}
	}

	fmt.Println(count)

}

type change struct {
	pos int
	changeTo byte
}