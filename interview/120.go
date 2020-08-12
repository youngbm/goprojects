
package main

import "fmt"


func GetValue(m map[int]string, id int)(string, bool){
    if _, exist := m[id]; exist {
        return "存在数据", true
    }
    //return nil, false // 返回string, bool, 这里却返回nil，改为“”
    return "", false 
}

func main(){
    fmt.Println(len("你好bj!")) // 汉字 utf8 3个字节

    // Next
    intmap := map[int]string{
        1: "aa",
        2: "bb",
        3: "cc",
    }

    v, err := GetValue(intmap, 3)
    fmt.Println(v, err)
}

