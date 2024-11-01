package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/apex-woot/monkey-v0/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is a monkey programming language\n", user.Username)
	fmt.Printf("Feel free to type comamnds\n")
	repl.Start(os.Stdin, os.Stdout)
}
