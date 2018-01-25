package b01_go_yu_yan_bian_cheng

import (
    "net/http"
    "fmt"
    "log"
    "time"
    "errors"
    "net/rpc"
    "net"
)

/**
 * User:  zxwtry
 * Date:  2018/1/20
 * Time:  11:24
 */

/*
    HTTP编程：
        HyperText Transfer Protocol
        超文本传输协议

    net/http包

    HTTP客户端
    无需借助第三方库，使用GET、POST方式请求数据。

    基本方法：
        func (c *Client) Get(url string) (r *Response, err error)
        func (c *Client) Post(url string, bodyType string, body io.Reader) (r *Response, err error)
        func (c *Client) PostForm(url string, data url.Values) (r *Response, err error)
        func (c *Client) HEAD(url string) (r *Response, err error)
        func (c *Client) Do(req *request) (r *Response, err error)

    http.Get()      等价于   http.DefaultClient.Get()
        resp, err := http.Get("http://www.zxwtry.com/")
        if err != nil {
            // 处理错误
            return
        }

        defer resp.Body.close()
        io.Copy(os.Stdout, resp.Body)
        // 上面这段代码请求一个网站，并将其网页内容打印到标准输出流中

    http.Post()
        以POST方式发送数据，http.Post()方法
        然后传入下面的参数：
        1，请求的目标URI
        2，将要POST数据的资源类型(MIMEType)
        3，数据的比特流([]byte 形式)

        下面的代码演示如何上传一张图片：
        resp, err := http.Post("http://example.com/upload",
                    "image/jpeg", &imageDataBuf)
        if err != nil {
            // 处理错误
            return
        }

        if resp.StatusCode != http.StatusOK {
            // 处理错误
            return
        }
        // ...


    http.PostForm()
        http.PostForm()方法，实现了
        标准编码格式为 application/x-www-form-urlencoded
        的表单提交。

        resp, err := http.PostForm("http://zxwtry.com/posts",
            url.Values{"title" : {"article title"}, "content": {"article body"}})


    http.Head()
        HTTP中的Head请求方式表明只请求目标URL的头部信息，
        即HTTP Header而不返回HTTP Body。

        Head()方法和Get()方法一样，只需要传入目标URL一个参数即可

        resp, err := http.Head("http://zxwtry.com/")


    (*http.Client).Do()
        发起的HTTP请求需要更多的定制信息
        希望设定一些自定义的HTTP Header字段
        比如：
            传递自定义的"User-Agent"，而不是磨人的"Go http package"
            传递Cookie

        此时可以使用net/http包中的http.Client对象的Do()方法实现

        req, err := http.NewRequest("GET", "http://zxwtry.com", nil)

        req.Header.Add("User-Agent", "Gobook Custom User-Agent")

        Client := &http.Client{}

        resp, err := Client.Do(req)
 */


