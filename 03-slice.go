package main

import "fmt"

func main() {
	//initialize an empty array
	var empty []int

	alphas := []string{"abc", "def", "ghi", "jkl"}

	// slice can not be modified, a new copy needs to be made
	// here is a example that appends elements to a slice
	empty = append(empty, 123)
	empty = append(empty, 456)
	fmt.Printf("%v \n", empty)

	alphas = append(alphas, "pqr", "stu")
	fmt.Printf("%v \n", alphas)

	//get length of a slice
	fmt.Println("Lenth: ", len(alphas))

	//retrieve a signle element  from slice
	fmt.Println(alphas[1])

	//retrieve a slice of a slice
	alphas2 := alphas[1:3]
	fmt.Println(alphas2)

	//check if element exists in array
	if elemExits("def", alphas) {
		fmt.Println("Exists!")
	}

}

func elemExits(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
