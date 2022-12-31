package main

import "fmt"

func main() {
	var (
		counter int
		V       int
		P       *int // zero value of pointer is nil
	)
	P = &counter // & gets an address of counter variable

	if P == nil {
		fmt.Printf("P is nil and it's adress is %p\n", P) // hex
	}

	if P == &counter {
		fmt.Printf("P is equal to counter's address: %p == %p\n", P, &counter)
	}

	counter = 100
	V = *P // copy of the value
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println("\n******** change counter")

	counter = 1
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V) // V = 100 because it's a copy

	fmt.Println("\n******** change counter in passVal()")

	counter = passVal(counter)
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)

	fmt.Println("\n******** change counter in passVal()")
	passPtrVal(&counter)
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
}

func passVal(n int) int {
	n = 50
	fmt.Printf("n      : %-13d addr: %-13p\n", n, &n)

	return n
}

func passPtrVal(pn *int) { // it changes the input value, we don't have to return anything
	// pointers are copied
	fmt.Printf("pn     : %-13p addr: %-13p *pn: %-13d\n", pn, &pn, *pn)
	*pn++
	fmt.Printf("pn     : %-13p addr: %-13p *pn: %-13d\n", pn, &pn, *pn)
}