/*
    之前的http.Get()、http.Post()、http.PostForm()、http.Head()
    都是http.DefaultClient.Get() ...

    Client的结构：
        type Client struct {
            // Transport用于确定Http的创建机制
            // 如果没有Transport，使用默认的DefaultTransport
            // http.Transport实现了http.RoundTripper接口
            Transport RoundTripper

            // CheckRedirect用于处理重定向策略
            // 如果CheckRedirect不为空，Client将在跟踪HTTP重定向前，调用CheckRedirect
            // req是即将发起的请求、via是已经发起的请求，最早访问的在最前面。
            // 如果CheckRedirect返回错误，Client将返回错误并停止。
            // 如果CheckRedirect为空，Client将在连续的10次请求后停止。
            // GET/HEAD时候，当响应返回30x(301, 302, 303, 307)，Client会调用CheckRedirect
            CheckRedirect func(req *Request, via []*Request) error

            // 如果Jar为空，不发送cookie，并在响应中忽略cookie
            // http.CookieJar中实现了 SetCookies()和Cookies()两个方法。
            // 一般通过http.SetCookie()方法去设定cookie
            Jar CookieJar
        }


    发送自定义请求：
        client := &http.Client{
            CheckRedirect: redirectPolicyFunc
        }

        resp, err := client.Get("http://zxwtry.com")

        req, err := client.NewRequest("GET", "http://zxwtry.com", nil)

        req.Header.Add("User-Agent", "My User Agent");
        req.Header.Add("If-None-Match", `w/"MyFileEtag`)

        resp, err := client.Do(req)


    Transport数据结构：

        type Transport struct {
            // Proxy指定：返回代理的函数
            // 如果error不为空，将停止执行，并返回错误
            // 如果Proxy为空、返回URL为空，将不使用代理
            Proxy func(*Request) (*url.URL, error)

            // Dial指定：创建TCP连接的dial()函数
            // 如果Dial为空，那么将使用默认的net.Dial()
            Dial func(net, addr string) (c net.Conn, err error)

            // TLSClientConfig指定：tls.Client的TLS配置
            // ssl专用
            TLSClientConfig *tls.Config

            DisableKeepAlives bool      // 默认false，启用长连接
            DisableCompression bool     // 默认false，启用压缩

            // 每个Host能维持的最大空连接数（非活跃连接）
            // 如果没有指定，使用DefaultMaxIdleConnsPerHost
            MaxIdleConnsPerHost
        }

        Transport还有三个函数：

            func (t *Transport) CloseIdleConnections()   关闭所有非活跃连接

            // 注册协议，比如File、FTP等等协议
            func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)

            // 实现RoundTripper接口
            func (t *Transport) RoundTrip(req *Request) (resp *Response, err error)


        自定义Client
            tr := &http.Transport {
                TLSClientConfig: &tls.Config {RootCAs: pool}
                DisableCompression: true
            }

            client := &http.Client {
                Transport: tr
            }

            resp, err := client.Get("http://zxwtry.com")

            Client和Transport在多个goroutine中执行是线程安全的
            可以创建一个Client和Transport，在多个goroutine中使用
 */


/*
    RoundTripper接口

    type RoundTripper interface {

        // RoundTrip执行单一HTTP请求并获得响应
        // err不为空：没有正确获取响应
        // 如果有响应，不管HTTP状态码是什么，err都为空
        // RoundTrip不会去理解更高级内容：认证、重定向、cookie等
        // RoundTrip不会修改请求内容
        RoundTrip (*Request) (*Response, err error)
    }
 */


type MyTransport struct {
    Transport http.RoundTripper
}

func (t * MyTransport) transport() http.RoundTripper {
    if t.Transport != nil {
        return t.Transport
    }
    return http.DefaultTransport
}

func (t * MyTransport) roundTrip(req *http.Request) (*http.Response, error) {
    return t.transport().RoundTrip(req)
}

func (t *MyTransport) client() * http.Client {
    return &http.Client{
        Transport: t.transport(),
    }
}

func P124MyTransport() {
    tr := &MyTransport{}
    client := tr.client()
    resp, err := client.Get("http://zxwtry.com")
    if err == nil {
        fmt.Println("right")
        fmt.Println(resp)
    } else {
        fmt.Println("error")
        fmt.Println(err)
    }
    /*
        输出：

    right
    &{
        200 OK 200
        HTTP/1.1 1 1
        map[    Etag:[W/"3683-1507730850000"]
                Last-Modified:[Wed, 11 Oct 2017 14:07:30 GMT]
                Content-Type:[text/html]
                Content-Length:[3683]u
                Date:[Tue, 23 Jan 2018 08:16:17 GMT]
                Server:[unknown-server]
                Accept-Ranges:[bytes]
        ]
        0xc0420362c0 3683 [] false false
        map[] 0xc0420e0000 <nil>
    }
     */
}


func P124MyTransportCompare() {
    client := &http.Client{}
    resp, err := client.Get("http://zxwtry.com")
    if err == nil {
        fmt.Println("right")
        fmt.Println(resp)
    } else {
        fmt.Println("error")
        fmt.Println(err)
    }
    /*
        得到的返回信息 和 P124MyTransport一样
     */
}



/*
    Client是业务层
    非业务层细节：
        1，HTTP传输过程
        2，代理
        3，GZIP
        4，连接池
        5，认证(SSL等)
 */


/*
    处理HTTP请求
    func ListenAndServe(addr string, handler Handler) error
    如果 handler为空，默认使用 http.DefaultServeMux
    编写的http.Handler()和http.HandlerFunc()会注入到http.DefaultServeMux

    http.Handler("/foo", fooHandler)
    http.HandlerFunc("/bar", func(w ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))


    自定义http.Server

    s := &http.Server {
        Addr: ":9998",
        Handler: http.DefaultServeMux,
        ReadTimeOut: 10 * time.Second,
        WriteTimeOut: 10 * time.Second,
        MaxHeaderBytes: 1 << 20
    }
 */

