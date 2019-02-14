//example6.go
/*
使用互斥锁
对临界区的访问，可以使用互斥锁来进行。
对于example4.go的竞争状态，可以使用互斥锁来解决：
*/
package example

import (
   "sync"
   "runtime"
   "fmt"
)

var (
   //counter为访问的资源
   counter int
   wg      sync.WaitGroup
   mutex   sync.Mutex
)

func addCount() {
   defer wg.Done()

   for count := 0; count < 2; count++ {
      //加上锁，进入临界区域
      mutex.Lock()
      {
         value := counter
         //当前goroutine从线程退出
         runtime.Gosched()
         value++
         counter = value
      }
      //离开临界区，释放互斥锁
      mutex.Unlock()
   }
}
func main() {
   wg.Add(2)
   go addCount()
   go addCount()
   wg.Wait()
   fmt.Printf("counter: %d\n", counter)
}

//output:
counter: 4