package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filename := "./extras/rabbits.txt"

	//read in file, one command loads file into content variable
	//if an error occurs returns it as the second return value
	//if no error, err will be nil
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error reading file", filename)
	}

	//content returned as []byte not string
	//so must cast to string first and the can display
	fmt.Println(string(content))

	outfile := "output.txt"

	err = ioutil.WriteFile(outfile, content, 0644)

	if err != nil {
		log.Fatalln("Error writing file: ", err)
	}

}
