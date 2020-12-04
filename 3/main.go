package main

import (
	"fmt"
	"io/ioutil"
	// "strconv"
	"strings"
)

type slope struct {
	x, y, tCount int
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(content), "\n")

	// currPos in the current position in the 1-D row
	var currPos int
	slopes := []*slope{{1, 1, 0}, {3, 1, 0}, {5, 1, 0}, {7, 1, 0}, {1, 2, 0}}

	for _, s := range slopes {
		currPos = 0
		for i := 0; i < len(split); i += s.y { // +/- rise
			rowLen := len(split[i])

			if rowLen == 0 {
				continue
			}
			if currPos >= rowLen {
				currPos = currPos - (rowLen)
			}
			if string(split[i][currPos]) == "#" {
				(*s).tCount += 1
			}
			currPos += s.x
		}
	}

	m := 1
	for l := 0; l < len(slopes); l++ {
		m *= slopes[l].tCount
	}
	fmt.Printf("tree count: %d", m)
}
