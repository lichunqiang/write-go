/**
 * A program showing how to use command-line arguments and flags in Go
 */
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	str  string
	num  int
	help bool
)

func main() {
	//command-line args are shorted as slice in os.Args
	//first argument in list is program itself
	num_args := len(os.Args)

	if num_args < 2 {
		fmt.Println(">> No args passed in")
	}

	//flag package provides parsing of command-line paramters
	//this example we create global variables and then pass them in
	//as pointers which BoolVar, StringVar and IntVar set as value
	flag.StringVar(&str, "str", "default value", "text description")
	flag.IntVar(&num, "num", 5, "text description")
	flag.BoolVar(&help, "help", false, "Display help")
	flag.Parse()

	if help {
		fmt.Println(">> Display help screen")
		os.Exit(1)
	}

	fmt.Println(">> String: ", str)
	fmt.Println(">> Number: ", num)

	//last command-line argument
	fmt.Println(">> Last Item: ", os.Args[num_args-1])

	//the os.Args will include flags for example
	//go run command-line-args.go --str=Foo filename
	//os.Args[1] = "--str=Foo"

	//if you have flags and want just the arguments
	//then after calling  flag.Parse() you an call
	//flag.Args which store only the args
	args := flag.Args()

	if len(args) > 0 {
		fmt.Println(">> Flag Arg:", args[0])
	}

}
