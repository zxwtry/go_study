package b01_go_yu_yan_bian_cheng

/**
 * User:  zxwtry
 * Date:  2018/1/11
 * Time:  17:16
 */


/*
    错误处理：
        error接口
        Go语言引入一个关于错误处理的标准模式，即error接口
        接口定义如下：
        type error interface {
            Error() string
        }

    对于大多数函数，如果要返回错误，大致上可以定义为如下模式：
        将error作为多种返回值中的最后一个，但这并不是强制要求
        func Foo(param int) (n int, err error) {
            // ...
        }

    调用时的代码建议按如下方式处理错误情况：
        n, err := Foo(0)
        if err != nil {
            // 错误处理
        } else {
            // 使用返回值n
        }

    自定义的error类型：
        首先，定义一个用于承载错误信息的类型。
            Go语言接口非常灵活
            不需要从error接口继承或者像Java一样使用implements
            具体代码实现：
            type PathError struct {
                Op string
                Path string
                Err error
            }
            编译期怎么知道PathError课可以当一个error来传递？
            关键在于下面的代码实现了Error()方法：
            func (e *PathError) Error() string {
                return e.Op + " " + e.Path + ": " + e.Err.Error()
            }
            之后就可以直接返回PathError变量了。
            在下面的代码中，当syscall.Stat()失败返回err时，将该err包装到一个PathError对象中返回。
            func Stat(name string) (fi FileInfo, err error) {
                var stat syscall.Stat_t
                err = syscall.Stat(name, &stat)
                if err != nil {
                    return nil, &PathError("stat", name, err)
                } else {
                    return fileInfoFromStat(&stat, name), nil
                }
            }

            fi, err := os.Stat("a.txt")
            if err != nil {
                if e, ok := err.(*os.PathError); ok && e.Err != nil {
                    // 获取PathError类型变量e中的其他信息并处理
                }
            }


 */