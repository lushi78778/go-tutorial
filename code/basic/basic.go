package main // 声明一个主包（包是对函数进行分组的一种方式，它由同一目录中的所有文件组成）

import "fmt"  // 导入fmt包，其中包含格式化文本的功能，包括打印到控制台。此程序包是安装Go时获得的标准库程序包之一。

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"


func main() { 
    // // 声明一个变量并初始化
    // var a = "lushi"
    // fmt.Println(a)

    // // 没有初始化就为零值
    // var b int
    // fmt.Println(b)

    // // bool 零值为 false
    // var c bool
    // fmt.Println(c)

	// var d = true
	// fmt.Println(d)

	// e := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	// fmt.Println(e)

	//这种不带声明格式的只能在函数体中出现
	g, h := 123, "hello"
    fmt.Println(x, y, a, b, c, d, e, f, g, h)
}