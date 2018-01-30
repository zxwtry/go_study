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
    fmt.Printf("value kind:\t%v\n", reflect.ValueOf(v).Kind() == reflect.Float64)
    /*
        value:	1.2
        type:	float64
        value kind:	float64
        value kind:	true
     */
}

