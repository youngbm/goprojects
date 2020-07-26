package main

import (
	"encoding/json"
	"fmt"
)

// 映射不许大写
type People struct {
	Name string `json: "name"`
	//name string `json: "name"`  // 小写则不能访问
}

func main() {
	var ch chan int // 读写都会堵塞
	select {
	case v1, ok := <-ch:
		fmt.Println(v1, ok)
	default:
		fmt.Println("No io event")
	}

	//  Next
	js := `{
		"name":  "seekload"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println(p)
}
