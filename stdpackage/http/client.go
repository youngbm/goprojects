package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var tr = http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    20 * time.Second,
	DisableCompression: true,
}

func main() {
	/*
			fmt.Println("--main-start--")
			c := &http.Client{Transport: tr}
			resp, err := c.Get("https://www.baidu.com")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Got body %s", resp.StatusCode)
		tr = &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}

	*/
	client := &http.Client{}
	fmt.Println("starting get baidu")
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
}
