package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, _ := ioutil.ReadFile("file.txt")
	fmt.Print(string(data))

	var file string
	fmt.Scanf("%s", &file)
	dataAgain, _ := ioutil.ReadFile(file)
	fmt.Print(string(dataAgain))

	f, _ := os.Open(file)

	b1 := make([]byte, 5)
	n1, _ := f.Read(b1)
	fmt.Printf("%d bytes @ %s \n", n1, string(b1))

	p2, _ := f.Seek(6, 0) // Position
	b2 := make([]byte, 2)
	n2, _ := f.Read(b2)
	fmt.Printf("%d bytes @ %d : %s", n2, p2, string(b2))
}
