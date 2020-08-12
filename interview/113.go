
package main

import "fmt"

func link(p ...interface{}) {
    fmt.Println(p)
}

func main() {
    link("seek", 1, 2, 3, 4) // 输出 [seek 1 2 3 4]
    a := []interface{}{"seek", 1, 2, 3, 4} // interfaces{} 空指针 可以是任意类型
    link(a...) // 输出 [seek [1 2 3 4]]
}
