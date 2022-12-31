package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name        string      `json:"username"`              // we can change the name of a key
	Password    string      `json:"-"`                     // visible to other packages, but hidden for json encoder
	Permissions permissions `json:"permissions,omitempty"` // omit empty will hide the key when it has a zero values
	// `json:",omitempty"` if we doesn't want to change the key name, only omit, we have to remember about comma
	// hidden to other packages
	hidden string // it won't be encoded, because it's not visible, private, lowercase
}

type permissions map[string]bool

func main() {
	users := []user{
		{"john", "paffaord", permissions{"write": true}, ""},
		{"kate", "paasdd", permissions{"read": true}, "secure"},
		{"martin", "paasdd", nil, "secure"},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
