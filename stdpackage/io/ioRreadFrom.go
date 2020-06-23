package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("vim-go")
	//use ioutil to print
	s, err := ioutil.ReadFile("io2-create.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

	// use writer.readFrom to print
	file, err := os.Open("io2-create.txt")
	file3, err := os.Open("io3-create.txt")
	if err != nil {
		panic(err)
	}
	//defer file.Close()
	defer func() {
		file.Close()
		file3.Close()
	}()

	write := bufio.NewWriter(os.Stdin)
	_, err = write.ReadFrom(file)
	if err != nil {
		panic(err)
	}
	write.Flush()
}
