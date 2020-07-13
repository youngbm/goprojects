package main

import (
	"expvar"
	"fmt"
)

func main() {
	fmt.Println("--main-start--")

	var visits = expvar.NewInt("visits")
	var a = 123

	fmt.Println(visits.String())
	fmt.Println(visits)
	visits.Add(1)
	fmt.Println(visits)

	fmt.Errorf("%s", a)

}
