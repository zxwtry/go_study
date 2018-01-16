package b01_go_yu_yan_bian_cheng

import (
    "fmt"
    "sync"
    "runtime"
)

/**
 * User:  zxwtry
 * Date:  2018/1/16
 * Time:  10:02
 */


/*
    并发意味着：程序运行时，有多个执行的上下文，对应多个调用栈

    并发价值：
    1，界面响应和运算密集/IO密集
    2，Web服务器处理大量用户请求
    3，分布式环境，不同计算机上执行同一个任务小部分
    4，CPU从单核到多核，硬件能力没有发挥
    5，IO阻塞

    并发优势：
    1，问题模型，很多都是多条线的
    2，CPU多核优势
    3，CPU与其他硬件设备的异步

    并发的实现模型：
    1，多进程：操作系统层面进行的并发。开销最大的并发。
        好处：简单、进程间互不影响
        坏处：系统开销大、所有进程由内核管理
    2，多线程：系统层面的并发，使用最多最有效模式。
        好处：比多进程开销小
        坏处：开销还是比较大、高并发下效率有影响
    3，基于回调的非阻塞/异步IO：
        来源于：多线程模式的危机。
                （多线程模式会很快耗尽内存和CPU）
        通过事件驱动方式使用异步IO，可以少用线程、降低开销
        在Node.js中得到很好的实践。
        编程比多线程复杂很多。
    4，协程：(Coroutine)本质上是用户态线程
        不需要操作系统来进行抢占式调度
        且在真正的实现中，寄存于线程
        系统开销极小
        有效提高线程的任务并发
        避免多线程的缺点
 */


/*
    共享内存系统   -->  消息传递系统
    对线程间共享状态的各种操作都被封装在线程之间传递的消息中。
    要求是：发送消息时，对状态进行复制
            在消息传递的边界上，交出状态的所有权
            从逻辑上看，这个操作与共享内存系统中执行的原子更新操作相同。
            从物理上看，非常不同。
            由于需要执行复制操作，所以大多数消息传递的实现
                在性能上，并不优越。但在线程中的状态管理工作通常变得简单。
 */


/*
    协程：
    执行体是个抽象概念。
    在操作系统层面，有多个概念与之对应。
    操作系统自己掌管的进程process、进程内的线程thread、进程内的协程coroutine
    与传统的系统级线程和进程相比，协程的最大优势在于“轻量级”
    能轻松创建百万个，不会导致系统资源枯竭。
    线程和进程最多不能超过1万个

    协程：轻量级线程的切换管理，不依赖于系统的线程和进程，也不依赖于CPU核数
 */


/*
    goroutine是Go语言中的轻量级线程实现。
    由Go运行时(runtime)管理。
    func Add(x, y int) {
        z := x + y
        fmt.Println(z)
    }
    如何让这个函数并发执行呢？
    go Add(1, 1)

    在一个函数调用前加上go关键字，这次调用就会在一个新的goroutine
        中并发执行。
    当被调用的函数返回时，这个goroutine也自动结束了。
    需要注意，如果这个函数有返回值，那么这个返回值会被丢弃。
 */


func Add(x, y int) {
    z := x + y
    fmt.Println(z)
}


func P088Add() {
    for i := 0; i < 10; i ++ {
        go Add(i, i)
    }
    /*
        执行结果：任何输出都没有
        原因是：
            主函数启动了10个goroutine
            然后程序就退出了
            另外10个goroutine还没有来得及执行
     */
}


/*
    并发通信：

 */


/*
    C语言中，线程间数据共享

 */


var counter int = 0


func Count(lock *sync.Mutex)  {
    lock.Lock()
    counter ++;
    fmt.Println(counter)
    lock.Unlock()
}


func P088Count() {
    lock := &sync.Mutex{}

    for i := 0; i < 10; i ++ {
        go Count(lock)
    }

    for {
        lock.Lock()
        c := counter
        lock.Unlock()
        runtime.Gosched()
        if c >= 10 {
            break;
        }
    }
}


/*
    锁
    共享变量
    错误分支
    让代码难以维护

    Go语言提供另一种通信模型：
        以消息机制，而非共享内存，作为通信方式。

    不通过共享内存来通信，
    通过通信来共享内存
 */


/*
    channel
    channel是Go语言在语言级别提供的goroutine间的通信方式。
    可以使用channel在两个或多个goroutine之间传递消息。
    channel是进程内的通信方式。
    通过channel传递对象的过程和调用函数时的传递行为比较一致。
    比如：可以传递指针。

    如果需要跨进程通信，可以用分布式系统的方法来解决。
    比如使用Socket或HTTP等通信协议。
    Go语言对于网络方面也有完善的支持。

    channel是类型相关的。
    一个channel只能传递一种类型的值。
    这个类型需要在声明channel时指定
    如果对Unix管道有了解，不难理解channel，可以认为是一种类型安全的管道。

 */

func CountChannel(ch chan int)  {
    ch <- 1
    fmt.Println("Counting")
}


func P088Channel() {
    chs := make([] chan int, 10)
    for i := 0; i < 10; i ++ {
        chs[i] = make(chan int)
        go CountChannel(chs[i])
    }
    for _, ch := range chs {
        <-ch
    }
    /*
        这个例子中，定义了一个包含10个
        channel的数组 chs
        并把数组中每个channel分配给10个不同的goroutine
        每个goroutine的Add()函数完成后，
        通过ch <- 1语句，向对应channel写入一个数据。
        在这个channel被读取前，这个操作是阻塞的。

        在所有goroutine启动完成后，
        通过<-ch语句从10个channel中一次读取数据。

        在对应的channel写入数据前，这个操作也是阻塞的。

        这样，使用channel实现了类似锁的功能，
        进而保证了所有goroutine完成后主函数才返回。

        使用Go语言开发时，经常遇到需要实现条件等待的场景。
        这也是channel可以发挥作用的地方。
     */
}


/*
    channel的声明形式：
        var charName chan ElementType
        与一般变量声明不同的地方仅仅在类型之前
            加了chan关键字

        ElementType指定这个channel所能传递的元素类型。

        例子：需要声明一个传递类型为int的channel
            var ch chan int

        或者，声明一个map，元素是bool型的channel
            var m map[string] chan bool

        定义一个channel，直接使用内置的函数make
            ch := make(chan int)

        这就声明并初始化了一个int型的名为ch的channel

        在channel的用法中，最常见的包括写入和读出
        将一个数据写入（发送）至channel的语法：
            ch <- value
        向channel写入数据  --> 程序阻塞
        直到有其他goroutine从这个channel中读取数据

        从channel中读取数据的语法：
            value := <-ch

        如果channel之前没有写入数据，那么channel中读取数据也会导致程序阻塞。
        知道channel中被写入数据为止。
 */


/*
    select

 */