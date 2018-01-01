package b01_go_yu_yan_bian_cheng

/**
 * User:  zxwtry
 * Date:  2017/12/30
 * Time:  15:05
 */

/*
    变量声明：
        var v1 int
        var v2 string
        var v3 [10]int
        var v4 []int
        var v5 struct {
            f int
        }
        var v6 *int     // 指针
        var v7 map[string] int      // map key为string, value为int
        var v8 func(a int) int
        var (
            v1 int
            v2 string
        )

    变量初始化：
        var v1 int = 10
        var v2 = 10
        v3 := 10
    (:=  进行变量声明和初始化工作)

    初始化编译出错：
        var i int
        i := 2

    多重赋值：
        (交换i和j)
        i,j = j,i

    匿名变量：
        func GetName() (firstName, lastName, nickName string) {
            return "May", "Chan", "Chibi Maruko"
        }
        _, _, nickName := GetName()
        只想获得nickName

    常量定义：
        const PI float64 = 3.14
        const zero = 0.0
        const (
            size int64 = 1024
            eof = -1
        )
        const u, v float32 = 0, 3       // u=0.0 v=3.0常量的多重赋值
        const a, b, c = 3, 4, "foo"
        // a = 3, b = 4, c = "foo"

        Go的常量定义可以限定常量类型，但不是必需的。
        如果定义常量没有指定类型，和字面常量一样，也是无类型常量

        常量定义的右值，也可以是一个在编译期运算的表达式
        const mask = 1 << 3
        由于常量的赋值是一个编译期行为，
        右值不能出现任何需要运行期才能得出结果的表达式。
        const Home = os.GetEnv("HOME")      //这里会有编译错误
        os.GetEnv()只有在运行期才能知道返回结果，编译期不能确定

 */

/*
    预定义常量：
        Go语言预定义了这些常量：true、false、iota
        iota比较特殊，可以被认为是一个可被编译器修改的变量
        在每一个const关键字出现时被重置为0
        然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增1
        const (             // iota被重设为0
            c0 = iota       // c0 == 0
            c1 = iota       // c1 == 1
            c2 = iota       // c2 == 2
            c3 = iota       // c3 == 3
        )
        const (	                // iota被重设为0
            c0 = 1 << iota      // c0 == 1
            c1 = 1 << iota      // c1 == 2
            c2 = 1 << iota      // c2 == 4
            c3 = 1 << iota      // c3 == 8
        )
        const (
            u 	      = iota * 42       // u == 0
            v float64 = iota * 42       // v == 42.0
            v float64 = iota * 42       // v == 84.0
        )
        const x = iota          // x == 0 (const出现，iota置0)
        const y = iota          // y == 0 (const出现，iota置0)

        如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式。
        const (                 // iota被重设为0
            c0 = iota           // c0 = 0
            c1                  // c1 = 1
            c2                  // c2 = 2
        )
        const (                      // iota被重设为0
            c0 = 1 << iota           // c0 = 1
            c1                       // c1 = 2
            c2                       // c2 = 4
        )
 */


/*
    枚举：
        const (
            Sunday = iota
            Monday
            Tuesday
            Wednesday
            Thursday
            Friday
            Saturday
            numberOfDays
        )
 */


/*
    类型：
    Go语言内置一下这些基础类型：
    1， 布尔类型：bool
    2， 整型：int8、byte、int16、int、uint、uintptr等
    3， 浮点类型：float32、float64
    4，
 */