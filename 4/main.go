package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const keyRe = `^[a-z]{3}`

var (
	compiledRe  = regexp.MustCompile(keyRe)
	goodKeys    = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	goodKeysLen = len(goodKeys)
)

func findKeysRegex(s string) []string {
	var lines []string
	for _, x := range strings.Split(s, "\n") {
		for _, y := range strings.Split(x, " ") {
			// fmt.Println(compiledRe.FindString(y))
			lines = append(lines, compiledRe.FindString(y))
		}

	}
	return lines
}

func areKeysValid(s []string) bool {
	var foundKeys []string
	cidIsMissing := true
	for _, x := range goodKeys {
		for _, y := range s {
			if y == "cid" {
				cidIsMissing = false
			}
			if x == y {
				foundKeys = append(foundKeys, y)
			}
		}
	}
	if len(foundKeys) == goodKeysLen-1 && cidIsMissing {
		fmt.Println("missing cid")
		return true
	} else if len(foundKeys) <= goodKeysLen-1 {
		fmt.Printf("only found: %d keys\n", len(foundKeys))
		return false
	} else {
		fmt.Printf("keys match: %d keys\n", len(foundKeys))
		return true
	}
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(content), "\r\n\r\n")
	good := 0
	for i := 0; i < len(split); i++ {
		keys := findKeysRegex(split[i])
		if areKeysValid(keys) {
			good++
		}
	}
	fmt.Printf("number of good passports %d\n", good)
}
