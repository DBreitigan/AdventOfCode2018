package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("DBreitigan/AdventOfCode2018/Day_7/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	instructions := make(map[rune] []rune)
	parents := make(map[rune] int)


	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		key := rune(line[5])
		value := rune(line[36])
		instructions[key] = append(instructions[key],value)
		parents[value] = parents[value] + 1
	}

	readyTasks := make([]rune, 0)
	for k := range instructions {
		if parents[k] == 0 {
			readyTasks = append(readyTasks, k)
		}
	}

	finishedTasks := make([]rune, 0)
	workersTasks := []rune{'.', '.', '.', '.', '.'}
	TimePerWorker := []int{0, 0, 0, 0, 0}

	t := 0
	working := 1
	for ; working > 0; t++ {
		working = 0

		for n := range TimePerWorker {
			//decrease time left
			if TimePerWorker[n] != 0 {
				TimePerWorker[n] = TimePerWorker[n] - 1
				working = working + 1
			} else {
				//check if more work to do on task
				if workersTasks[n] != '.' {
					finishedTask := workersTasks[n]
					workersTasks[n] = '.'

					//check children
					for _, v := range instructions[finishedTask] {
						parents[v] = parents[v] - 1
						if parents[v] == 0 {
							readyTasks = append(readyTasks, v)
						}
					}
				}
			}
		}

		//try add new tasks to TimePerWorker
		for ; len(readyTasks) > 0 && working < len(TimePerWorker); {
			temp := make([]rune, len(readyTasks))
			copy(temp, readyTasks)
			sort.Sort(runes(temp))
			x := temp[0]
			for i := 0; i < len(readyTasks); i++ {
				if readyTasks[i] == x {
					readyTasks = append(readyTasks[:i], readyTasks[i+1:]...)
				}
			}
			finishedTasks = append(finishedTasks, x)
			for n, _ := range TimePerWorker {
				if workersTasks[n] == '.' {
					workersTasks[n] = x
					TimePerWorker[n] = int(x) - 5
					//TimePerWorker[n] = int(x) - 'A'
					working = working + 1
					break;
				}
			}
		}

		fmt.Print(t," ")
		for n, _ := range TimePerWorker {
			fmt.Print("{", TimePerWorker[n]," ",string(workersTasks[n]),"} ")
		}
		fmt.Print(string(finishedTasks))
		fmt.Println()
	}

	fmt.Println(t-1)

}

type runes []rune

func (s runes) Len() int {
	return len(s)
}

func (s runes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s runes) Less(i, j int) bool {
	return s[i] < s[j]
}
