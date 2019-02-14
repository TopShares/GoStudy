//example6.go

/*
利用通道共享数据
Go语言采用CSP消息传递模型。通过在goroutine之间传递数据来传递消息，
而不是对数据进行加锁来实现同步访问。这里就需要用到通道chan这种特殊的数据类型。
当一个资源需要在goroutine中共享时，chan在goroutine中间架起了一个通道。
通道使用make来创建：

unbuffered := make(char int) //创建无缓存通道，用于int类型数据共享
buffered := make(chan string,10)//创建有缓存通道，用于string类型数据共享
buffered<- "hello world" //向通道中写入数据
value:= <-buffered //从通道buffered中接受数据

通道用于放置某一种类型的数据。创建通道时指定通道的大小，将创建有缓存的通道。
无缓存通道是一种同步通信机制，它要求发送goroutine和接收goroutine都应该准备好，
否则会进入阻塞。
*/


/*
无缓存的通道
无缓存通道是同步的——一个goroutine向channel写入消息的操作会一直阻塞，
直到另一个goroutine从通道中读取消息。
反过来也是，一个goroutine从channel读取消息的操作会一直阻塞，
直到另一个goroutine向通道中写入消息。
《Go in action》中关于无缓存通道的解释有一个非常棒的例子：网球比赛。
在网球比赛中，两位选手总是处在以下两种状态之一：要么在等待接球，要么在把球打向对方。
球的传递可看为通道中数据传递。下面这段代码使用通道模拟了这个过程：
*/
package main

import (
    "sync"
    "fmt"
    "math/rand"
    "time"
)

var wg sync.WaitGroup

func player(name string, court chan int) {
    defer wg.Done()
    for {
        //如果通道关闭,那么选手胜利
        ball, ok := <-court
        if !ok {
            fmt.Printf("Player %s Won\n", name)
            return
        }
        n := rand.Intn(100)

        //随机概率使某个选手Miss
        if n%13 == 0 {
            fmt.Printf("Player %s Missed\n", name)
            //关闭通道
            close(court)
            return
        }
        fmt.Printf("Player %s Hit %d\n", name, ball)
        ball++
        //否则选手进行击球
        court <- ball
    }
}


func main() {
    rand.Seed(time.Now().Unix())
    court := make(chan int)
    //等待两个goroutine都执行完
    wg.Add(2)
    //选手1等待接球
    go player("candy", court)
    //选手2等待接球
    go player("luffic", court)
    //球进入球场（可以开始比赛了）
    court <- 1
    wg.Wait()
}
//output:
// Player luffic Hit 1
// Player candy Hit 2
// Player luffic Hit 3
// Player candy Hit 4
// Player luffic Hit 5
// Player candy Missed
// Player luffic Won