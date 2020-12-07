package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	strs := strings.Fields(string(input))
	values := make([]int, len(strs))
	for i, str := range strs {
		values[i], err = strconv.Atoi(str)
		if err != nil {
			fmt.Println("Failed to parse int, skipping", err)
		}
	}

	for i := range values {
		for j := i; j < len(values); j++ {
			val1, val2 := values[i], values[j]
			if val1+val2 == 2020 {
				fmt.Println("Part 1:", val1, " * ", val2, " = ", val1*val2)
			}
		}
	}

	for i := range values {
		for j := i; j < len(values); j++ {
			for k := j; k < len(values); k++ {
				val1, val2, val3 := values[i], values[j], values[k]
				if val1+val2+val3 == 2020 {
					fmt.Printf("Part 2: %d * %d * %d = %d\n", val1, val2, val3, val1*val2*val3)
				}
			}
		}
	}
}
