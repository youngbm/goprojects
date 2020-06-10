
package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string{
    if x < 0 {
        return sqrt(-x) +"i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main(){
    x := 0
    if x > 10 {
        fmt.Println(x)
    }else{
        fmt.Println("Under 10")
    }
    fmt.Println(sqrt(-10))
}
