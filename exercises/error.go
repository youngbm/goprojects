package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// No.1
	f, err := os.Open("/tmp/123")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "is  not found!")
	} else {
		fmt.Println("Open file ", f.Name(), "successfully!!")
	}

	// No.2
	addr, err := net.LookupHost("www.baidu.com1")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("Operation timeout")
		} else if err.Temporary() {
			fmt.Println("Temporary error")
		} else {
			fmt.Println("generic err: ", err)
		}
	}

	// No.3

}
