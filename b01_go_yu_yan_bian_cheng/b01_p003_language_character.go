package b01_go_yu_yan_bian_cheng

import (
	"fmt"
	"reflect"
	"unsafe"
)
/*
#include<stdio.h>
#include<stdlib.h>
 */
import "C"

/**
 * User:  zxwtry
 * Date:  2017/12/29
 * Time:  11:25
 */


/*
	Go语言最主要的特性：
	1，	自动垃圾回收
	2，	丰富的内置类型
	3，	函数多返回值
	4，	错误处理
	5，	匿名函数和闭包
	6，	类型和接口
	7，	并发编程
	8，	反射
	9，	语言交互性
 */


 /*
 	匿名函数和闭包：
 	在Go语言中，所有函数也是值类型，可以作为参数传递。
 	下列代码定义了一个叫f的匿名函数，可以随意对该匿名函数变量进行传递和调用
 	f := func(x, y int) int {
 		return x - y;
 	}
 */


/*
	类型和接口：
	1， 不支持继承和重载
	2， 支持最基本的类型组合
	3， 非“侵入式”接口

 */

type Bird struct {

}

func (b *Bird) Fly() {
	fmt.Println("飞。。。")
}

type IFly interface {
	Fly()
}

func P003Bird()  {
	var fly IFly = new (Bird)
	fly.Fly()
	// Bird类型实现时，没有声明与接口IFly的关系
	// 接口和类型可以直接转换
	// 接口的定义不需要在类型定义之前
	/*
		输出：
			飞。。。
	 */
}


/*
	并发编程：
	1， goroutine使用消息传递来共享内存，不是裸用OS的并发机制
	2， 函数调用前使用关键字go，让该函数以goroutine方式执行
	3， goroutine是一种比线程更加轻盈、更省资源的协程
	4， Go用通道实现CSP（通信顺序进程），进行跨goroutine通信
	5， Go读写锁：一个进程的所有goroutine运行在同一内存地址空间
 */

func sum(values[] int, resultChan chan int)  {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum	// 将计算结果发送到channel中
}

func P003Sum() {
	values := [] int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	resultChan := make(chan int, 2)
	go sum(values[: len(values) / 2], resultChan)
	go sum(values[len(values) / 2 :], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan //接收结果
	fmt.Println("result:", sum1, sum2, sum1 + sum2)
	/*
		输出：
		result: 35 10 45
	 */
}


/*
	反射：
	1，	没有内置类型工厂，无法通过类型字符串创建对象实例
	2，	使用场景：对象的序列化
 */

type CatReflect struct {
	Name string
	LifeExpect int
}

func MainReflect()  {
	cat := &CatReflect{"bosi", 1}
	catRef := reflect.ValueOf(cat).Elem()
	typeOfT := catRef.Type()
	for i := 0; i < catRef.NumField(); i ++ {
		f := catRef.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name,
			f.Type(), f.Interface())
	}
	/*
		0: Name string = bosi
		1: LifeExpect int = 1
	 */
}


/*
	语言交互性：
	1，	Go重用现有C模块   -->  Cgo
 */

func P003Cgo()  {
	str := C.CString("Hello World")
	C.puts(str)
	C.free(unsafe.Pointer(str))
}

