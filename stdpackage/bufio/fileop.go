package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("main")
	var (
		filename = "/tmp/test284057930"
		f        *os.File
		err      error
	)
	//fInfo, err := os.Stat(filename)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(fInfo.Size())
	//fmt.Println(fInfo.Name())
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		f, err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}
	f, err = os.OpenFile("/tmp/123123", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//  io 模块去写 string
	_, err = io.WriteString(f, "io\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("os fd + io write")

	//打开的文件句柄  file.WriteString(wstring)
	_, err = f.Write([]byte("fd write\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("os fd + fd write(os single)")

	//打开的文件句柄  file.WriteString(wstring)
	_, err = f.WriteString("fd writestring\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("os fd + fd writeString(os single)")

	// bufio 模块写 需要Flush
	w := bufio.NewWriter(f)
	_, err = w.WriteString("fd + bufid\n")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("os fd + bufio ")
	w.Flush()
	f.Close()
}
