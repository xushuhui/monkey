package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	currentuser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s", currentuser.Username)
	repl.Start(os.Stdin, os.Stdout)
}
