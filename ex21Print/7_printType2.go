
package main
import "fmt"
// ---------------------------------------------------------
// EXERCISE: Print the Type #2
//
//  Print the type and value of 3.14 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3.14 is float64
// ---------------------------------------------------------

func main() {
	myVal :=3.14
	fmt.Printf("Type of %v is %[1]T\n",myVal)
}


