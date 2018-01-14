package b01_go_yu_yan_bian_cheng

import "fmt"

/**
 * User:  zxwtry
 * Date:  2018/1/12
 * Time:  18:04
 */


/*
    Go语言并没有沿袭传统面向对象编程中的
        继承、虚函数、构造函数、析构函数、this指针等
    Go语言优雅之处：
        对面向对象编程的支持是语言类型系统中的天然组成成分
        整个类型系统通过接口串联，浑然一体

    类型系统：
        1，基础类型：byte、int、bool、float等
        2，复合类型：数组、结构体、指针等
        3，可以指向任意对象的类型（Any类型）
        4，值语义和引用语义
        5，面向对象，即所有具备面向对象特征（比如成员方法）的类型
        6，接口

    Go语言
        大多数类型都是值语义，并且都一个包含对应的操作方法
        在需要的时候，可以给任何类型（包括内置类型）“增加”新方法
        在实现某个接口时，无需从该接口继承
        （Go语言不支持面向对象思想中的继承语法）
        只需要实现该几口要求的所有方法即可。
        任何类型都可以被Any类型引用。
        Any类型就是空接口，即interface{}

 */


/*
    为类型添加方法
    在Go语言中，可以给任意类型（包括内置类型、不包括指针类型）
    添加相应的方法

    type Integer int
    func (a Integer) Less(b Integer) bool {
        return a < b
    }
 */

type Integer int


func (a Integer) less(b Integer) bool  {
    return a < b
}


func P062AddMethod()  {
    var a Integer = 1
    var b Integer = 2
    if a.less(b) {
        fmt.Println("a.less(b) is true")
    } else {
        fmt.Println("a.less(b) is false")
    }
    /*
        a.less(b) is true
     */
}


/*
    在Go语言中没有隐藏的this指针
    方法施加的目标(也就是“对象”)显式传递，没有被隐藏起来
    方法施加的目标(也就是“对象”)不需要非得是指针，也不用非地叫this

    Go语言和C语言一样，类型是基于值传递的。
    如果想要修改变量的值，只能传递指针。

    只有需要修改对象的时候，才必须用指针

    $GOROOT/src/pkg/http/header.go
    HTTP头部信息的Header类型
    就是通过Go内置的map类型赋予新的语义来实现的。

    // Header类型用于表达HTTP头部的键值对信息
    type Header map[string][] string
    // Add()方法用于添加一个键值对到HTTP头部
    // 如果该键已存在，则会将值添加到已存在的值后面
    func (h Header) Add(key, value string) {
        textproto.MIMEHeader(h).Add(key, value)
    }

    // Set()方法用于设置某个键对应的值，如果存在，则替换
    func (h Header) Set(key, value string) {
        textproto.MIMEHeader(h).Set(key, value)
    }
 */

func (a *Integer) add(b Integer)  {
    *a += b
}


func P062PassAddress()  {
    var a Integer = 1
    var b Integer = 2
    a.add(b)
    fmt.Println("a", a)
    /*
        a 3
     */
}


/*
    值语义和引用语义
    值语义和引用语义的差在于赋值
        b = a
        b.Modify()
    如果b的修改不会影响a的值，那么此类型属于值类型
    如果b的修改会影响a的值，那么此类型属于引用类型

    Go语言中的大多数类型都属于值语义，包括：
    1， 基本类型：如byte、int、bool、float32、float64、string等
    2， 复合类型：如数组（array）、结构体（struct）、指针（pointer）等

    Go语言中类型的值语义表现得非常彻底，可以看数组
    （C语言中的数组：通过函数传递一个数组的时候，基于引用语义，）
    （但是在结构体中定义数组变量的时候，基于值语义）
    （表现在为结构体赋值的时候，该数组会被完整地复制）

    Go语言中的数组和基本类型没有区别，是很纯粹的值类型
    var a = [3]int{1, 2, 3}
    var b = a
    b[1] ++
    fmt.Println(a, b)
    // 运行结果： [1, 2, 3] [1, 3, 3]
    表示b = a是将数组内容完整复制。

    表达引用，使用指针：
    var a = [3]int{1, 2, 3}
    var b = &a
    b[1] ++
    fmt.Println(a, b)
    // 运行结果：[1, 3, 3] [1, 3, 3]
 */

func P062ByVauleByReference()  {
    var a = [3]int{1, 2, 3}
    var b = a
    var c = &a
    b[1] ++
    fmt.Println(a, b)
    c[1] ++
    fmt.Println(a, c)
    fmt.Println(a, *c);
    /*
        [1 2 3] [1 3 3]
        [1 3 3] &[1 3 3]
        [1 3 3] [1 3 3]
     */

    /*
        c的类型不是 [3]int
              而是 *[3]int
     */
}


