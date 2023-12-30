# Find and fix vulnerable dependencies with govulncheck

使用 `govulncheck` 查找和修复易受攻击的依赖项。

示例代码 vuln-tutorial ，建议用时 5 分钟。


```Bash
go mod init vuln.tutoria

# 在文件夹中创建一个名为 `main.go` 的文件

go mod tidy

# 观察 go.mod

# 版本降级到 v0.3.5，其中包含漏洞
go get golang.org/x/text@v0.3.5

# 观察 go.mod

# 安装 govulncheck
go install golang.org/x/vuln/cmd/govulncheck@latest

# 分析文件夹
govulncheck ./...

# 可以看到漏洞信息

# 升级到 v0.3.8
go get golang.org/x/text@v0.3.8

# 分析文件夹
govulncheck ./...

# 最后，govulncheck 确认未发现任何漏洞。可以使用命令 govulncheck 定期扫描依赖项
```

```Go
// main.go

package main

import (
        "fmt"
        "os"

        "golang.org/x/text/language"
)

func main() {
    for _, arg := range os.Args[1:] {
        tag, err := language.Parse(arg)
        if err != nil {
            fmt.Printf("%s: error: %v\n", arg, err)
        } else if tag == language.Und {
            fmt.Printf("%s: undefined\n", arg)
        } else {
            fmt.Printf("%s: tag %s\n", arg, tag)
        }
    }
} 
```