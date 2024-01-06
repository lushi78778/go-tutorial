# Getting Started


本教程的示例代码为 hello

接下来，你将用约六十分钟的时间，快速完成go的安装与入门，包括第一个hello world和关于模块的小教程

在本教程中，你将编写打包到两个模块中的函数：一个带有发送问候语的逻辑的模块（greetings）;另一个是作为模块的调用方（hello）。

## 安装 Go {id="go_1"}

访问一下地址以获取go：https://go.dev/doc/install

```Bash 
# 验证安装
go version
```
go version go1.21.5 windows/amd64

## 第一个脚本：hello world

在你认为合适的目录新建文件夹hello，用VS code 打开

```Bash
# 初始化项目
go mod init hello      

# 定义系统 （最好加上环境变量！！！！！！！！！！！！！！！！！）                                        
$env:GOOS="windows"
```

新建```hello.go```

```Go
package main // 声明一个主包（包是对函数进行分组的一种方式，它由同一目录中的所有文件组成）

import "fmt"  // 导入fmt包，其中包含格式化文本的功能，包括打印到控制台。此程序包是安装Go时获得的标准库程序包之一。

func main() { 
    //实现一个主要功能，将消息打印到控制台。当您运行主程序包时，默认情况下会执行一个主函数。
    fmt.Println("Hello, World!")
}
```



```Bash
# 运行 hello.go
go run hello.go
```

好了，现在你已经学会了最基础的一部分，包括注释写法

## 外部包的使用

修改刚刚的代码

```Go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

运行一下语句

```Bash
# 添加新的模块要求到 go.sum
go mod tidy

# 运行 hello.go
go run .
# 或者
go run .\hello.go
```

参考输出

```Bash
PS D:\projectCode\go\hello> go run .\hello.go
Don't communicate by sharing memory, share memory by communicating.
```

如果想要查询具体的外部包信息，请访问：https://pkg.go.dev/

比如quote包：https://pkg.go.dev/search?q=quote
![pkg-search.png](pkg-search.png)

可以在内页查询具体的版本信息
![pkg-search-info](pkg-search-info.png)

##  创建 Go 模块

还是在我刚刚的文件夹内，请看目录结构

```Bash
hello
│   go.mod
│   go.sum
│   hello.go   # 上面是之前建立的
│
├───greetings
│       go.mod
│       greetings.go
│
└───hello
        go.mod
        hello.go
```

现在我们一点点建立项目结构，新建greetings，greetings.go；新建hello，hello.go

greetings 下初始化项目

```Bash
# 这里example.com/greetings是命名模块的基本范式，<prefix>/<descriptive-text> eg:github.com/<project-name>
go mod init example.com/greetings 
```

```Go
//greetings.go

package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

hello 下初始化项目

```Bash
go mod init example.com/hello 
```

```Go
//hello.go

package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```

现在你可能大致了解了，但是有一个疑问，匿名没有把example.com/greetings 发布到网上，hello.go怎么调用呢呢？

```Bash
# 在 hello 目录中的命令提示符下，运行以下命令
go mod edit -replace example.com/greetings=../greetings

# 现在你再去看看hello\go.mod，然后同步模块的依赖项，再去观察
go mod tidy # 该命令在 greetings 目录中找到了本地代码，然后 添加了 require 指令以指定导入包时的依赖关系

go run .
```

输出样例

```Bash
PS D:\projectCode\go\hello\hello> go mod edit -replace example.com/greetings=../greetings   
PS D:\projectCode\go\hello\hello> go mod tidy
go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000 
PS D:\projectCode\go\hello\hello> go run .
Hi, Gladys. Welcome!
```

好了，我想你已经掌握了一些小技巧，让我们继续，做一些异常处理

## 异常处理

修改greetings\greetings.go

```Go
// greetings.go

package greetings

import (
    "errors" //导入 Go 标准库包，以便您可以 使用其错误。新功能。errors
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    if name == "" { //添加语句以检查无效请求（ 名称应为空字符串），并在 请求无效。该函数返回一个包含您的消息的函数。iferrors.Newerror
        return "", errors.New("empty name") //更改函数，使其返回两个值：a 和 .检查 第二个值用于查看是否发生错误。（任何 Go 函数都可以 返回多个值。有
    }

    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil // 添加（表示没有错误）作为 成功返回。这样，调用方可以看到函数 成功。nil
}
```

修改 hello\hello.go

```Go
//hello.go

package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ") // 将日志包配置为 在其日志消息的开头打印命令名称 （“greetings： ”）， 没有时间戳或源文件信息。
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")

    if err != nil {
        //使用标准库中的函数 输出错误信息。如果出现错误，请使用包的 Fatal 函数打印错误并停止程序。
        log.Fatal(err)
    }


    fmt.Println(message)
}
```

`在hello路径下`（hello/hello）,运行`hello.go`

输出样例

```Bash
PS D:\projectCode\go\hello\hello> go run .
greetings: empty name
exit status 1
```

## 返回随机问候语


修改greetings\greetings.go

```Go
// greetings.go

package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}


func randomFormat() string {  //添加一个函数，该函数随机返回 选定的问候语格式。
    // 声明一个切片 三种消息格式
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    使用 math/rand 包生成一个随机数，用于从切片中选择项目
    return formats[rand.Intn(len(formats))]
}
```

修改 hello\hello.go

