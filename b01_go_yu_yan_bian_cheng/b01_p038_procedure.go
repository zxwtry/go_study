package b01_go_yu_yan_bian_cheng

import (
    "fmt"
    "errors"
)

/**
 * User:  zxwtry
 * Date:  2018/1/8
 * Time:  17:15
 */


/*
    流程控制：
        if、else、else if、
        switch、case、select
        for、range、goto
        break、continue、fallthrough
 */


func P038Switch()  {
    i := 201
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    default:
        fmt.Println("others")
    }

    switch {
    case 0 <= i && i < 100:
        fmt.Println("0-100")
    case 200 <= i && i < 400:
        fmt.Println("200-400")
    case 600 <= i && i < 800:
        fmt.Println("600-800")
    }
}


func P038For()  {
    sum := 0

    for sum = 0; sum < 5; sum ++ {
        fmt.Println("sum is ", sum)
    }

    sum = 0
    for {
        sum ++
        if sum >= 5 {
            break
        }
        fmt.Println("sum is ", sum)
    }
    /*
        sum is  0
        sum is  1
        sum is  2
        sum is  3
        sum is  4
        sum is  1
        sum is  2
        sum is  3
        sum is  4
     */
}


func P038Goto()  {
    i := 0
    HERE:
        fmt.Println("i=", i)
        i ++
        for i < 10 {
            goto HERE
        }
    /*
        i= 0
        i= 1
        i= 2
        i= 3
        i= 4
        i= 5
        i= 6
        i= 7
        i= 8
        i= 9
     */
}


func P038Func()  {
    var add = func(a int, b int) (ret int, err error) {
        if a < 0 || b < 0 {
            err = errors.New("不能输入负数！")
            return
        }
        return a + b, nil
    }
    ret, err := add(1, 2)
    fmt.Println("answer is ", ret);
    fmt.Println("err is ", err)
    /*
        answer is  3
        err is  <nil>
     */

    ret, err = add(-1, 2)
    fmt.Println("answer is ", ret);
    fmt.Println("err is ", err)
    /*
        answer is  0
        err is  不能输入负数！
     */
}


/*
    函数调用：
    小写字母开头的函数，只在本包内可见
    大写字母开头的函数，才能被其他包使用
 */


/*
    不定参数：
    1， 不定参数类型
        func myfunc(args ...int) {
            for _, arg := range args {
                fmt.Println(arg)
            }
        }
 */