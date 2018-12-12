package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	vectors := readInput()

	closeToResult := false
	for i := 1; i < 100000; i++ {
		move(vectors)
		res := printStars(vectors, i)

		if !closeToResult && res {
			closeToResult = true
		}
		//Break once we get too far away
		if closeToResult && !res {
			break
		}
	}
}

type star struct {
	px int
	py int
	vx int
	vy int
}

func readInput() (r []*star) {
	dat, _ := ioutil.ReadFile("DBreitigan/AdventOfCode2018/Day_10/input.txt")
	lines := strings.Split(string(dat), "\n")
	re := regexp.MustCompile(`position=<\s*(-*\d+),\s*(-*\d+)> velocity=<\s*(-*\d+),\s*(-*\d+)>$`)
	for _, l := range lines {
		matches := re.FindStringSubmatch(l)
		r = append(r, &star{toInt(matches[1]), toInt(matches[2]), toInt(matches[3]), toInt(matches[4])})
	}
	return r
}

func move(vectors []*star) {
	for _, v := range vectors {
		v.px = v.px + v.vx
		v.py = v.py + v.vy
	}
}

func printStars(vectors []*star, i int) bool {
	minX, minY, maxX, maxY := findEdges(vectors)
	m := make(map[string]string)
	for _, v := range vectors {
		m[keyToString(v.px, v.py)] = "#"
	}

	if maxY-minY < 30 {
		fmt.Println("Second: ", i)
		for y := minY; y <= maxY; y++ {
			line := ""
			for x := minX; x <= maxX; x++ {
				_, exists := m[keyToString(x, y)]
				if exists {
					line += "#"
				} else {
					line += " "
				}
			}
			fmt.Println(line)
		}
		return true
	}
	return false

}

func keyToString(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func findEdges(vectors []*star) (int, int, int, int) {
	minX, minY, maxX, maxY := 1000, 1000, -1000, -1000
	for _, v := range vectors {
		if v.px > maxX {
			maxX = v.px
		} else if v.px < minX {
			minX = v.px
		}
		if v.py > maxY {
			maxY = v.py
		} else if v.py < minY {
			minY = v.py
		}
	}
	return minX, minY, maxX, maxY
}