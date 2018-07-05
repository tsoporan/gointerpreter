package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome to the REPL %s!\n", user.Username)
	fmt.Printf("Type in commands ...\n")
	repl.Start(os.Stdin, os.Stdout)
}
