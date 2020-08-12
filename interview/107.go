package main

import "fmt"
import "sync"

func main(){
    var m sync.Map
    m.LoadOrStore("a", "aa")

    m.Delete("a")
    fmt.Println(m)
    //{{0 0} {{map[] true}} map[] 0}
    //{{0 0} {{map[] true}} map[a:0xc00000e030] 0}

    // Next
    var wg sync.WaitGroup
    var m1 sync.Mutex
    fmt.Println(m1)
    N := 1000
    wg.Add(N)
    var ints = make([]int, 0, 1000)

    for i := 0; i < N; i++ {
        go func(){
            for i := 0; i < 10000; i++ {
                m1.Lock()
                ints = append(ints, i)  // append 不是线程安全,请上锁, 有并发就要考虑上锁
                m1.Unlock()
            }
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(len(ints))
}
