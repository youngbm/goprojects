package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("--main-start--")

	res, err := http.Get("http://www.google.com/robots.txt")
	//res, err := http.Get("http://www.baidu.com/robots.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(res.StatusCode)
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Printf("%T", robots)
}
