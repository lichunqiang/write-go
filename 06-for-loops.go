/**
 * Go uses "for" for all loops, there is no "while" or "do-while" loops
 */
package main

import (
	"fmt"
)

func main() {
	//create an array of strings
	alphas := []string{"abc", "def", "ghi"}

	//standard for loop
	for i := 0; i < len(alphas); i++ {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	//if iterating over an array, this would be how you would actually do it
	//the standard loop would be used if uncommon step function
	//range return two values, i -> index, val -> value
	for i, v := range alphas {
		fmt.Printf("%d: %s \n", i, v)
	}

	//if you did not care about the value and only wanted the index
	for i := range alphas {
		fmt.Println(i)
	}

	//if you did not care about the index and only the value, you use the _
	//which means don't save this value
	for _, val := range alphas {
		fmt.Println(val)
	}

	x := 0

	//how to use for like whild
	for x < 10 {
		fmt.Println(x)
		x++
	}

	//infinite loop
	// for {
	// 	fmt.Println(".")
	// }
}
