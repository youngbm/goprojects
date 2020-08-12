
package main


import "fmt"


var y int

func f(x int)int {
    return 7
}

func t(){
    y = 7
    switch y { // 这个是可以是变量，可以是空, 注意这里有没有分号(;)
        case 7: // 变量这里就是值， 空则是表达式
            fmt.Println("rt 7")
            return
    }
}

func main(){
    t()
}
