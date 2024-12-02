package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")

	dataArr := strings.Split(string(input), "\n")

	if (err != nil) {
		log.Fatalf("Failed to read file: %v", err)
	}

	fmt.Printf("Solution of Part 1: %d\n", partOne(dataArr))
	fmt.Printf("Solution of Part 2: %d\n", partTwo(dataArr))
}

func partTwo(input []string) int {
	locationIdsLeft := make([]int, 0, 1000)
	locationIdsRight := make([]int, 0, 1000)

	output := 0

	for _, item := range input {
		rows := strings.Fields(item)
		leftId, _  := strconv.Atoi(rows[0])
		rightId, _ := strconv.Atoi(rows[1])

		locationIdsLeft = append(locationIdsLeft, leftId)
		locationIdsRight = append(locationIdsRight, rightId)
	}

	for _, leftId := range locationIdsLeft {
		occurence := 0

		for _, rightId := range locationIdsRight {
			if (leftId == rightId) {
				occurence++
			}
		}

		output += leftId * occurence
	}

	return output
}

func partOne(input []string) int {
	locationIdsLeft := make([]int, 0, 1000)
	locationIdsRight := make([]int, 0, 1000)

	for _, item := range input {
		rows := strings.Fields(item)
		id1, _  := strconv.Atoi(rows[0])
		id2, _ := strconv.Atoi(rows[1])

		locationIdsLeft = append(locationIdsLeft, id1)
		locationIdsRight = append(locationIdsRight, id2)
	}

	sort.Ints(locationIdsLeft)
	sort.Ints(locationIdsRight)

	output := 0

	for i := 0; i < len(locationIdsLeft); i++ {
		if (locationIdsLeft[i] < locationIdsRight[i]) {
			output += locationIdsRight[i] - locationIdsLeft[i]
		} else {
			output += locationIdsLeft[i] - locationIdsRight[i]
		}
	}

	return output
}
