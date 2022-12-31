package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Populate and Lookup
//
//  Add elements to the maps that you've declared in the
//  first exercise, and try them by looking up for the keys.
//
//  Either use the `make()` or `map literals`.
//
//  After completing the exercise, remove the data and check
//  that your program still works.
//
//
//  1. Phone numbers by last name
//     --------------------------
//     bowen  202-555-0179
//     dulin  03.37.77.63.06
//     greco  03489940240
//
//     Print the dulin's phone number.
//
//
//  2. Product availability by Product ID
//     ----------------
//     617841573 true
//     879401371 false
//     576872813 true
//
//     Is Product ID 879401371 available?
//
//
//  3. Multiple phone numbers by last name
//     ------------------------------------------------------
//     bowen  [202-555-0179]
//     dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
//     greco  [03489940240 03489900120]
//
//     What is Greco's second phone number?
//
//
//  4. Shopping basket by Customer ID
//     -------------------------------
//     100 [617841573:4 576872813:2]
//     101 [576872813:5 657473833:20]
//     102 []
//
//     How many of 576872813 the customer 101 is going to buy?
//                (Product ID)  (Customer ID)
//
//
// EXPECTED OUTPUT
//
//   1. Run the solution to see the output
//   2. Here is the output with empty maps:
//
//      dulin's phone number: N/A
//      Product ID #879401371 is not available
//      greco's 2nd phone number: N/A
//      Customer #101 is going to buy 5 from Product ID #576872813.
//
// ---------------------------------------------------------

func main() {
	// 1st way
	var phones map[string]string

	phones = map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}

	person, number := "dulin", "N/A"

	n, ok := phones[person]
	if ok {
		number = n
	}

	fmt.Printf("%s's phone: %s\n", person, number)

	// 2nd way
	products := make(map[string]bool)

	products["617841573"] = true
	products["879401371"] = false
	products["576872813"] = true

	status := "available"

	if !products["879401371"] {
		status = "not " + status
	}

	fmt.Printf("Product 879401371 is %s\n", status)

	// 3rd way
	phoneBook := map[string][]string{}

	phoneBook["bowen"] = []string{"202-555-0179"}
	phoneBook["dulin"] = []string{"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"}
	phoneBook["greco"] = []string{"03489940240", "03489900120"}

	person, number = "greco", "N/A"

	if phones := phoneBook[person]; len(phones) == 2 {
		number = phoneBook["greco"][1]
	}

	fmt.Printf("greco 2nd phone number: %s\n", number)

	// 4th way map literal
	basket := map[int]map[int]int{
		100: {841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	//Customer #101 is going to buy 5 from Product ID #576872813.
	pId, cId := 576872813, 101
	fmt.Printf("Customer %d is going to buy %d from Product ID #%d.\n", cId, basket[cId][pId], pId)

	fmt.Printf("phones: %#v\n", phones)
	fmt.Printf("products: %#v\n", products)
	fmt.Printf("phoneBook: %#v\n", phoneBook)
	fmt.Printf("basket: %#v\n", basket)
}
