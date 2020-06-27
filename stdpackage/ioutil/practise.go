package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("main")
	// 读文件
	p, err := ioutil.ReadFile("/tmp/testIOutil.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	// 创建目录

	dirname, err := ioutil.TempDir("/tmp/", "test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dirname)
	//  创建临时文件
	filename, err := ioutil.TempFile(dirname, "testfile")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(filename.Name())
	fmt.Printf("You created %s\n", filename.Name())
}
