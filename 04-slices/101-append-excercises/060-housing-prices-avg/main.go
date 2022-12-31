package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		location                 []string
		size, beds, baths, price []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		cols := strings.Split(row, separator)

		location = append(location, cols[0])

		n, _ := strconv.Atoi(cols[1])
		size = append(size, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		price = append(price, n)

	}

	h := strings.Split(header, separator)
	for k, v := range h {
		fmt.Printf("%-15s", v)

		if k+1%5 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Print(strings.Repeat("=", 75), "\n")

	for i := range rows {
		fmt.Printf("%-15s", location[i])
		fmt.Printf("%-15d", size[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", price[i])
		fmt.Println()
	}

	fmt.Print(strings.Repeat("=", 75), "\n")
	var total float64
	var sum float64

	for i := range size {
		sum += float64(size[i])
	}
	total = float64(len(size))
	fmt.Printf("%-15s", "")
	fmt.Printf("%-15.2f", sum/total)

	sum = 0
	for i := range beds {
		sum += float64(beds[i])
	}
	total = float64(len(beds))
	fmt.Printf("%-15.2f", sum/total)

	sum = 0
	for i := range baths {
		sum += float64(baths[i])
	}
	total = float64(len(baths))
	fmt.Printf("%-15.2f", sum/total)

	sum = 0
	for i := range price {
		sum += float64(price[i])
	}
	total = float64(len(price))
	fmt.Printf("%-15.2f", sum/total)
	fmt.Println()

}
