package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	p := newParser()
	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		parsed := parse(p, in.Text())

		update(p, parsed)
	}

	if p.lerr != nil {
		fmt.Println(p.lerr)
		return
	}

	summarize(p)
	// Let's handle the error
	// if err := in.Err(); err != nil {
	// 	fmt.Println("> Err:", err)
	// }
	dumpErrs([]error{in.Err(), err(p)})
}

func summarize(p *parser) {
	// Print the visits per domain
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		//parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

}

func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}
