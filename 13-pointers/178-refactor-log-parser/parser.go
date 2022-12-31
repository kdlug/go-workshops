package main

import (
	"fmt"
	"strconv"
	"strings"
)

// r stores the parsed r for a domain
type r struct {
	domain string
	visits int
	// add more metrics if needed
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]r // metrics per domain
	domains []string     // unique domain names
	total   int          // total visits for all domains
	lines   int          // number of parsed lines (for the error messages)
	lerr    error        // last error
}

func newParser() *parser {
	return &parser{sum: make(map[string]r)}
}

func parse(p *parser, line string) (parsed r) {
	//for in.Scan() {
	//	p.lines++
	// var (
	// 	parsed r
	// 	err    error
	// )
	// Parse the fields
	if p.lerr != nil {
		return
	}

	p.lines++
	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	parsed.domain = fields[0]

	// Sum the total visits per domain
	var err error
	parsed.visits, err = strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
	}

	return
}

func update(p *parser, parsed r) {
	if p.lerr != nil {
		return
	}
	domain, visits := parsed.domain, parsed.visits
	// Collect the unique domains
	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}

	// Keep track of total and per domain visits
	p.total += visits

	// You cannot assign to composite values
	// p.sum[domain].visits += visits

	// create and assign a new copy of `visit`
	p.sum[domain] = r{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}
}

func err(p *parser) error {
	return p.lerr
}
