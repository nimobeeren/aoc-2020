package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Could not read input")
	}

	valid1 := 0
	valid2 := 0
	for true {
		var lower int
		var upper int
		var char string
		var password string
		n, _ := fmt.Fscanf(file, "%d-%d %1s: %s\n", &lower, &upper, &char, &password)
		if n == 0 {
			break
		}

		// fmt.Printf("n:%d lower:%d upper:%d char:%s password:%s\n", n, lower, upper, char, password)

		count := strings.Count(password, char)
		if (lower <= count && count <= upper) {
			valid1++
		}

		a := string(password[lower - 1]) == char
		b := string(password[upper - 1]) == char
		if a != b {
			valid2++
		}
	}

	fmt.Println(valid1)
	fmt.Println(valid2)
}
