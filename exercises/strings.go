package main

import "fmt"

func printString(s string) {
	for _, v := range s {
		fmt.Printf("cding %x, char \"%c\"\n", v, v)
	}

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("cding %v, char \"%c\"\n", s[i], s[i], runes[i])
	}
}

func main() {
	fmt.Println("strings")
	printString("Señor!")
}
