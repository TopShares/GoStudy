//example5.go
/*
使用原子函数
对整形变量或指针的同步访问，可以使用原子函数来进行。
这里使用原子函数来修复example4.go中的竞争状态问题
*/
package main

import (
    "sync"
    "runtime"
    "fmt"
    "sync/atomic"
)

var (
    //counter为访问的资源
    counter int64
    wg      sync.WaitGroup
)

func addCount() {
    defer wg.Done()
    for count := 0; count < 2; count++ {
        //使用原子操作来进行
        atomic.AddInt64(&counter,1)
        //当前goroutine从线程退出
        runtime.Gosched()
    }
}

func main() {
    wg.Add(2)
    go addCount()
    go addCount()
    wg.Wait()
    fmt.Printf("counter: %d\n",counter)
}
//output：
// counter: 4