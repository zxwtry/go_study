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
		var v6 *int		// 指针
		var v7 map[string] int  	// map key为string, value为int
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



 */
