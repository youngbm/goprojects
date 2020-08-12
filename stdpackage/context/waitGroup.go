
package main

import "sync"
import "fmt"
import "time"


func main(){

    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        time.Sleep(time.Second)
        fmt.Println("FUNC1")
        wg.Done()
    }()

    go func() {
        fmt.Println("FUNC2")
        time.Sleep(2 * time.Second)
        wg.Done()
    }()

    wg.Wait()
    fmt.Println("End!!!")
}
