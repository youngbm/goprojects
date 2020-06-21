package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("vim-go")
	//// read all
	//data, err := ioutil.ReadFile("/tmp/1231")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(data))

	// read one by one line
	f, err := os.Open("/tmp/123")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
