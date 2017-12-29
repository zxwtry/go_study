package main

import "os"
import "fmt"
import "simplemath"
import "strconv"

/*
    在Windows中，使用calc.exe add 2 3时候
    会出现参数比linux多一个
        args[0] is  calc.exe
        args[1] is  add

    Windows下编译过程：
    set GOPATH={calcproj所在目录}
    mkdir bin
    cd bin
    go build calc
    之后就会在bin目录下生成calc.exe文件

    calc.exe add 2 3
    Result:  5
    calc.exe sqrt 8
    Result:  2

    测试：
    go test simplemath
    ok      simplemath      0.028s

    Linux下没有测试，应该也类似吧，set 换成 export
*/


var Usage = func() {
    fmt.Println("USAGE: calc command [arguments] ...")
    fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

var argsOff = 1;        // Windows 1;   Linux 0

func main() {
    args := os.Args
    if args == nil || len(args) < 2 {
        Usage()
        return
    }
    switch args[0 + argsOff] {
        case "add":
            if len(args) != (3 + argsOff) {
                fmt.Println("USAGE: calc add <integer1><integer2>")
                return
            }
            v1, err1 := strconv.Atoi(args[1 + argsOff])
            v2, err2 := strconv.Atoi(args[2 + argsOff])
            if err1 != nil || err2 != nil {
                fmt.Println("USAGE: calc add <integer1><integer2>")
                return
            }
            ret := simplemath.Add(v1, v2)
            fmt.Println("Result: ", ret)
        case "sqrt":
            if len(args) != (2 + argsOff) {
                fmt.Println("USAGE: calc add <integer>")
                return
            }
            v, err := strconv.Atoi(args[1 + argsOff])
            if err != nil {
                fmt.Println("USAGE: calc sqrt <integer>")
                return
            }
            ret := simplemath.Sqrt(v)
            fmt.Println("Result: ", ret)
        default:
            Usage()
    }
}