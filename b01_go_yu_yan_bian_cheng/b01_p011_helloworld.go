package b01_go_yu_yan_bian_cheng

import "fmt"

/**
 * User:  zxwtry
 * Date:  2017/12/29
 * Time:  11:25
 */

/*
    1， 包是Go语言中最基本的分发单位
    2，	Go程序入口：名为main的包，名为main()的函数
    3，	main()函数不能带参数，不能定义返回值
        命令行传入的参数在os.Args变量
    4，	函数实例：
        func Compute(value1 int, value2 float64)(result float64, err error) {
            // 函数体
        }
    5，	不是所有的返回值都必须复制，没有赋值的返回值会设置为默认值
        result被设为0.0，err被设为nil
 */

func P011Hello() {
    fmt.Println("Hello World!")
}


/*
    1， 验证Go被正确安装
    2， go run hello.go
        直接运行，将编译、链接和运行合为一步
    3，	go build hello.go
        ./hello
    4， 6g hello.go
        6l hello.6
        ./6.out
        32位使用 8g 8l
 */