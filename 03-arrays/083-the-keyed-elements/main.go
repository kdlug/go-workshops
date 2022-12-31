package main

import "fmt"

const (
	ICX = iota // 0
	WAN
	ETH
)

func main() {
	// we can add indexes manually
	// rates := [...]float64{
	// 	0: 1.25,
	// 	1: 1,
	// 	2: 1.5,
	// }

	// we can eve use const
	rates := [...]float64{
		ICX: 1.25,
		WAN: 1,
		ETH: 1.5,
	}

	fmt.Printf("1 BTC is %g ICX\n", rates[ICX])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])
	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
}
