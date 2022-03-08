
Skip to content
Pull requests
Issues
Marketplace
Explore
@PaulAyling
inancgumus /
learngo
Public

Code
Issues 2
Pull requests 5
Discussions
Actions
Security

    Insights

learngo/11-if/exercises/04-vowel-or-cons/solution/main.go /
@inancgumus
inancgumus update: copyright
Latest commit ede39bf on 30 Oct 2019
History
1 contributor
35 lines (29 sloc) 828 Bytes
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	// I didn't use a short-if here because, it's already
	// hard to read. Do not make it harder.

	s := args[1]
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}
}