```Go
// hello.go 只是添加 Gladys 的名字（或者如果您愿意，可以改用其他名字） 作为 hello.go 中函数调用的参数。

package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ") // 将日志包配置为 在其日志消息的开头打印命令名称 （“greetings： ”）， 没有时间戳或源文件信息。
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("Gladys") //这里

    if err != nil {
        //使用标准库中的函数 输出错误信息。如果出现错误，请使用包的 Fatal 函数打印错误并停止程序。
        log.Fatal(err)
    }


    fmt.Println(message)
}
```

`在hello路径下`（hello/hello）,运行`hello.go`

输出样例

```Bash
PS D:\projectCode\go\hello\hello> go run .
Hi, Gladys. Welcome!
PS D:\projectCode\go\hello\hello> go run .
Hail, Gladys! Well met!
PS D:\projectCode\go\hello\hello> go run .
Great to see you, Gladys!
```

## 在一个请求中获取多个人的问候语

修改greetings\greetings.go

```Go
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

//新添加的函数 添加一个参数为名称切片的函数
func Hellos(names []string) (map[string]string, error) {

    messages := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }

        messages[name] = message
    }
    return messages, nil
}


func randomFormat() string {

    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    
    return formats[rand.Intn(len(formats))]
}
```

修改 hello\hello.go

```Go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // 创建一个变量作为包含三个的切片类型 名字。names
    names := []string{"Gladys", "Samantha", "Darrin"}

    // 将变量作为参数传递给函数。namesHellos
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}
```

其实就是用循环输出了好多人的问候语

输出样例

```Bash
PS D:\projectCode\book\go-tutorial\code\hello\hello>  go run .
map[Darrin:Great to see you, Darrin! Gladys:Hail, Gladys! Well met! Samantha:Great to see you, Samantha!]
```

## 测试

以 `_test.go` 结尾的文件名会告诉命令 此文件包含测试函数。使用 `go test`可以进行测试。

- 在 greetings 目录中，创建名为 greetings_test.go 的文件
- 在 greetings_test.go 中，粘贴以下代码并保存文件

```Go
//greetings_test.go

package greetings

import (
    "testing"
    "regexp"
)

//在与要测试的代码相同的包中实现测试函数。

// TestHelloName调用函数， 传递函数应具有的值 能够返回有效的响应消息。
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty调用函数 替换为空字符串。该测试旨在确认错误处理有效。
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    // 如果调用返回非空字符串或错误，则使用参数的方法将消息打印到控制台并结束执行。
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
```

测试应通过:

```Bash
PS D:\projectCode\book\go-tutorial\code\hello\greetings> go test
PASS
ok      example.com/greetings   0.126s
PS D:\projectCode\book\go-tutorial\code\hello\greetings> go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s) 
=== RUN   TestHelloEmpty        
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      example.com/greetings   0.052s
```

故意编写错误函数

```Go
// 在 greetings/greetings.go 中，粘贴以下代码来代替 Hello 函数。
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    // message := fmt.Sprintf(randomFormat(), name)
    message := fmt.Sprint(randomFormat())
    return message, nil
}
```

测试应不通过:

```Bash
PS D:\projectCode\book\go-tutorial\code\hello\greetings> go test   
--- FAIL: TestHelloName (0.00s)
    greetings_test.go:16: Hello("Gladys") = "Great to see you, %v!", <nil>, want match for `\bGladys\b`, nil
FAIL
exit status 1
FAIL    example.com/greetings   0.037s
PS D:\projectCode\book\go-tutorial\code\hello\greetings> go test -v
=== RUN   TestHelloName
    greetings_test.go:16: Hello("Gladys") = "Hail, %v! Well met!", <nil>, want match for `\bGladys\b`, nil
--- FAIL: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
FAIL
exit status 1
FAIL    example.com/greetings   0.054s
```

## 编译应用程序

```Bash
PS D:\projectCode\book\go-tutorial\code\hello> cd .\hello\

# 编译
PS D:\projectCode\book\go-tutorial\code\hello\hello> go build

# 执行编译文件
PS D:\projectCode\book\go-tutorial\code\hello\hello> .\hello.exe
map[Darrin:Hi, %v. Welcome! Gladys:Great to see you, %v! Samantha:Hail, %v! Well met!]      

# 发现安装路径
PS D:\projectCode\book\go-tutorial\code\hello\hello> go list -f '{{.Target}}'
C:\Users\lushi\go\bin\hello.exe # 可以把C:\Users\lushi\go\bin\放环境中，完成安装
```

## 比较 Go 和 Java
- Go 和 Java 有很多共同之处
  - C 系列 (强类型，括号)
  - 静态类型
  - 垃圾收集
  - 内存安全 (nil 引用，运行时边界检查)
  - 变量总是初始化 (zero/nil/false)
  - 方法
  - 接口
  - 类型断言 (实例)
  - 反射
- Go 与 Java 的不同之处
  - 代码程序直接编译成机器码，没有 VM
  - 静态链接二进制
  - 内存布局控制
  - 函数值和词法闭包
  - 内置字符串 (UTF-8)
  - 内置泛型映射和数组/片段
  - 内置并发
- Go 特意去掉了大量的特性
  - 没有类
  - 没有构造器
  - 没有继承
  - 没有 final
  - 没有异常
  - 没有注解
  - 没有自定义泛型
- 为什么 Go 要省去那些特性？
  - 代码清晰明了是首要的
  - 当查看代码时，可以很清晰的知道程序将会做什么
  - 当编写代码的时候，也可以很清晰的让程序做你想做的
  - 有时候这意味着编写出一个循环而不是调用一个模糊的函数。
  - (不要变的太枯燥)