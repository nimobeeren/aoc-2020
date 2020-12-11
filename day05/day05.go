package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func readBinaryString(bin string) int {
	n := 0
	for i, char := range bin {
		if char == '1' {
			n += int(math.Pow(2., float64(len(bin)-i-1)))
		}
	}
	return n
}

func decode(seat string) (int, int) {
	var rowBinStr, colBinStr string
	for i, char := range seat {
		if i < 7 {
			if char == 'F' {
				rowBinStr += "0"
			} else {
				rowBinStr += "1"
			}
		} else {
			if char == 'L' {
				colBinStr += "0"
			} else {
				colBinStr += "1"
			}
		}
	}

	row := readBinaryString(rowBinStr)
	col := readBinaryString(colBinStr)

	return row, col
}

func seatID(row, col int) int {
	return row*8 + col
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Could not read input")
	}

	lines := strings.Split(string(input), "\n")
	max := 0
	var otherSeats []int
	for _, seat := range lines {
		id := seatID(decode(seat))
		otherSeats = append(otherSeats, id)
		if id > max {
			max = id
		}
	}

	includes := func(target int) bool {
		for _, candidate := range otherSeats {
			if candidate == target {
				return true
			}
		}
		return false
	}

	mySeat := -1
	for i := 1; i < max; i++ {
		if (includes(i-1) && !includes(i) && includes(i+1)) {
			mySeat = i
			break
		}
	}

	fmt.Println(max)
	fmt.Println(mySeat)
}
