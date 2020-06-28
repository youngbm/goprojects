package main

import (
	"fmt"
	"log"
)

type user struct {
	s string
}

func (u user) String() string {
	return fmt.Sprintf("I am  a strings: %s", u.s)
}

func main() {
	fmt.Println("--main-fmt-start--")
	s := user{"xiezhiyang"}

	fmt.Println(s)
	fmt.Printf("%s\n", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%T\n", s)
	fmt.Println()
	fmt.Printf("%b\n", 123)
	fmt.Printf("%d\n", 0x123)
	fmt.Printf("%14.4f\n", 987654321.123456789)

	fmt.Printf("%q\n", 11)
	var a int
	a = 123
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", a)

	fmt.Println()
	n, err := fmt.Printf("|%-3s|", "a")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	//scan
	var x string
	fmt.Println("Scan")
	fmt.Scan(&x)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//scanf
	fmt.Println("Scanf")
	_, err = fmt.Scanf("%s", &x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//scanln
	fmt.Println("Scanln")
	_, err = fmt.Scanln(&x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)
	fmt.Printf("%d\n", x)

	//Print
	fmt.Print(x)

	// +  operator
	fmt.Println("a" + "b" + "c" + "d")
	//Stringer
	u := user{"xiezhiyang"}
	fmt.Println(u)
}
