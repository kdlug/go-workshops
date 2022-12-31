package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type permissions map[string]bool

	type user struct {
		Name        string      `json:"username"`
		Permissions permissions `json:"permissions,omitempty"` // omit empty will hide the key when it has a zero values
	}

	var input []byte
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		input = append(input, in.Bytes()...)
	}

	var users []user

	err := json.Unmarshal(input, &users)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", users)
}
