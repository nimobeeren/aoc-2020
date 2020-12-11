package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func setCreate() map[rune]struct{} {
	return make(map[rune]struct{})
}

func setAdd(set map[rune]struct{}, valueToAdd rune) {
	set[valueToAdd] = struct{}{}
}

func setSize(set map[rune]struct{}) int {
	return len(set)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Could not read input")
	}

	groups := strings.Split(string(input), "\n\n")

	numYesses := 0
	for _, group := range groups {
		groupYesQuestions := setCreate()
		for _, question := range group {
			if question == '\n' { continue }
			setAdd(groupYesQuestions, question)
		}
		numYesses += setSize(groupYesQuestions)
	}

	fmt.Println(numYesses)
}