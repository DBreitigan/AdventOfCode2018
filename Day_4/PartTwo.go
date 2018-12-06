package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("DBreitigan/AdventOfCode2018/Day_4/input.txt")
	var guardEvents [1045]event
	var dateLayout = "2006-01-02 15:04"

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//create events from input file
	var counter = 0
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), "]")

		t, err := time.Parse(dateLayout, line[0][1:])

		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(t)

		e := event{t, line[1]}
		guardEvents[counter] = e
		counter++
	}

	//Sort the events
	sort.Slice(guardEvents[:], func(i, j int) bool {
		if guardEvents[i].timestamp.Before(guardEvents[j].timestamp) {
			return true
		} else {
			return false
		}
	})



	var fellAsleepAt int
	var currentGuard string

	var mostSleptGuard string

	maxTimeSleptInAMin := 0
	mostSleptMin := 0

	resultMap := make(map[string]guardData)
	for _, e := range guardEvents {
		timeEvent := strings.Split(e.event, " ")
		//fmt.Println(e.timestamp)
		switch timeEvent[1] {
		case "Guard":

			currentGuard = timeEvent[2]
			fellAsleepAt = -1
		case "wakes":
			sleepTime := e.timestamp.Minute() - fellAsleepAt

			currentGuardData, present := resultMap[currentGuard]

			totalSleep := 0
			var sleepPerMin [60]int
			if present {
				totalSleep += currentGuardData.totalSleep
				sleepPerMin = currentGuardData.sleepPerMin
			}

			totalSleep += sleepTime
			for i := fellAsleepAt; i < e.timestamp.Minute(); i++ {
				sleepPerMin[i]++
				if sleepPerMin[i] > maxTimeSleptInAMin {
					maxTimeSleptInAMin = sleepPerMin[i]
					mostSleptGuard = currentGuard
					mostSleptMin = i
				}
			}

			resultMap[currentGuard] = guardData{totalSleep, sleepPerMin}

		case "falls":
			fellAsleepAt = e.timestamp.Minute()
		}
	}

	fmt.Println(mostSleptGuard)
	fmt.Println(mostSleptMin)
}

type guardData struct {
	totalSleep int
	sleepPerMin [60]int
}

type event struct {
	timestamp time.Time
	event     string
}
