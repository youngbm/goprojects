package main

import "fmt"

type c interface {
	Calculator() float32
}

type manager struct {
	name   string
	salary float32
	bonus  float32
}

type staff struct {
	name   string
	salary float32
}

func (x manager) Calculator() float32 {
	return x.salary + x.bonus
}

func (x staff) Calculator() float32 {
	return x.salary
}

type myInt int16

func (x myInt) myFuncName(string) {
	fmt.Println("I am so smart!")

}

func testFunc(s string) {
	fmt.Println(s)

}

func main() {
	var x myInt
	x = 123
	x.myFuncName("hah")

	fmt.Println("Using Interface")
	//staff
	var staff1, staff2 staff
	staff1 = staff{
		"Xzy",
		10000,
	}
	staff2 = staff{
		"Fjh",
		8000,
	}
	//manager
	var manager1, manager2 manager
	manager1 = manager{
		"John",
		25000,
		-1000,
	}
	manager2 = manager{
		"John",
		35000,
		1000,
	}

	var department = []c{staff1, staff2, manager1, manager2}
	var salarySum float32
	for _, v := range department {
		salarySum += v.Calculator()
	}

	fmt.Println(salarySum)
}