func P124HttpServe() {
    log.Fatal(http.ListenAndServe(":9999", nil))
}


func P124HttpServer() {
    s := &http.Server {
        Addr: ":9998",
        Handler: http.DefaultServeMux,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())
}


/*
    HTTPS：需要认证文件和认证私钥
    http.ListenAndServeTLS(addr string,
        verFile string, verKey string, handler Handler)

    s := &http.Server {
        Addr: ":9996",
        Handler: http.DefaultServeMux,
        ReadTimeOut: 10 * time.Second,
        WriteTimeOut: 10 * time.Second,
        MaxHeaderBytes: 1 < 20
    }

    log.Fatal(s.ListenAndServeTLS(verFile string, verKey string))
 */


/*
    RPC特点：
    1，首字母大写
    2，类型：外部包可访问、Go内置类型
    3，第二个参数是指针
    4，返回一个error

    func (t *T) MethodName(arg T1, ret *T2) error {}
    T、T1、T2都必须是encoding/gob能够编码的

    可以通过rpc.ServeConn()来处理单个连接
 */


/*
    rpc提供rpc.Dial()和rpc.DialHttp()
    rpc.Call()：同步
    rpc.Dial()：异步
 */


type Arith int


type Args struct {
    A, B int
}


type Quotient struct {
    Quo, Rem int
}


func (a *Arith) Multiply(arg *Args, quo *int) error {
    *quo = arg.A * arg.B
    return nil
}


func (a *Arith) Divide(arg *Args, quo *Quotient) error {
    if arg.B == 0 {
        err := errors.New("被除数为0")
        return err
    }
    quo.Quo = arg.A / arg.B
    quo.Rem = arg.A % arg.B
    return nil
}


func P124RPC_MD() {
    arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()

    l, e := net.Listen("tcp", ":9995")
    if e != nil {
        fmt.Println("RPC 服务出错")
        fmt.Println(e)
    }
    //go http.Serve(l, nil)
    http.Serve(l, nil)
}


func P124RPC_Client() {
    client, err := rpc.DialHTTP("tcp", "localhost:9995")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    args := &Args{7, 8}
    var reply int
    e := client.Call("Arith.Multiply", args, &reply)
    if e != nil {
        log.Fatal("call:", e)
    }
    fmt.Println("reply is", reply)
    var re2 Quotient
    e2 := client.Call("Arith.Divide", args, &re2)
    if e2 != nil {
     log.Fatal("call re2:", e)
    }
    fmt.Println("re2 is", re2.Quo, re2.Rem)
    /*
        reply is 56
        re2 is 0 7
     */
}


func P124RPC_ClientGo() {
    client, err := rpc.DialHTTP("tcp", "localhost:9995")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    args := &Args{A:7, B:8}
    var reply int
    e := client.Go("Arith.Multiply", args, &reply, nil)
    done := <- e.Done
    if done.Error == nil {
        fmt.Println("no error")
        fmt.Println("result is", reply)
    } else {
        fmt.Println("reply error")
        fmt.Println("error is", done.Error)
    }
    var re2 Quotient
    e2 := client.Go("Arith.Divide", args, &re2, nil)
    done2 := <- e2.Done
    if done2.Error == nil {
        fmt.Println("no error")
        fmt.Println("result is", done2.Reply)
    } else {
        fmt.Println("reply error")
        fmt.Println("error is", done2.Error)
    }
    /*
        no error
        result is 56
        no error
        result is &{0 7}
     */
}


/*
    Gob：二进制流，Go语言专用
    自解析、高效率、完整表达能力
 */



/*
    RPC编码解码

    type ClientCodec interface {
        WriteRquest(*Request, interface{}) error
        ReadResponseHeader(*Response) error
        ReadResponseBody(interface{}) error

        Close() error
    }

    type ServerCodec interface {
        ReadRequestHeader(*Request) error
        ReadRequestBody(interface{}) error
        WriteResponse(*Response, interface{}) error

        Close() error
    }
 */