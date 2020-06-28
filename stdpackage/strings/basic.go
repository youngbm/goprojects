package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("--main-start--")
	//count
	fmt.Println(strings.Count("123123123123123abdsdsfsdfsf", "1"), "1")
	fmt.Println(strings.Count("123123123123123abdsdsfsdfsf", "2"), "2")
	fmt.Println(strings.Count("123123123123", "3"), "3")
	fmt.Println(strings.Count("1231231sssssssssssssss231", "s"), "s")
	fmt.Println(strings.Compare("asdf", "b"))
	fmt.Println(strings.Join([]string{"1", "2", "3", "4"}, "|"))
	fmt.Println(strings.Split("1|2|3|4|5|6|7|8|9|0", "4"))

	fmt.Println(strings.Index("123333", "3"))
	fmt.Println(strings.LastIndex("123333", "3"))
	fmt.Println(strings.IndexRune("123333", rune('3')))
	fmt.Println(strings.LastIndexAny("123333", "xxas1f"))
	fmt.Println(strings.Repeat("123\n", 10))

	s := strings.NewReader("abcdefghijklmnopqrstuvwxyz我")
	fmt.Println(s.Len())
	fmt.Println(s.Size())
	r, err := s.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	//fmt.Println(s.WriteTo(os.Stdin))
}
