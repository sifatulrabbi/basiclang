package main

import (
	"fmt"
	"os"
	"os/user"

	"basiclang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, this is BasicLang programming language.\nFeel free to type in commands.\n",
		user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
