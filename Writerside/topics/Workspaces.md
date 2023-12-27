# Workspaces

本教程的示例代码为 workspace

简单介绍一下工作区的概念与使用，建议用时十分钟。

```Go
# 创建 hello 模块
go mod init example.com/hello

# 添加对 golang.org/x/example/hello/reverse 包的依赖关系。go get
go get golang.org/x/example/hello/reverse
```

```Go
// hello/hello.go

package main

import (
    "fmt"

    "golang.org/x/example/hello/reverse"
)

func main() {
    fmt.Println(reverse.String("Hello"))
}
```

```Go
# 初始化工作区
go work init ./hello
# 该命令生成文件：go.work 该文件的语法类似go.mod

# 在工作区目录下运行程序
go run ./hello
```

- 在工作区目录中,克隆存储库 `git clone https://go.googlesource.com/example`
- 将模块添加到工作区 `go work use ./example/hello`

```Go
// 创建 workspace/example/hello/reverse/int.go

package reverse

import "strconv"

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
    i, _ = strconv.Atoi(String(strconv.Itoa(i)))
    return i
}
```

```Go
// 修改hello/hello.go

package main

import (
    "fmt"

    "golang.org/x/example/hello/reverse"
)

func main() {
    fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
```

- 这两个模块位于同一个工作区中，因此很容易在一个模块中进行更改并在另一个模块中使用它