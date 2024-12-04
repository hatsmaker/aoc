package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")

	if (err != nil) {
		log.Fatalf("Failed to read file: %v", err)
	}

	inputStr := string(input)

	fmt.Printf("Solution of Part 1: %d\n", partOne(inputStr))
	fmt.Printf("Solution of Part 2: %d\n", partTwo(inputStr))
}

func partTwo(input string) int {
	semaphore := true

	stopInstruction := "don't()"
	startInstruction := "do()"

	str := ""

	for index := 0; index < len(input); index++ {
		if index > len(stopInstruction) {
			subStr := input[index - len(stopInstruction):index]

			if (subStr == stopInstruction) {
				semaphore = false
			}
		}

		if index > len(startInstruction) {
			subStr := input[index - len(startInstruction):index]

			if (subStr == startInstruction) {
				semaphore = true
			}
		}

		if (semaphore) {
			str += string(input[index]);
		}
	}

	return partOne(str)
}

func partOne(input string) int {
	output := 0

	regex := regexp.MustCompile(`mul\((\d+,\d+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		res := strings.Split(match[1], ",")
		output += parseInt(res[0]) * parseInt(res[1])
	}

	return output
}

func parseInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}

