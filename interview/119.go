package main


func test123() (a int) {
    if true {
        a := 123 // 这里是短= 那么作用域只是局限于 if
        println(a)
    }
    return a
}

func main(){
    println(test123())
}
