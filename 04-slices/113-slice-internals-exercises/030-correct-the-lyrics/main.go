package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//
//  [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
// STEPS
//
//  INITIAL SLICE:
//    [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  2. Put the words to the correct positions in the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
//  3. Print the `lyric` slice.
//
//
// BONUS
//
//   + Think about when does the append allocate a new backing array.
//
//   + Check whether your conclusions are correct.
//
//
// HINTS
//
//   If you get stuck, check out the hints.md file.
//
// ---------------------------------------------------------

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:

	//  1. Prepend "yesterday" to the `lyric` slice.
	lyric = append([]string{"yesterday"}, lyric...)

	// 2. Put the words to the correct positions in the `lyric` slice.

	// 9 oh I believe in yesterday
	//end := append([]string(nil), lyric[9:]...)
	lyric = append(lyric, lyric[8:13]...)
	lyric = append(lyric[:8], lyric[13:]...)
	fmt.Printf("%s\n", lyric)

}
