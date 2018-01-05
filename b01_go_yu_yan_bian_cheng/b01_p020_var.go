package b01_go_yu_yan_bian_cheng

import "fmt"

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
    4， 复数类型：complex64、complex128
    5， 字符串：string
    6， 字符类型：rune
    7， 错误类型：error

    复合类型：
    1， 指针：pointer
    2， 数组：array
    3， 切片：slice
    4， 字典：map
    5， 通道：chan
    6， 结构体：struct
    7， 接口：interface
 */


 /*
    布尔类型：

    var v1 bool
    v1 = true
    v2 := (1 == 2)  // v2也会自动推导为bool类型

    var b bool
    b = 1   // 编译出错
    b = bool(1)   // 编译出错
    布尔类型不能接受其他类型的赋值
    布尔类型不支持自动或强制的类型转换

    b = (1 != 0)   // 编译正确
  */


/*
    整型：

    int8        1           -128 ~ 127
    uint8       1             0  ~ 255
    int16       2         -32768 ~ 32767
    uint16      2              0 ~ 65535
    int32       4          -2^31 ~ 2^31 - 1
    uint32      4              0 ~ 2^32 - 1
    int64       8          -2^63 ~ 2^63 - 1
    uint64      8              0 ~ 2^64 - 1
    int        平台相关
    uint       平台相关
    uintptr     32位平台4字节，64位平台8字节

    int和int32认为是两种不同的类型。
    编译器不会自动做类型转换
    var value int32
    value1 := 64        // value1自动推导为int类型
    value = value1      // 编译出错
    // 出错信息： cannot use value1 (tyype int) as type int32 in assignment

    可以使用强制类型转换：
    value = int32(value1)

    两个不同类型的整型数不能直接比较
    各种类型的整型变量可以直接与字面常量（literal）进行比较

    var i int32
    var j int64
    i, j = 1, 2
    if i == j {     // 编译错误

    }
    if i == 1 || j == 2 {       // 编译正确

    }


    取反：
    C语言： ~x
    GO：   ^x
    ^2  // 结果是-3
 */


/*
    Go语言浮点数：  float32和float64
    小数默认是 float64

    浮点数比较：
    import "math"
    // 用户自定义比较精度，比如0.0001
    func IsEqual(f1, f2, p float64) bool {
        return math.Fdim(f1, f2) < p
    }
 */


/*
    复数类型：

    var v1 complex64     // 由两个float32构成的复数类型
    v1 = 3.2 + 12i
    v2 := 3.2 + 12i     // v2是complex128类型
    v3 := complex(3.2, 12)    // v3是complex128类型

    实部：real(v1)
    虚部：imag(v1)
 */


/*
    字符串：
    var s1 string    // 声明一个字符串变量
    s1 = "Hello World"
    c0 := s1[0]     // 取字符串的第一个字符
    fmt.Printf("The length of \"%s\" is %d.\n", s1, len (s1))
    fmt.Printf("The first character of \"%s\" is %c.\n", s1, c0)

    常见字符串操作：
        x + y       字符串连接
        len(s)      字符串长度
        s[i]        取字符


 */


func P020String() {
    s := "Hello，世界"
    n := len(s)
    for i := 0; i < n; i ++ {
        ch := s[i]
        fmt.Println(i, ch)
    }
    fmt.Println("==============")
    for i, ch := range s {
        fmt.Println(i, ch)
    }
    /*
        0 72
        1 101
        2 108
        3 108
        4 111
        5 239
        6 188
        7 140
        8 228
        9 184
        10 150
        11 231
        12 149
        13 140
        ==============
        0 72
        1 101
        2 108
        3 108
        4 111
        5 65292
        8 19990
        11 30028

        （UTF-8中，一个中文占3个字节）
     */
}


/*
    字符类型：
        Go支持两种字符类型：byte（uint8的别名）和 rune（单个Unicode字符）
        rune（Unicode）比较少使用
 */


/*
    数组：
    声明方式：
        [32]byte        // 长度为32的数组，每个元素为一个字节
        [2 * N] struct { x, y int32}    // 复杂类型数组
        [1000]*float64                  // 指针数组
        [3][5]int                       // 二维数组
        [2][2][2]float64    // 等同于 [2][2]([2]float64)
        数组长度在定义后，不可更改
        数组长度可以为一个常量或者一个常量表达式
        （常量表达式：编译期即可计算结果的表达式）
        数组长度是该数组类型的一个内置常量，可通过len()获取

    元素访问：
        for i := 0; i < len(array); i ++ {
            fmt.Println("Element ", i , " of array is ", array[i])
        }
        for i, v := range array {
            fmt.Println("Element ", i , " of array is ", v)
        }

    值类型：
        数组是一个值类型（value type）。
        所有的值类型变量在赋值和作为参数传递时，都会产生一次复制动作
 */

func P020ValueType()  {
    var modify = func(array [5]int) {
        array[0] = 10
        fmt.Println("In modify(), array is: ", array);
    }
    array := [5]int{1, 2, 3, 4, 5}
    modify(array)
    fmt.Println("Out modify(), array is: ", array);
    /*
        In modify(), array is:  [10 2 3 4 5]
        Out modify(), array is:  [1 2 3 4 5]
     */
}


/*
    数组切片：
    数据结构：   1，指向原生数组的指针
                2，数组切片中元素个数
                3，数组切片已分配的存储空间

    创建数组切片：
            1， 基于数组
                var mySlice []int = myArray[:5]     下标 0~4
                var mySlice []int = myArray[:]      所有元素
                var mySlice []int = myArray[5:]     下标从5开始
            2， 直接创建

 */

func P020ArraySlice()  {
    // 先定义一个数组
    var myArray [10]int = [10]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    // 基于数组创建一个数组切片
    var mySlice []int = myArray[:5]

    fmt.Println("Elements of myArray : ")
    for _, v := range myArray {
        fmt.Print(v, " ")
    }
    fmt.Println()

    fmt.Println("Elements of mySlice : ")
    for _, v := range mySlice {
        fmt.Print(v, " ")
    }
    fmt.Println()

    /*
        Elements of myArray :
        0 1 2 3 4 5 6 7 8 9
        Elements of mySlice :
        0 1 2 3 4
     */
}


func P020()  {

}






