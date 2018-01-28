package b01_go_yu_yan_bian_cheng

import (
    "fmt"
    "encoding/json"
    "os"
)

/**
 * User:  zxwtry
 * Date:  2018/1/26
 * Time:  10:53
 */



/*
     JSON
    1，json.Marshal()函数  --- 结构体
    2，Go中除了channel、complex、函数，其余都可以JSON
    3，浮点数、整数 --> 数字
    4，字符串 --> unicode字符串
    5，指针  --> 指向的内容
    6，数组  --> JSON中的数组
    7，map必须是 map[string]T
    8，结构体中，只有首字母大写，才会被JSON
 */


type book struct {
    Name string
    Author []string
    Press string
    isPublished bool
    Price float64
    Seller map[string]string
    readFunc func()
    PreBook *book
}
/*
    如果将readFunc修改成ReadFunc
    那么报运行错误：json: unsupported type: func()
    首字母小写，不会被JSON

    如果PreBook写成：PreBook book
    那么在初始化的时候：PreBook:*preBook,
    编译没有错，运行报错：invalid recursive type book
 */


func P136JsonBook() {
    preBook := & book {
        Name: "编程入门",
    }
    goBook := & book {
        Name: "Go语言编程",
        Author: []string {"许世伟","好多人"},
        Press: "图灵出版社",
        isPublished:true,
        Price: 100,
        Seller: map[string]string {
            "北京":"北京销售",
            "上海":"上海销售",
        },
        readFunc:P136JsonBook,
        PreBook:preBook,
    }
    fmt.Println(goBook)
    bs, err := json.Marshal(goBook)
    if err == nil {
        fmt.Println(string(bs))
    } else {
        fmt.Println(err)
    }
    /*
        &{Go语言编程 [许世伟 好多人] 图灵出版社 true 100}
        {"Name":"Go语言编程","Author":["许世伟","好多人"],"Press":"图灵出版社","Price":100}
     */
     var a map[string]string
     a = make(map[string]string)
     a["k1"] = "v1"
     a["k2"] = "v2"
     fmt.Println(a)
     // map[k1:v1 k2:v2]

     b := map[string]string {
         "k1": "v1",
     }
     b["k2"] = "v2"
     fmt.Println(b)
     // map[k1:v1 k2:v2]

     var newBook book;
     nbs,_ := json.Marshal(goBook)
     err2 := json.Unmarshal(nbs, &newBook)
     if err2 == nil {
         mbs, _ := json.Marshal(newBook)
         fmt.Println(string(mbs))
     } else {
         fmt.Println(err2)
     }
}


/*
    解码未知JSON
    布尔      --> bool
    数字      --> float64
    字符串    --> string
    数组      --> [] interface{}
    字典      --> map[string] interface{}
    null        --> nil
 */

func P136JsonUnknown() {
    preBook := & book {
        Name: "编程入门",
    }
    goBook := & book {
        Name: "Go语言编程",
        Author: []string {"许世伟","好多人"},
        Press: "图灵出版社",
        isPublished:true,
        Price: 100,
        Seller: map[string]string {
            "北京":"北京销售",
            "上海":"上海销售",
        },
        readFunc:P136JsonBook,
        PreBook:preBook,
    }
    bs, err := json.Marshal(goBook)
    if err != nil {
        fmt.Println("err:", err)
        os.Exit(-1)
    }
    var r interface{}
    err2 := json.Unmarshal(bs, &r)
    if err2 != nil {
        fmt.Println("err2:", err2)
        os.Exit(-1)
    }
    newBook, err3 := r.(map[string]interface{})
    if ! err3 {
       fmt.Println("err3", err3)
        os.Exit(-1)
    }

    for k, v := range newBook {
        fmt.Println("key:", k)
        fmt.Println("val:", v)
        fmt.Println()
    }
    /*
        key: Author
        val: [许世伟 好多人]

        key: Press
        val: 图灵出版社

        key: Price
        val: 100

        key: Seller
        val: map[上海:上海销售 北京:北京销售]

        key: PreBook
        val: map[Seller:<nil> PreBook:<nil> Name:编程入门 Author:<nil> Press: Price:0]

        key: Name
        val: Go语言编程
     */
}


func P136EncodeDecode()  {
    bs, e := json.Marshal(map[string]string{"k1":"v1"})
    if e == nil {
        fmt.Println(string(bs))
    }
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            fmt.Println(err)
            os.Exit(-1)
        }
        for k, v := range v {
            fmt.Println("key:", k)
            fmt.Println("val:", v)
            fmt.Println()
        }
        if err := enc.Encode(v); err != nil {
            fmt.Println(err)
            os.Exit(-1)
        }
    }
}