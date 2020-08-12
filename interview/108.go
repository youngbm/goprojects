package main

import "fmt"
import "encoding/json"
import "flag"

type People struct {
    Name string `json: "name"`     // json 解析必须要大写
    //name string `json: "name"` //小写为私有属性不会别打印
}

var ip string
var port int

func init(){
    //           变量, 参数,  默认值, 描述
    flag.IntVar(&port, "port", 8000, "port")
    flag.StringVar(&ip, "ip", "localhost", "port")
}

func main(){
    js := `{
        "name": "xiezhiyang"
    }`

    var p People
    err := json.Unmarshal([]byte(js), &p)
    if err != nil {
        fmt.Println("err:", err)
        return
    }
    fmt.Println("people", p)

    // Next
    flag.Parse() // 自动把参数解析并赋值给变量
    // go run 108.go --ip 127.0.0.1 --port 8080
    fmt.Println(ip, port)
}
