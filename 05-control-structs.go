package main

import (
	"fmt"
)

func main() {
	//define an int variable
	num := 1

	if num > 3 {
		fmt.Println("Many")
	}

	if num == 1 {
		fmt.Println("one")
	} else if num == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("Many")
	}

	//Switch statement takes conditionals and auto break
	switch {
	case num == 1:
		fmt.Println("one")
	case num == 2:
		fmt.Println("two")
	default:
		fmt.Println("Many")
	}
}
