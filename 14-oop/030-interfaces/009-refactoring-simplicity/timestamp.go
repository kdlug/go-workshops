package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) string() string {
	if ts.IsZero() {
		return "N/A"
	}

	const Layout = "01/02 03:04:05PM '06 -0700"
	layout := "2006 Jan"

	return ts.Format(layout)
}

func toTimestamp(v interface{}) timestamp {
	var t int

	switch v := v.(type) { // we have to assign it to v variable, because we have to extract a type
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts := timestamp{
		time.Unix(int64(t), 0),
	}

	return ts
}
