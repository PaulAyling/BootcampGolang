package main
import "fmt"
// ---------------------------------------------------------
// EXERCISE: Print the Type #3
//
//  Print the type and value of "hello" using fmt.Printf
//
// EXPECTED OUTPUT
// 	Type of hello is string
// ---------------------------------------------------------

func main() {
	myVar :="hello"
	fmt.Printf("Type of %v is a %[1]T\n",myVar)
}