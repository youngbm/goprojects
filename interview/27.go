package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return []string{"North", "East", "South", "West"}[d] //  [] == [...]
	//return [...]string{"North", "East", "South", "West"}[d]
}

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3}, //  这里结尾要加  逗号,
}

// Math -> *Math 变为 可以寻址了
var mm = map[string]*Math{
	"foo": &Math{2, 3}, //  这里结尾要加  逗号,
}

func main() {
	fmt.Println(South) // 就是打印  fmt.Println(South.String())

	// Next
	m["foo"] = Math{30, 40}

	// [Importan] go 中的 map 的 value 本身是不可寻址的, 使用临时变量,或者变为可以寻址的 &

	// m["foo"].x = 30 //cannot assign to struct field m["foo"].x in map
	c := m["foo"] // 使用临时变量，赋值到别的变量也不能修改原来的 m["foo"]
	c.x = 100
	fmt.Println(c.x)      // 100
	fmt.Println(m["foo"]) // 30
}