/*
    Go语言中有4个类型比较特别，看起来像引用类型
    1，数组切片：指向数组(array)的一个区间
    2，map
    3，channel：执行体(goroutine)间的通信设施
    4，interface：对一组满足某个契约的类型的抽象

    数组切片本质上是一个区间，可以大致将[]T表示为：
    type slice struct {
        first *T
        len int
        cap int
    }
    数组切片类型本身的赋值仍然是值语义

    type MAP_K_V struct {
        // ...
    }
    type map[K][V] struct {
        impl *MAP_K_V
    }
    基于指针，完全可以自定义一个引用类型
    type IntegerRef struct {
        impl *int
    }

    channel 和 map类似，本质上是一个指针
    将他们设计为引用类型，而不是统一的值类型的原因：
        完整复制一个channel或map并不是常规需求

    同样，接口具备引用语义，是因为内部维持了两个指针：
    type interface struct {
        data *void
        itab *Itab
    }
    接口在Go语言中的地位非常重要。
 */


/*
    结构体：
    Go语言中的结构体(struct)和其他语言的类(class)具有相同地位。
    Go语言放弃了包括继承在内的大量面向对象特性，
        只保留组合（composition）这个基础的特性
        组合只是复合类型的基础
    所有的Go语言类型（指针类型除外）都可以有自己的方法
    Go语言的结构体只是很普通的复合类型，平淡无奇。
    比如定义一个矩阵类型：
        type Rec struct {
            x, y float64
            width, height float64
        }
    定义成员方法Area()来计算矩形的面积
        func (r *Rect) Area() float64 {
            return r.width * r.height
        }

    初始化：
        定义Rect类型后，如何创建并初始化Rect类型的对象实例
        rect1 := new(Rect)
        rect2 := &Rect{}
        rect3 := &Rect{0, 0, 100, 200}
        rect4 := &Rect{width: 100, height: 200}
        在Go语言中，未进行显式初始化的变量，都会被初始化为该类型的零值
        bool --> false
        int  --> 0
        string --> ""

        Go语言没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成。
        以NewXXX来命名，表示“构造函数”

        func NewRect(x, y, width, height float64) *Rect {
            return &Rect{x, y, width, height}
        }

    匿名组合：
        Go语言提供继承，但是采用了组合的文法，称为匿名组合
        type Base struct {
            Name string
        }
        func (base *Base) Foo() {...}
        func (base *Base) Bar() {...}
        type Foo struct {
            Base
            ...
        }
        type (foo *Foo) Bar() {
            foo.Base.Bar()
            ...
        }
        // 上面代码定义了一个Base类（实现了Foo()和Bar()两个成员方法）
        // 然后定义了一个Foo类，该类从Base类“继承”并改写了Bar()方法
        // 该方法实现时先调用了基类的Bar()方法
        // 调用 foo.Foo() 和 调用 foo.Base.Foo()效果一致

        Go语言很清晰告知类的内存布局情况
        在Go语言中可以随心所欲地修改内存布局

        type Foo struct {
            ... //其他成员
            Base
        }
        // 从语义上，这段代码和上面没有区别；但内存布局改变了。
        // “基类”Base的数据放在了“派生类”Foo的最后

        另外，在Go语言中，还可以以指针方式从一个类型“派生”
        type Foo struct {
            *Base
            ...
        }
        这段Go代码仍然有“派生”的效果，只是在Foo创建实例的时候，
        需要外部提供一个Base类实例的指针。

        在C++中有类似的功能，虚基类。

        匿名组合的一个小价值：
        type Job struct {
            Command string
            *log.Logger
        }
        在合适的赋值后，在Job类型的所有成员方法中可以借用所有log.Logger提供的方法。
        比如如下的写法：
        func (job *Job)Start() {
            job.Log("starting now...")
            .. //做一些事情
            job.Log("started.")
        }
        对于Job的实现者，无需意识到log.Logger类型的存在。
        这是匿名组合的魅力所在。

        不管是非匿名的类型还是匿名组合，
        被组合的类型所包含的方法虽然都升级成了
        外部这个组合类型的方法，
        但是它们被组合方法调用时接收者并没有改变。

        即使组合后调用的方式变成了 job.Log(...)
        但Log函数的接收者仍然是log.Logger指针，
        因此在Log中不可能访问到job的其他成员方法和变量。

        接口组合中的名字冲突问题：
        type X struct {
            Name string
        }
        type Y struct {
            X
            Name string
        }
        组合的类型和被组合的类型都包含一个Name成员
        会不会有问题？不会有问题。

        所有的Y类型的Name成员的访问都只会访问到
        最外层的那个Name变量
        X.Name变量相当于被隐藏。


        type Logger struct {
            Level int
        }

        type Y struct {
            *Logger
            Name string
            *log.Logger
        }

        这里会有问题，匿名组合类型相当于以其类型名称（去掉包名部分）
        作为成员变量的名字。
        按照规则，Y类型中相当于存在两个名为Logger的成员。

        这个编译错误，不是一定会发生的。
        加入这两个Logger在定义后再也没有被用过
        那么编译期会直接忽略这个冲突问题，
        直到开发者开始使用其中的某个Logger


    可见性
        Go语言很少增加关键字
        没有private、protected、public这样的关键字。
        需要让某个符号对其他包（package）可见，符号定义为大写开头。

        如：
        type Rect struct {
            X, Y float64
            Widtch, Height float64
        }

        成员方法的可见性，同样遵循相同的规则
        func (r *Rect) area() float64 {
            return r.Width * r.Height
        }





 */