package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	sync.Mutex //[Import]锁失效。将 Mutex 作为匿名字段时 相关的方法必须使用指针接收者
	//*sync.Mutex  //方法2
}

//func (d data) test(s string) {   //  接收者改为指针
func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add()
	var d data

	go func() {
		defer wg.Done()
		d.test("Read")
	}()

	go func() {
		defer wg.Done()
		d.test("Write")
	}()

	wg.Wait()
	fmt.Println("End")

}
