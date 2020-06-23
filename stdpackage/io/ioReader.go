package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func readMenu() {
	fmt.Println("")
	fmt.Println("*******从不同的源读入数据*********")
	fmt.Println("1.表示标准输入")
	fmt.Println("2.表示普通文件")
	fmt.Println("3.表示从字符串中")
	fmt.Println("4.表示从网络中")
	fmt.Println("b返回上级菜单")
	fmt.Println("q退出菜单")
	fmt.Println("***********************************")
}

func main() {
	for {
		readMenu()
		var (
			input string
		)
		fmt.Scan(&input)
		fmt.Println("you input:", input)
		switch input {
		case "1":
			fmt.Println("最多输入9998个字节. 请用回车结束输入:")
			p, err := ReadFrom(os.Stdin, 10000)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("你输入了:", string(p))
		case "2":
			var filepath string
			fmt.Println("你输入了文件路径")
			fmt.Scan(&filepath)
			f, err := os.Open(filepath)
			if err != nil {
				log.Fatalln(err)
			}
			p, err := ReadFrom(f, 100000000)
			if err != nil {
				log.Fatal(err)
			}
			f.Close()
			fmt.Println("你的文件内容是:")
			fmt.Println("------------Start--------------")
			fmt.Println(string(p))
			fmt.Println("------------ End --------------")
		case "3":
			fmt.Println("从字符串中读入，暂未实现!")
			p, err := ReadFrom(strings.NewReader("my string!!!"), 1000)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(p))
		case "4":
			fmt.Println(input)
			fmt.Println("从网络读入，暂未实现!")
		case "b":
			fmt.Println(input)
			continue
		case "q":
			fmt.Println("退出程序！！")
			os.Exit(0)
		default:
			fmt.Println("你输入不正确， 请输入[1234bq]")
		}
	}
}
