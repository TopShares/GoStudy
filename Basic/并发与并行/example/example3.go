//example3.go
package example

import (
   "sync"
   "runtime"
   "fmt"
)

var (
   //counter为访问的资源
   counter int64
   wg      sync.WaitGroup
)

func addCount() {
   defer wg.Done()
   for count := 0; count < 2; count++ {
      value := counter
      //当前goroutine从线程退出
      runtime.Gosched()
      value++
      counter=value
   }
}

func main() {
   wg.Add(2)
   go addCount()
   go addCount()
   wg.Wait()
   fmt.Printf("counter: %d\n",counter)
}
//output:
//counter: 4 或者counter: 2