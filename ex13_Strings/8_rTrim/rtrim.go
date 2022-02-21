
package main

import (
	"fmt"
	"strings"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Right Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     only the right-most part of the given string
//  3. Trim it from the right part only
//  4. Print the number of characters it contains.
//
// RESTRICTION
//  Your program should work with unicode string values.
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func main() {
	msg := os.Args[1]
	trimmed := strings.TrimRight(msg , " ")
	fmt.Println(trimmed)

}