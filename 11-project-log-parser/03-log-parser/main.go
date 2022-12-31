package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		total   int
		lines   int
		domains []string // we need a slice in order to get always the same order of domain list, because map can return it everytime in different order
	)
	in := bufio.NewScanner(os.Stdin)
	log := make(map[string]int)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("ERROR: Wrong log file format: %v (line #%d)\n", fields, lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], lines)
			return
		}

		if _, ok := log[domain]; !ok {
			domains = append(domains, domain)
		}

		log[domain] += visits
		total += visits
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range domains {
		visits := log[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	// Print the total visits for all log
	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

}
