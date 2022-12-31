package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	spendings := fetch()

	// spendings daily
	for i, daily := range spendings {
		var total int

		for _, spending := range daily {
			total += spending
		}
		fmt.Printf("Day #%d: %d\n", i+1, total)
	}

	fmt.Printf("Spendings: %v\n", spendings)
}

func fetch() [][]int {
	content := `200 100
	25 10 45 60
	5 15 35
	95 10
	50 25`

	lines := strings.Split(content, "\n")
	spendings := make([][]int, len(lines)) // length 0 can append starting from index 0

	for i, line := range lines { // each line string
		fields := strings.Fields(line)
		spendings[i] = make([]int, len(fields)) // slice is empty, so initialize a line

		for j, field := range fields {
			n, _ := strconv.Atoi(field)
			spendings[i][j] = n
		}
	}

	return spendings
	//return [][]int(nil)
}
