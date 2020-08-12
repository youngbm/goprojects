
package main


import "fmt"
import "context"
import "time"

func main(){
    ctx, cancel := context.WithCancel(context.Background())

    go func() {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("got stop channel")
                return
            default:
                fmt.Println("still working")
                time.Sleep(time.Second)
            }
        }
    }()

    time.Sleep(3 * time.Second)
    fmt.Println("stop the goroutine")
    cancel()
    time.Sleep(time.Second)
    fmt.Println("End process")
}

