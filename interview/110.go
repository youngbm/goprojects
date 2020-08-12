
package main

import (
    "fmt"
    "sync"
    "time"
    "runtime"
)

type T struct {
    V int
}

func (t *T) Incr(wg *sync.WaitGroup){
    t.V++
    wg.Done()
}

func (t *T) Print(){
    time.Sleep(1)
    fmt.Print(t.V)
}

func main(){
    var wg sync.WaitGroup
    wg.Add(10)
    var ts = make([]T, 10)
    for i := 0; i <10; i++ {
        ts[i] = T{i}
    }

    //for i, t := range ts { // 这里的ts是副本，修改无效
    for i := 0; i <10; i++ {
        fmt.Println(i)
        go ts[i].Incr(&wg)
    }

    wg.Wait()
    fmt.Println(ts)
    for _, t := range ts {
        t.Print()
    }
    time.Sleep(1 * time.Second)

    // Next
    fmt.Println("\n\n")
    const N = 26
    const GOMAXPROCS = 1
    runtime.GOMAXPROCS(GOMAXPROCS) // 限定用一个逻辑CPU去跑
    var wg1 sync.WaitGroup
    wg1.Add(3 * N)
    for i := 0; i < N; i++ {
        go func(i int) {
            defer wg1.Done()
            runtime.Gosched()  // 把CPU时间分出去
            fmt.Printf("%c", 'a'+i)
        }(i)
        go func(i int) {
            defer wg1.Done()
            fmt.Printf("%c", 'A'+i)
        }(i)

        go func(i int) {
            defer wg1.Done()
            runtime.Gosched()  // 把CPU时间分出去
            fmt.Printf("%d", i)
        }(i)
    }
    wg1.Wait()
}
