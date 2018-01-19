package b01_go_yu_yan_bian_cheng

/**
 * User:  zxwtry
 * Date:  2018/1/18
 * Time:  18:22
 */



/*
    网络编程：
        1，建立Socket
        2，绑定Socket
        3，监听
        4，接受连接
        5，接收

        Go语言：net.Dial()


    Dial()函数
        函数原型：
        func Dial(net, addr string) (Conn, error)

    常见协议的调用方式：
        TCP：
            conn, err := net.Dial("tcp", "192.168.0.10:2100")

        UDP：
            conn, err := net.Dial("udp", "192.168.0.12:975")

        ICMP：
            conn, err := net.Dail()

 */


