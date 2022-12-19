package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {
	currentuser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s", currentuser.Username)
	repl.Start(os.Stdin, os.Stdout)
}
