
package main

import  (
    "fmt"
    "time"
    "context"
    )


func main(){

    ctx1, cancel1 := context.WithCancel(context.Background())
    ctx2, cancel2 := context.WithCancel(context.Background())
    ctx3, cancel3 := context.WithCancel(context.Background())

    go worker(ctx1, "w1")
    go worker(ctx2, "w2")
    go worker(ctx3, "w3")

    time.Sleep(2 * time.Second)
    cancel1()
    time.Sleep(1 * time.Second)
    cancel2()
    time.Sleep(1 * time.Second)
    cancel3()
    fmt.Println("End main process")
}

func worker(ctx context.Context, s string){
    for {
        select {
        case <-ctx.Done():
            fmt.Println(s+" END")
            return
        default:
            fmt.Println(s + " Still running")
            time.Sleep(time.Second)
        }
    }
}
