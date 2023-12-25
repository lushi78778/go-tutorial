# Hello Go

接下来，你将用一小段时间，快速完成go的安装与入门，包括第一个hello world和关于模块的小教程

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

# 定义系统                                         
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




