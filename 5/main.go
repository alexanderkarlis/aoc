package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	uids := []int{}
	for i := 0; i < len(split); i++ {
		if split[i] == "" {
			continue
		}
		var row, col int
		upperRow := 127.0
		lowerRow := 0.0
		upperCol := 7.0
		lowerCol := 0.0
		for j := 0; j < len(split[i]); j++ {
			// fmt.Println(string(split[i][j]))
			middleRow := math.Floor((upperRow + lowerRow) / 2)
			middleCol := math.Floor((upperCol + lowerCol) / 2)

			if j == 6 {
				upperRow = float64(middleRow)
				lowerRow = float64(lowerRow)
				if string(split[i][j]) == "F" {
					row = int(lowerRow) + 1
				} else {
					row = int(upperRow) + 1
				}
			} else if string(split[i][j]) == "F" {
				upperRow = float64(middleRow)
				lowerRow = float64(lowerRow)
			} else if string(split[i][j]) == "B" {
				upperRow = float64(upperRow)
				lowerRow = float64(middleRow)
			}

			if j == 9 {
				upperCol = float64(middleCol)
				lowerCol = float64(lowerCol)
				if string(split[i][j]) == "L" {
					col = int(lowerCol) + 1
				} else {
					col = int(upperCol) + 1
				}
			} else if string(split[i][j]) == "L" {
				upperCol = float64(middleCol)
				lowerCol = float64(lowerCol)
			} else if string(split[i][j]) == "R" {
				upperCol = float64(upperCol)
				lowerCol = float64(middleCol)
			}
		}
		fmt.Printf("row: %d , column: %d\n", row, col)
		uids = append(uids, row*8+col)
	}
	sort.Ints(uids)
	fmt.Println(uids)
	maxId := uids[len(uids)-1]
	minId := uids[0]
	missingSeats := []int{}
	fmt.Println(uids)

	for i := minId; i < maxId; i++ {
		if _, in := Find(uids, i); !in {
			missingSeats = append(missingSeats, i)
		}
	}
	for _, seat := range missingSeats {
		_, in1 := Find(uids, seat+1)
		_, in2 := Find(uids, seat-1)
		if in1 && in2 {
			// fmt.Println("HERE1", seat)
		}
	}
	fmt.Println(missingSeats)
}

func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
