/**
 * functions.go
 *
 * Examples on how to use functions in go
 */
package main

import (
	"fmt"
	"math"
)

func Echo(s string) {
	fmt.Println(s)
}

//Function with single return value
//the type of the return value is speicfied after function declaration
func Say(s string) string {
	pharse := "Hello " + s

	return pharse
}

//Function with single named return value
//you can specify the return variable name, which  initializes it
//the := notation is for new variables, and = for initial ones
//Also you do not need to include the variable in return statement
//it will return the current value of the variable at return
func Say2(s string) (phrase string) {
	phrase = "Hello " + s
	return
}

//Function with multiple parameters and return values
//If the types are the same you can specify the type once at end.
func Divide(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)

	return q, r
}

func main() {
	Echo("Bonjour tout le module")

	fmt.Println(Say("Gopher"))

	q, r := Divide(11, 3)

	fmt.Printf("Quotient: %v, Reminder: %v \n", q, r)
}
