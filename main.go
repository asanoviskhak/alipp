package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/asanoviskhak/alipp/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Салам %s! Бул alipp программалоо тили!\n", user.Username)
	fmt.Printf("Өзүңүз каалагандай тилди изилдеп көрүңүз\n\n")
	fmt.Printf("Бул жерден чыгуу үчүн Ctrl жана C баскычтарын басыңыз\n\n")
	repl.Start(os.Stdin, os.Stdout)
}