package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFreqRange(s string) []int {
	var array []int
	splitRange := strings.Split(s, "-")

	lowInt, err := strconv.Atoi(splitRange[0])
	if err != nil {
		panic(err)
	}

	highInt, err := strconv.Atoi(splitRange[1])
	array = append(array, lowInt)
	array = append(array, highInt)
	return array
}

func getLetter(s string) string {
	return strings.ReplaceAll(s, ":", "")
}

func isLetterFrequencyValid(min, max int, letter, seq string) bool {
	if strings.Count(seq, letter) > max || strings.Count(seq, letter) < min {
		return false
	}
	return true
}

func isLetterPositioningValid(firstPos, secondPos int, letter, seq string) bool {
	if string(seq[firstPos-1]) == letter && string(seq[secondPos-1]) == letter {
		return false
	}

	if string(seq[secondPos-1]) == letter {
		return true
	}

	if string(seq[firstPos-1]) == letter {
		return true
	}

	return false
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	passwordLines := strings.Split(string(content), "\n")

	isValid := 0
	for _, line := range passwordLines {
		if line == "" {
			continue
		}
		var freq, letter, seq string

		// split each line on space
		splitLine := strings.Split(line, " ")
		freq = splitLine[0]
		letter = splitLine[1]
		seq = splitLine[2]
		freqRange := getFreqRange(freq)
		actualLetter := getLetter(letter)

		if isLetterPositioningValid(freqRange[0], freqRange[1], actualLetter, seq) {
			isValid += 1
		}
	}
	fmt.Printf("Number of valid = %d\n", isValid)
}
