package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result // total visits per domain
	domains []string          // we need a slice in order to get always the same order of domain list, because map can return it everytime in different order
	total   int               // total visits to all domains
	lines   int               // number of parsed lines
}

func main() {
	p := parser{
		sum: make(map[string]result),
	}

	in := bufio.NewScanner(os.Stdin)
	//log := make(map[string]int)

	for in.Scan() {
		p.lines++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("ERROR: Wrong log file format: %v (line #%d)\n", fields, p.lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], p.lines)
			return
		}

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}
		// you can only change an addresable value
		// but map elements is not an addresable value
		//p.sum[result.domain].visits += visits
		// SOLUTION:
		// create a copy
		r := result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
		p.sum[domain] = r
		p.total += visits
	}

	sort.Strings(p.domains)
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	// Print the total visits for all log
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

}
