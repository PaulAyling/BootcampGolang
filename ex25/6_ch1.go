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

func main() {

	username := "jack"
	password := "1888"

	args := os.Args

	if len(args) >3 || len(args) <3{
		fmt.Println("Please enter both username AND password")
		return
	} 

	if args[1]==username && args[2] == password {
		fmt.Println("Access Granted")
	} else if args[1]!=username {
		fmt.Println("Invalid username" ,args[1])
	} else if args[1]==username &&  args[2]!=password {
		fmt.Println("Invalid Password ", args[1])
	}

}