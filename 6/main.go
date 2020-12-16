package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	groups := strings.Split(string(content), "\n\n")
	count := 0
	everyoneAnsCount := 0
	for i := 0; i <= len(groups)-1; i++ {
		person := strings.Split(groups[i], "\n")
		m := make(map[string]int)
		personLen := len(person)
		if person[len(person)-1] == "" {
			personLen = len(person) - 1
		}
		for j := 0; j < len(person); j++ {
			if person[j] == "" {
				continue
			}
			for _, ans := range person[j] {
				m[string(ans)] = m[string(ans)] + 1
			}
		}
		count += len(m)
		allAns := []string{}
		for key := range m {
			if m[key] == personLen {
				allAns = append(allAns, key)
			}
		}
		everyoneAnsCount += len(allAns)
	}
	fmt.Println(count)
	fmt.Println(everyoneAnsCount)
}
