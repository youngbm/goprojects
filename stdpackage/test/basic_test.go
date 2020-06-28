package mystrings

import (
	"fmt"
	"testing"
)

type usCase struct {
	in, out string
}

var a = []usCase{
	usCase{"sdf", "SDF"},
	usCase{"abcd", "ABCD"},
	usCase{"123", "123"},
	usCase{"abc-123", "ABC-123"},
}

func TestUpperCase(t *testing.T) {
	for _, v := range a {
		fmt.Println(v)
		if v.out == UpperCase(v.in) {
			fmt.Println("This case pass")
		} else {
			t.Errorf("uppercase(%s), must be %s", v.in, v.out)
		}
	}

}
