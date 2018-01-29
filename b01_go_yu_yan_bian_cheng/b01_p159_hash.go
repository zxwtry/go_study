package b01_go_yu_yan_bian_cheng

import (
    "crypto/md5"
    "fmt"
    "crypto/sha1"
    "io/ioutil"
    "os"
    "io"
)

/**
 * User:  zxwtry
 * Date:  2018/1/31
 * Time:  16:22
 */

func P159Md5()  {
    str := "hello world"
    md5Ins := md5.New()
    md5Ins.Write([]byte(str))
    fmt.Printf("%x\n", md5Ins.Sum([]byte("")))
}

func P159Sha1() {
    str := "hello world"
    sha1Ins := sha1.New()
    sha1Ins.Write([]byte(str))
    fmt.Printf("%x\n", sha1Ins.Sum([]byte("")))
}

func P159Md5File()  {
    bs, err := ioutil.ReadFile("D:/vm_php/vmdk/ubuntu14.04.4server-s001.vmdk")
    if err == nil {
        md5Ins := md5.New()
        md5Ins.Write(bs)
        fmt.Printf("%x\n", md5Ins.Sum([]byte("")))
    } else {
        fmt.Println("err:", err)
    }
}

func P159Md5File2() {
    fileName := "D:/vm_php/vmdk/ubuntu14.04.4server-s001.vmdk"
    f, err := os.Open(fileName)
    if err == nil {
        md5Ins := md5.New()
        io.Copy(md5Ins, f)
        fmt.Printf("%x\n", md5Ins.Sum([]byte("")))
    } else {
        fmt.Println("err:", err)
    }
}

func P159Sha1File() {
    fileName := "F:/a.rar"
    bs, err := ioutil.ReadFile(fileName)
    if err == nil {
        sha1Ins := sha1.New()
        sha1Ins.Write(bs)
        fmt.Printf("%x\n", sha1Ins.Sum([]byte("")))
    } else {
        fmt.Println("err:", err)
    }
}

func P159Sha1File2() {
    fileName := "F:/a.rar"
    f, err := os.Open(fileName)
    if err == nil {
        sha1Ins := sha1.New()
        io.Copy(sha1Ins, f)
        fmt.Printf("%x\n", sha1Ins.Sum([]byte("")))
    } else {
        fmt.Println("err:", err)
    }
}
