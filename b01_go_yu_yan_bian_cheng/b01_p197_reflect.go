package b01_go_yu_yan_bian_cheng

import (
    "io"
    "reflect"
    "fmt"
)



/**
 * User:  zxwtry
 * Date:  2018/1/31
 * Time:  17:08
 */

type MyName struct{
    Name string
}

/*
    io.Reader接口只有一个函数Read()
    实现了函数Read()，就实现了接口 io.Reader
 */
func (m MyName) Read(p []byte) (n int, err error) {
    return 100, nil;
}
 
func P197Reflect()  {
    var name io.Reader
    name = &MyName{"aabb"}
    fmt.Printf("value:\t%v\n", reflect.ValueOf(name))
    fmt.Printf("type:\t%v\n", reflect.TypeOf(name))
    fmt.Printf("value kind is:\t%v\n", reflect.ValueOf(name).Kind())
    /*
        value:	&{aabb}
        type:	*b01_go_yu_yan_bian_cheng.MyName
        value kind is:	ptr
     */
}

func P197Reflect2() {
    var v float64
    v = 1.2
    fmt.Printf("value:\t%v\n", reflect.ValueOf(v))
    fmt.Printf("type:\t%v\n", reflect.TypeOf(v).Name())
    fmt.Printf("value kind:\t%v\n", reflect.ValueOf(v).Kind())
    fmt.Printf("value kind == reflect.Float64:\t%v\n", reflect.ValueOf(v).Kind() == reflect.Float64)
    fmt.Printf("Value类型的Float方法：\t%v\n", reflect.ValueOf(v).Float())

    //fmt.Printf("Value类型的Int方法：\t%v\n", reflect.ValueOf(v).Int())
    //panic: reflect: call of reflect.Value.Int on float64 Value

    fmt.Printf("Value类型的String方法：\t%v\n", reflect.ValueOf(v).String())

    /*
        value:	1.2
        type:	float64
        value kind:	float64
        value kind == reflect.Float64:	true
        Value类型的Float方法：	1.2
        Value类型的String方法：	<float64 Value>
     */
}

func P198ValueType()  {
    /*
        Go中，所有的类型都是值类型
        这些变量在传输给函数的时候，会发生一次复制
     */

    /*
        var x float64 = 3.4
        v := reflect.ValueOf(x)
        v.Set(4.1)  这句会报错
        通过reflect.ValueOf(x)创建一个x的副本
        如果v允许set，修改的将是x的副本

        在Go中，引入可设属性这个概念（Settability）
        如果CanSet()返回false，不应该调用Set()和SetXxx()方法
    */
    var x float64 = 3.4
    v := reflect.ValueOf(x)
    fmt.Printf("v.Type() is\t%v\n", v.Type())
    fmt.Printf("v.CanSet() is\t%v\n", v.CanSet())
    // v.Type() is	float64
    // v.CanSet() is	false

    u := reflect.ValueOf(&x)
    fmt.Printf("u.Type() is\t%v\n", u.Type())
    fmt.Printf("u.CanSet() is\t%v\n", u.CanSet())
    // u.Type() is	*float64
    // u.CanSet() is	false

    w := u.Elem()
    fmt.Printf("w.Type() is\t%v\n", w.Type())
    fmt.Printf("w.CanSet() is\t%v\n", w.CanSet())
    // w.Type() is	float64
    // w.CanSet() is	true

    w.SetFloat(5.5)
    fmt.Printf("v.Interface() is\t%v\n", v.Interface())
    fmt.Printf("v is\t%v\n", v)
    fmt.Printf("x is\t%v\n", x)
    // v.Interface() is	3.4
    // v is	3.4
    // x is	5.5
}

/*
    对结构的反射操作
 */

type P197T struct {
    A int
    B string
}

func P199StructReflect() {
    t := P197T{101,"101string"}
    s := reflect.ValueOf(&t).Elem()
    typeOfT := s.Type()
    for i := 0; i < s.NumField(); i ++ {
        f := s.Field(i)
        fmt.Printf("%d: %s %s = %v %v\n", i, typeOfT.Field(i).Name,
            f.Type(), f.Interface(), f.CanSet())
    }
    // 0: A int = 101 true
    // 1: B string = 101string true
}

