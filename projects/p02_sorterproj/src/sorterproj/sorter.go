package main

import "flag"
import "fmt"
import "bufio"
import "io"
import "os"
import "strconv"
import "time"
import "algorithms/bubblesort"
import "algorithms/qsort"

/*
    任务：
        获取并解析命令行输入
        从对应文件中读取输入数据
        调用对应的排序函数
        将排序的结果输出到对应的文件中
        打印排序所花费时间的信息
*/

/*
    项目编译过程：
    set GOPATH={sorterproj所在目录}
    mkdir bin
    cd bin

    go build sorterproj     // 在bin目录生成exe文件

    在bin目录下准备一份unsort.txt文件，运行
    sorterproj.exe -i unsort.txt -o sorted.txt -a qsort

    测试：go test algorithms/qsort

    go install algorithms/qsort
        // 在pkg/windows_amd64/algorithms
        // 生成qsort.a文件

    go install algorithms/bubblesort
        // 在pkg/windows_amd64/algorithms
        // 生成bubblesort.a文件

 */


var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "file to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")


func readValues(infile string) (values []int, err error) {
    file, err := os.Open(infile)
    if err != nil {
        fmt.Println("Failed to open the input file ", infile)
        return
    }
    defer file.Close()
    br := bufio.NewReader(file)
    values = make([]int, 0)
    for {
        line, isPrefix, err1 := br.ReadLine()
        if err1 != nil {
            if err1 != io.EOF {
                err = err1
            }
            break
        }
        if isPrefix {
            fmt.Println("A too long line, seems unexpected.")
            return
        }
        str := string(line) //转换字符数组为字符串
        value, err1 := strconv.Atoi(str)
        if err1 != nil {
            err = err1
            return
        }
        values = append(values, value)
    }
    return
}


func WriteValues(values []int, outfile string) error {
    file, err := os.Create(outfile)
    if err != nil {
        fmt.Println("Failed to create the output file: ", outfile)
        return err
    }
    defer file.Close()
    for _,value := range values {
        str := strconv.Itoa(value)
        file.WriteString(str + "\n")
    }
    return nil
}


func main() {
    flag.Parse()
    if infile != nil {
        fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
    }
    values, err := readValues(*infile)
    if err == nil {
        t1 := time.Now()
        switch *algorithm {
        case "qsort":
            qsort.QSort(values)
        case "bubblesort":
            bubblesort.BubbleSort(values)
        default:
            fmt.Println("algorithm", *algorithm, "is unknown")
        }
        fmt.Println("cost times: ", (time.Now().Sub(t1)))
        WriteValues(values, *outfile)
    } else {
        fmt.Println(err)
    }
    /*
        sorterproj.exe -i unsort.data -o sorted.data -a bubblesort
        infile = unsort.data outfile = sorted.data algorithm = bubblesort
    */
}
