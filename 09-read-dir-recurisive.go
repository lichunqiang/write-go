/**
 * Read directory and recursively print on all files and directories.
 *
 */
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Walker(fn string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walker error: ", err)
		return nil
	}

	if fi.IsDir() {
		fmt.Println("Direcotry: ", fn)
	} else {
		fmt.Println("File: ", fn)
	}

	return nil
}

func main() {
	path := "./"

	filepath.Walk(path, Walker)
}
