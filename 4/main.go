package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	keyRe      = `^[a-z]{3}`
	hexColorRe = `^#(?:[0-9a-fA-F]{3}){1,2}$`
)

var (
	compiledRe  = regexp.MustCompile(keyRe)
	hcRe        = regexp.MustCompile(hexColorRe)
	goodKeys    = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	goodKeysLen = len(goodKeys)
)

func mapBuilder(s string) map[string]string {
	var lines []string

	m := make(map[string]string)
	for _, x := range strings.Split(s, "\n") {
		for _, y := range strings.Split(x, " ") {
			splitKv := strings.Split(y, ":")
			m[splitKv[0]] = strings.Replace(splitKv[1], "\r", "", -1)
			lines = append(lines, compiledRe.FindString(y))
		}
	}
	return m
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
		return true
	} else if len(foundKeys) <= goodKeysLen-1 {
		return false
	} else {
		return true
	}
}

func validateDateRange(s string, d1, d2 int) bool {
	sint, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	if d1 <= sint && sint <= d2 {
		return true
	}
	return false
}

func validateHeight(s string) bool {
	var min, max int
	if len(s) <= 3 {
		return false
	}
	sint, err := strconv.Atoi(s[:len(s)-2])
	if err != nil {
		panic(err)
	}

	if s[len(s)-2:] == "cm" {
		min = 150
		max = 193
	} else if s[len(s)-2:] == "in" {
		min = 59
		max = 76
	}
	if min <= sint && sint <= max {
		return true
	}
	return false

}

func validateHair(s string) bool {
	return hcRe.MatchString(s)
}

func validateEye(s string) bool {
	for _, x := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if x == s {
			return true
		}
	}
	return false
}

func validatePID(s string) bool {
	_, err := strconv.Atoi(s)
	if len(s) == 9 && err == nil {
		return true
	}
	return false
}

func checkMapValues(m map[string]string, keys []string) bool {
	var good []bool
	for _, key := range keys {
		switch key {
		case "byr":
			good = append(good, validateDateRange(m[key], 1920, 2002))
		case "iyr":
			good = append(good, validateDateRange(m[key], 2010, 2020))
		case "eyr":
			good = append(good, validateDateRange(m[key], 2020, 2030))
		case "hgt":
			good = append(good, validateHeight(m[key]))
		case "hcl":
			good = append(good, validateHair(m[key]))
		case "ecl":
			good = append(good, validateEye(m[key]))
		case "pid":
			good = append(good, validatePID(m[key]))
		case "cid":
		default:
		}
	}
	for i := 0; i < len(good); i++ {
		if good[i] == false {
			return false
		}
	}
	return true
}

func isMapValid(m map[string]string) bool {
	var foundKeys []string
	cidIsMissing := true
	for _, x := range goodKeys {
		for k := range m {
			if k == "cid" {
				cidIsMissing = false
			}
			if x == k {
				foundKeys = append(foundKeys, k)
			}
		}
	}
	if len(foundKeys) == goodKeysLen-1 && cidIsMissing { // missing only `cid` code
		return checkMapValues(m, foundKeys)
	} else if len(foundKeys) <= goodKeysLen-1 { // missing another key, don't bother checking values
		return false
	} else { // nothing missing, proceed to check keys
		return checkMapValues(m, foundKeys)
	}
}

func pp(x map[string]string) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(content), "\r\n\r\n")
	good := 0
	for i := 0; i < len(split); i++ {
		m := mapBuilder(split[i])
		if isMapValid(m) {
			good++
		}
	}
	fmt.Printf("number of good passports %d\n", good)
}
