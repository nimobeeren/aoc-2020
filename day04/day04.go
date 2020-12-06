package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Could not read input")
	}

	passports := strings.Split(string(input), "\n\n")

	numValid := 0
	for _, passport := range passports {
		isValid := true

		for _, field := range requiredFields {
			if !strings.Contains(passport, field+":") {
				isValid = false
				break
			}
		}

		if isValid {
			numValid++
		}
	}

	fmt.Println(numValid)
}
