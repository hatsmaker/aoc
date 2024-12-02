package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")

	if (err != nil) {
		log.Fatalf("Failed to read file: %v", err)
	}

	dataArr := strings.Split(string(input), "\n")

	fmt.Printf("Solution of Part 1: %d\n", partOne(dataArr))
	fmt.Printf("Solution of Part 2: %d\n", partTwo(dataArr))
}

func partTwo(input []string) int {
	output := 0

	for _, fields := range input {
		row := strings.Fields(fields)

		var isSafe bool = false

		if isRowSafe(row) {
			isSafe = true
		} else {
			for index := 0; index < len(row); index++ {
				rowCopy := make([]string, len(row))
				copy(rowCopy, row)

				if isRowSafe(append(rowCopy[:index], rowCopy[index+1:]...)) {
					isSafe = true
					break
				}
			}
		}

		if isSafe {
			output++
		}
	}

	return output
}

func partOne(input []string) int {
	output := 0

	for _, fields := range input {
		row := strings.Fields(fields)

		var isSafe bool = isRowSafe(row)

		if isSafe {
			output++
		}
	}

	return output
}

func isRowSafe(row []string) bool {
	upperBound := 3
	lowerBound := 1

	var isIncreasing bool = parseInt(row[1]) - parseInt(row[0]) > 0

	for i := 1; i < len(row); i++ {
		previous := parseInt(row[i-1])
		current := parseInt(row[i])

		diff := current - previous
		absDiff := absInt(diff)

		var condition bool = diff == 0 ||
			(absDiff < lowerBound || absDiff > upperBound) ||
			(isIncreasing && diff < 0) ||
			(!isIncreasing && diff > 0)

		if condition {
			return false
		}
	}

	return true
}

func absInt(n int) int {
    if n < 0 {
        return -n
    }

    return n
}

func parseInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}

