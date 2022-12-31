package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	product
	published interface{}
}

func (b *book) print() {
	b.product.print()
	p := format(b.published)
	fmt.Printf("\t - (%v)\n", p)
}

func format(v interface{}) string {
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
