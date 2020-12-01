package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./inputs.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(content), "\n")
	for i, x := range split {
		for _, y := range split[i : len(split)-1] {
			for _, z := range split[i+1 : len(split)-1] {
				numx, err := strconv.Atoi(x)
				if err != nil {
					panic(err)
				}
				numy, err := strconv.Atoi(y)
				if err != nil {
					panic(err)
				}
				numz, err := strconv.Atoi(z)
				if err != nil {
					panic(err)
				}
				if numx+numy+numz == 2020 {
					fmt.Println(numx, numy, numz)
					fmt.Println(numx * numy * numz)
				}
			}
		}
	}
}
