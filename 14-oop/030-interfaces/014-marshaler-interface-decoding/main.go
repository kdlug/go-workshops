package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `[
 {
  "Title": "Star Wars II",
  "Price": 20,
  "Released": -62135596800
 },
 {
  "Title": "Lego",
  "Price": 30,
  "Released": -62135596800
 },
 {
  "Title": "Little prince",
  "Price": 5.5,
  "Released": 1671953739
 },
 {
  "Title": "Moby Dick",
  "Price": 5,
  "Released": 1672299339
 }
]`

func main() {
	var l list
	err := json.Unmarshal([]byte(data), &l)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(l)
}
