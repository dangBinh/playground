package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	d1 := []byte("Hello world!")
	err := ioutil.WriteFile("file1.txt", d1, 0644)

	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("file2.txt")

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, _ := f.Write(d2)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, _ := f.WriteString("write\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()
}
