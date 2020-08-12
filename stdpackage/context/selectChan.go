
package main

import "fmt"
import "time"


// 通过channel 来控制进程的退出
func main(){
    var ch1 chan int = make(chan int)
    var ch2 chan int = make(chan int)

    go func(){
        for {
            select {
                case <-ch1: // 等待退出
                    fmt.Println("get value, F1 exit!")
                default: // 函数处理逻辑在这里
                    fmt.Println("F1 still running!")
                    time.Sleep(time.Second)
            }
        }
    }()

    go func(){
        for {
            select {
                case <-ch2: // 等待退出
                    fmt.Println("get value, F2 exit!")
                default: // 函数处理逻辑在这里
                    fmt.Println("F2 still running!")
                    time.Sleep(time.Second)
            }
        }
    }()

    fmt.Println("Main Process")
    time.Sleep(time.Second * 10 )
    ch1 <- 1 //
    ch2 <- 1
    time.Sleep(time.Second * 1 )
    fmt.Println("End Process")
}
