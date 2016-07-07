/**
 * The io/ioutil examples previously worked on the whole file at once
 * The majority of the time, that is all you need.However, if for some
 * reason you want to read or write line-by-line
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "./extras/rabbits.txt"

	//read the file by using buffer io
	//The Scan() method scans token by token, default token is Scanlines
	//Scanwords and other methods avaliable
	f, _ := os.Open(filepath)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("output.txt")

	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}

	defer f.Close()

	for _, str := range []string{"apple", "banana", "carrot"} {
		bytes, err := f.WriteString(str)

		if err != nil {
			log.Fatalln("Error writing string: ", err)
		}

		fmt.Println("Wrote %d bytes to file\n", bytes)
	}
}
