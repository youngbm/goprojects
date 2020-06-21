package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("main")

	//files, _ := filepath.Glob("[]")   this means  skips the err  return
	files, error := filepath.Glob("*go")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println(len(files), "files found!", "matched files:")
	for _, v := range files {
		fmt.Println(v)

	}
}
