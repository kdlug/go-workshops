package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

const (
	USER, PASSWORD     string = "john", "doe"
	usage                     = "Usage: [username] [password]\n"
	errInvalidPassword        = "Invalid password for %q.\n"
	errInvalidUser            = "Access denied for %q.\n"
	msgAccessOK               = "Access granted to %q.\n"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf(usage)
		return
	}

	u, p := os.Args[1], os.Args[2]

	if u != USER {
		fmt.Printf(errInvalidUser, u)
	} else if p != PASSWORD {
		fmt.Printf(errInvalidPassword, u)
	} else {
		fmt.Printf(msgAccessOK, u)
	}
}
