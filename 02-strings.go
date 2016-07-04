package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "HI, I'M UPPER CASE"
	//convert to lower case
	lower := strings.ToLower(str)

	fmt.Println(lower)

	//check if a string contain other string
	if strings.Contains(lower, "case") {
		fmt.Println("Yes, exits ...")
	}

	//strings are arrays of characters
	fmt.Println("Characters 3-9: " + str[3:9])

	fmt.Println("First Five: " + str[:5])

	//split a string
	sentence := "I'm a sentence made up of words"
	words := strings.Split(sentence, " ")

	fmt.Printf("%v \n", words)

	//If you are spliting on whitespace, using Fields is better becase
	//it will split on more than just the space, but all whitespace chars
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)
}
