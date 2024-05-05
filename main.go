package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func copyToDestination(src, dest string) (string, error) {
	// The defer keyword executes the close statement when this function returns

	source, err := os.Open(src)
	check(err)
	defer source.Close()

	destination, err := os.Create(dest)
	check(err)
	defer destination.Close()
	_, err = io.Copy(destination, source)
	check(err)

	return dest, err
}

func main() {
	files, err := os.ReadDir("./folderA")
	check(err)
	for _, file := range files {
		go func() {
			src := "./folderA/" + file.Name()
			dest := "./folderB/" + file.Name()
			_, err = copyToDestination(src, dest)
			if err != nil {
				fmt.Printf("%s -> %s ❌\n", src, dest)
			} else {
				fmt.Printf("%s -> %s ✅\n", src, dest)
			}
		}()
	}
}
