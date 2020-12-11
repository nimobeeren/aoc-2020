package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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

	numYesAny := 0
	numYesAll := 0
	for _, group := range groups {
		groupNumYesAny := 0
		groupNumYesAll := 0
		for _, question := range "abcdefghijklmnopqrstuvwxyz" {
			if strings.ContainsRune(group, question) {
				groupNumYesAny++
			}

			allAnsweredYes := func(question rune) bool {
				for _, member := range strings.Split(group, "\n") {
					if len(member) > 0 && !strings.ContainsRune(member, question) {
						return false
					}
				}
				return true
			}

			if allAnsweredYes(question) {
				groupNumYesAll++
			}
		}

		numYesAny += groupNumYesAny
		numYesAll += groupNumYesAll
	}

	fmt.Println(numYesAny)
	fmt.Println(numYesAll)
}
