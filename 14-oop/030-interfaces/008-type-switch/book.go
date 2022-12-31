package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	title     string
	price     float64
	published interface{}
}

func (b book) print() {
	published := format(b.published)
	fmt.Printf("Title: %s, price: %.2f, v: %v\n", b.title, b.price, published)
}

func format(v interface{}) string {

	switch v.(type) {
	case int:
		fmt.Print("int ->  ")
	case string:
		fmt.Print("string ->  ")
	default:
		fmt.Print("nil ->  ")
	}

	//--
	var t int

	switch v := v.(type) { // we have to assign it to v variable, because we have to extract a type
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "N/A"
	}

	ut := time.Unix(int64(t), 0)
	// const Layout = "01/02 03:04:05PM '06 -0700" ...
	layout := "2006 Jan"
	return ut.Format(layout)
}
