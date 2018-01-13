package b01_go_yu_yan_bian_cheng

/**
 * User:  zxwtry
 * Date:  2018/1/12
 * Time:  9:38
 */


/*
    C++ 文件操作

    class file_closer {
        FILE _f;
    public:
        file_closer(FILE f) : _f(f) {}
        ~file_closer() {if (f) fclose(f);}
    }

    使用：
    void f() {
        FILE f = open_file("file.txt");
        f_closer _closer(f);
        // 对f句柄进行操作
    }

    为什么需要file_closer这个包装类？
        如果没有这个类，代码中所有退出函数的环节
        都可能抛出异常，每一个return之前都需要关闭之前打开的文件句柄


    C++中还有另外一种解决方案：
        所有需要释放的资源变量都声明在函数的开头部分
        并在函数的末尾部分同意释放资源。
        函数需要退出时，必须goto到函数末尾，先完成资源清理工作
 */


/*
    Go 文件操作
    使用defer
    func CopyFile(dst, src string) (w int64, err error) {
        srcFile, err := os.Open(src)
        if err != nil {
            return
        }

        defer srcFile.Close()

        dstFile, err := os.Create(dstName)
        if err != nil {
            return
        }
        defer dstFile.Close()

        return io.Copy(dstFile, srcFile)

        // Copy()函数抛出异常，Go仍然会保证dstFile和srcFile被正常关闭
    }

    如果一条语句无法完成清理工作：
        defer func() {
            // 执行复杂的清理工作
        } ()
    一个函数中可以存在多个defer语句，
    defer语句的调用顺序是遵照后进先出的原则
    最后一个defer语句最先被执行
 */


/*
    Go语言引入两个内置函数panic()和recover()
    用于报告和处理运行时错误和程序中错误场景
    func panic(interface{})
    func recover() interface{}
    当一个函数执行过程中调用panic()函数时，
    正常的函数执行流程将立即终止，
    但函数中之前使用defer关键字延迟执行的语句将正常展开执行，
    之后该函数将返回到调用函数，并逐层向上执行panic流程，
    直到所属的goroutine中所有正在执行的函数被终止。
    错误信息将被报告，包括在调用panic()函数时传入的参数，
    这个过程被称为错误处理流程。

    从panic()的参数类型interface{}可以得知，该函数可以接收任何类型的数据。
    比如整型、字符串、对象等等。
    panic(404)
    panic("nextwork broken")
    panic(Error("file not exists"))

    recover()函数用于种植错误处理流程。
    在一般情况下，recover(0应该在一个使用defer关键字的函数中执行
    以有效截取错误处理流程。如果没有在反生异常的goroutine中
    明确调用恢复过程（recover关键字），会导致该goroutine所属的进程
    打印异常信息后直接退出。

    常见场景：
        对于foo()函数的执行，可能触发错误处理。
        可以用如下方式在调用代码中截取recover()
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Runtime error caught: %v", r)
            }
        }()
        foo()

        无论foo()中是否出发了错误处理流程， 该匿名defer函数都将
        在函数退出时得到执行。
        假如foo()中触发了错误处理流程，recover()函数执行将
        使得该函数处理过程终止。
        如果错误处理流程被触发时，程序传给panic函数的参数不为nil
        则该函数还会打印详细的错误信息。
 */