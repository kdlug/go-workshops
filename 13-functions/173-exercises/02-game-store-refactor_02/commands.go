package main

import (
	"fmt"
	"strconv"
	"strings"
)

//     1- Add a func that runs the given command from the user:
//
//        Name  : runCmd
//        Inputs: input string, []game, map[int]game
//        Output: bool
//
//        This func returns true if it wants the program to
//        continue. When it returns false, the program will
//        terminate. So, all the commands that it calls need
//        to return true or false as well depending on whether
//        they want to cause the program to terminate or not.
func runCmd(input string, games []game, byID map[int]game) bool {
	fmt.Println()

	cmd := strings.Fields(input)

	if len(cmd) == 0 {
		return true
	}

	switch cmd[0] {
	case "quit":
		// cmdQuit()
		return cmdQuit()

	case "list":
		// cmdList()
		return cmdList(games)

	case "id":
		// cmdByID
		return cmdByID(cmd, games, byID)
	}

	return true
}

//     2- Add a func that handles the quit command:
//
//        Name  : cmdQuit
//        Input : none
//        Output: bool
//
func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

//     3- Add a func that handles the list command:
//
//        Name  : cmdList
//        Inputs: []game
//        Output: bool
func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}

	return true
}

//     4- Add a func that handles the id command:
//
//        Name  : cmdByID
//        Inputs: cmd []string, []game, map[int]game
//        Output: bool
func cmdByID(cmd []string, games []game, byID map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := byID[id]
	if !ok {
		fmt.Println("sorry. I don't have the game")
		return true
	}

	printGame(g)
	return true
}
