package b01_go_yu_yan_bian_cheng

import (
    "fmt"
)

/**
 * User:  zxwtry
 * Date:  2018/1/10
 * Time:  10:08
 */


/*
    匿名函数：
        func(a, b int, z float64) bool {
            return a * b < int(z)
        }
        匿名函数可以直接复制给一个变量或者直接执行
        f := func(x, y int) int {
            return x + y
        }
        func(ch chan int) {
            ch <- ACK
        } (reply_chan)      // 花括号后直接跟参数列表，表示函数调用
 */

func B045Anonymous()  {
    f1 := func() {
        fmt.Println("f1 called")
    }
    f1()

    func(ch int) {
        fmt.Println("anonymous called val: ", ch)
    } (2)

    /*
        f1 called
        anonymous called val:  2
     */
}


/*
    闭包：
    基本概念：
        闭包是可以包含自由（未绑定到特定对象）变量的代码块
        这些变量不在这个代码块内或者任何全局上下文中定义
        而是在定义代码块的环境中定义
        要执行的代码块为自由变量提供绑定的计算环境（作用域）

    闭包的价值：
        可以作为函数对象或者匿名函数
        支持闭包的多数语言都将函数作为第一级对象
        这些函数可以存储到变量中作为参数传递给其他函数
        能够被函数动态创建和返回

    Go语言中的闭包：
        同样也会引用到函数外的变量。
        只要闭包还被使用，那么被闭包引用的变量会一直存在
 */


func P045Closure()  {
    var j int = 5
    a := func() (func()) {
        var i int = 10
        return func() {
            fmt.Printf("i, j: %d, %d\n", i, j)
        }
    }()
    a()
    j *= 2
    a()
    /*
        i, j: 10, 5
        i, j: 10, 10
     */
}


