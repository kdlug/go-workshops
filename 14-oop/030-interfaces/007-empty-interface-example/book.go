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
	fmt.Printf("Title: %s, price: %.2f, published: %v\n", b.title, b.price, published)
}

func format(published interface{}) string {

	if published == nil {
		return "N/A"
	}

	var t int

	if v, ok := published.(int); ok {
		t = v
	}

	if v, ok := published.(string); ok {
		t, _ = strconv.Atoi(v)
	}

	ut := time.Unix(int64(t), 0)

	return ut.String()
}
