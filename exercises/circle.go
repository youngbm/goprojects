package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (c *areaError) Error() string {
	return fmt.Sprintf("You give me a nagetive number", c.radius)
}

func circleArea(r float64) (float64, error) {
	if r < 0 {
		//return 0, errors.New("You give a wrong  number!")
		return 0, &areaError{"fuck you!", r}
	}
	return math.Pi * r * r, nil
}

func main() {
	fmt.Println("main")
	r1 := float64(-10.00)
	if area, err := circleArea(r1); err != nil {
		if err, _ := err.(*areaError); err != nil {
			fmt.Println(err.err, err.radius)
		} else {
			fmt.Println(err, r1)
			fmt.Println("The area of circle r1 is", area)
		}
	}
}
