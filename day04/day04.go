package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func hasAllRequiredFields(passport string) bool {
	for _, field := range requiredFields {
		if !strings.Contains(passport, field+":") {
			return false
		}
	}
	return true
}

var fieldPattern = regexp.MustCompile(`(\S+):(\S+)`)
var hgtValuePattern = regexp.MustCompile(`(\d+)(cm|in)`)
var hclValuePattern = regexp.MustCompile(`#([a-f]|[0-9]){6}`)
var eclValuePattern = regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
var pidValuePattern = regexp.MustCompile(`^\d{9}$`)

func validateField(fieldName, fieldValue string) bool {
	switch fieldName {
	case "byr":
		byr, err := strconv.Atoi(fieldValue)
		if err != nil || byr < 1920 || byr > 2002 {
			return false
		}
	case "iyr":
		iyr, err := strconv.Atoi(fieldValue)
		if err != nil || iyr < 2010 || iyr > 2020 {
			return false
		}
	case "eyr":
		eyr, err := strconv.Atoi(fieldValue)
		if err != nil || eyr < 2020 || eyr > 2030 {
			return false
		}
	case "hgt":
		heightValue := hgtValuePattern.FindStringSubmatch(fieldValue)
		if heightValue == nil {
			return false
		}
		heightNumber, err := strconv.Atoi(heightValue[1])
		if err != nil {
			return false
		}
		heightUnit := heightValue[2]
		switch heightUnit {
		case "cm":
			if heightNumber < 150 || heightNumber > 193 {
				return false
			}
		case "in":
			if heightNumber < 59 || heightNumber > 76 {
				return false
			}
		default:
			// invalid unit
			return false
		}
	case "hcl":
		if !hclValuePattern.MatchString(fieldValue) {
			return false
		}
	case "ecl":
		if !eclValuePattern.MatchString(fieldValue) {
			return false
		}
	case "pid":
			if !pidValuePattern.MatchString(fieldValue) {
				return false
			}
	}
	return true
}

func hasOnlyValidValues(passport string) bool {
	fields := fieldPattern.FindAllStringSubmatch(passport, -1)
	if fields == nil {
		// passport has no fields
		return true
	}
	for _, field := range fields {
		if !validateField(field[1], field[2]) {
			return false
		}
	}
	return true
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Could not read input")
	}

	passports := strings.Split(string(input), "\n\n")

	numValid1 := 0
	numValid2 := 0
	for _, passport := range passports {
		if hasAllRequiredFields(passport) {
			numValid1++

			if hasOnlyValidValues(passport) {
				numValid2++
			}
		}
	}

	fmt.Println(numValid1)
	fmt.Println(numValid2)
}
