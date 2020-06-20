package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func sayHello(c chan int) {
	fmt.Println("Subprocess: Hi world!~~~~~")
	c <- 1000
}

func input(c chan int) {
	var i = 0
	for {
		c <- i
		i++
		//time.Sleep(1000 * time.Microsecond)
	}
}

func output(c chan int) {

	for {
		fmt.Println(time.Now(), len(c))
		fmt.Println(<-c)
		time.Sleep(1 * time.Second)
	}
}

func printCap(c chan int) {

	file, err := os.OpenFile("/tmp/chanCap.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic("Can not open this file!")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	var s string
	for {
		s = fmt.Sprintln(len(c))
		writer.WriteString(s)
		writer.Flush()
	}
}

func main() {
	//fmt.Println("Main process!!!")
	c := make(chan int, 100)
	//go sayHello(c)
	//time.Sleep(1000 * time.Millisecond)
	//fmt.Println(<-c)

	go input(c)
	go printCap(c)
	output(c)
}
