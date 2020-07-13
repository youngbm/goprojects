package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("--main-start--")
	l, err := net.Listen("tcp", "localhost:7788")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10000000; i++ {
		//for {
		conn, err := l.Accept()
		if err != nil {
			//log.Fatal(err)
			log.Println(err)
			continue
		}
		//handle
		go handle(conn, i)
	}
}

func handle(c net.Conn, n int) {
	defer c.Close()
	fmt.Println("Connection: ", n)
	for i := 0; i < 10000000; i++ {
		//for {
		//_, err := io.WriteString(c, time.Now().Format("2006/01/02 15:04:05\n"))
		_, err := io.WriteString(c, fmt.Sprintf("%d -- %s", i, time.Now().Format("2006/01/02 15:04:05\n")))
		if err != nil {
			log.Fatal(err)
		}

		//time.Sleep(1 * time.Microsecond)
		//  阻塞读 客户端输入内容
		/*
			i := bufio.NewReader(c)
			s, err := i.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("I got string:", s)
		*/
	}
}
