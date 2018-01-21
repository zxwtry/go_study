package b01_go_yu_yan_bian_cheng

import (
    "net"
    "fmt"
    "os"
    "bytes"
    "io"
    "io/ioutil"
)

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

        ICMP(使用协议名称)：
            conn, err := net.Dial("ip4:icmp", "www.zxwtry.com")

        ICMP(使用协议编号)：
            conn, err := net.Dial("ip4:1", "10.0.0.3")


    Dial支持的网络协议：
        tcp     tcp4    tcp6
        udp     udp4    udp6
        ip      ip4     ip6
 */


func readFully(conn net.Conn) ([]byte, error)  {
    defer conn.Close()

    result := bytes.NewBuffer(nil)
    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        result.Write(buf[0:n])
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
    }
    return result.Bytes(), nil
}


func checkSum(msg []byte) uint16  {
    sum := 0

    // 先假设为偶数
    for n := 1; n < len(msg) - 1; n +=2 {
        sum += int(msg[n]) * 256 + int(msg[n + 1])
    }
    sum = (sum >> 16) + (sum & 0xffff)
    sum += sum >> 16
    var answer uint16 = uint16(^sum)
    return answer
}


func checkError(err error)  {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        fmt.Printf("Fatal error: %s", err.Error())
        os.Exit(1)
    }
}


func P118ICMP() {
    service := "www.baidu.com"
    conn, err := net.Dial("ip4:icmp", service)
    checkError(err)

    var msg [512]byte

    msg[0] = 8  //  echo
    msg[1] = 0  //  code 0
    msg[2] = 0  //  checksum
    msg[3] = 0  //  checksum
    msg[4] = 0  //  identifier[0]
    msg[5] = 13 //  identifier[1]
    msg[6] = 0  //  sequence[0]
    msg[7] = 37 //  sequence[1]

    len := 8
    check := checkSum(msg[0:len])
    msg[2] = byte(check >> 8)
    msg[3] = byte(check & 255)

    _, err = conn.Write(msg[0:len])
    checkError(err)

    _, err = conn.Read(msg[0:])
    checkError(err)

    fmt.Println("Got response")

    if msg[5] == 13 {
        fmt.Println("Identifier matches.")
    }
    if msg[7] == 37 {
        fmt.Println("Sequence matches.")
    }
    os.Exit(0)
}


func P118TCP()  {
    service := "www.zxwtry.com:80"

    conn, err := net.Dial("tcp", service)
    checkError(err)

    _, err = conn.Write([] byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    result, err := readFully(conn)
    checkError(err)

    fmt.Println("result is:")
    fmt.Println(string(result))

    os.Exit(0)

    /*
        result is:
        HTTP/1.1 200 OK
        Accept-Ranges: bytes
        ETag: W/"3683-1507730850000"
        Last-Modified: Wed, 11 Oct 2017 14:07:30 GMT
        Content-Type: text/html
        Content-Length: 3683
        Date: Sat, 20 Jan 2018 02:53:24 GMT
        Connection: close
        Server: unknown-server
     */
}


/*
    Dial函数是对DialTCP、DialUDP、DialIP、DialUnix的封装
    可以直接调用这些函数：
    func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err error)
    func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err error)
    func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, err error)
    func DialUnix(net string, laddr, raddr *UnixAddr) (c *UnixConn, err error)
 */


func P118TCP2()  {
    service := "www.zxwtry.com:80"

    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    result, err := ioutil.ReadAll(conn)
    checkError(err)

    fmt.Println(string(result))

    os.Exit(0)

    /*
        HTTP/1.1 200 OK
        Accept-Ranges: bytes
        ETag: W/"3683-1507730850000"
        Last-Modified: Wed, 11 Oct 2017 14:07:30 GMT
        Content-Type: text/html
        Content-Length: 3683
        Date: Sat, 20 Jan 2018 03:14:50 GMT
        Connection: close
        Server: unknown-server
     */

     /*
        与之前使用的Dial例子不同：
            net.ResolveTCPAddr()，用于解析地址和端口号
            net.DialTCP()，用户建立连接
            这两个函数都在Dial里面进行了封装
      */
}


/*
    验证IP地址有效性：
        func net.ParseIP()
    创建子网掩码
        func IPv4mask(a, b, ,c, d byte) IPMask
    获取默认子网掩码
        func (ip IP) DefaultMask() IPMask
    根据域名查找IP的
        func ResolveIPAddr(net, addr string) (*IPAddr, error)
        func LookupHost(name string) (cname string, addrs []string, ,err error)
 */