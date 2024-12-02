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

	if (err != nil) {
		log.Fatalf("Failed to read file: %v", err)
	}

	dataArr := strings.Split(string(input), "\n")

	locationIds1 := make([]int, 0)
	locationIds2 := make([]int, 0)

	for _, item := range dataArr {
		rows := strings.Fields(item)
		id1, _  := strconv.Atoi(rows[0])
		id2, _ := strconv.Atoi(rows[1])

		locationIds1 = append(locationIds1, id1)
		locationIds2 = append(locationIds2, id2)
	}

	sort.Ints(locationIds1)
	sort.Ints(locationIds2)

	output := 0

	for i := 0; i < len(locationIds1); i++ {
		if (locationIds1[i] < locationIds2[i]) {
			output += locationIds2[i] - locationIds1[i]
		} else {
			output += locationIds1[i] - locationIds2[i]
		}
	}

	fmt.Printf("Solution: %d\n", output)
}

