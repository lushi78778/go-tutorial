# Finally

最后再讲一些事情，这个教程到此就基本结束了。

目前你可以完成一些基础的操作，如输出素数等，甚至能开发一些简单的 API 

## 管理依赖关系

你可以通过 Go 工具获取和使用有用的包。在 pkg.go.dev 上，您可以搜索可能找到的软件包，然后使用命令将这些包导入到您自己的代码中调用其函数。

```Bash
# 查找和导入有用的软件包 单击“复制路径”按钮将路径复制到 您的剪贴板。在您自己的代码中，将路径粘贴到 import 语句中
import "rsc.io/quote"

# 在代码中启用依赖项跟踪
go mod init example/mymodule

# 命名模块
<prefix>/<descriptive-text>

# 添加依赖项 （“.”是指当前目录中的包）
go get .
# 添加特定依赖项
go get example.com/theirmodule

# 获取特定依赖项版本
go get example.com/theirmodule@v1.3.4
go get example.com/theirmodule@latest

# 发现可用更新
# 出作为当前模块依赖项的所有模块， 以及每个可用的最新版本
go list -m -u all
# 显示可用于特定模块的最新版本
go list -m -u example.com/theirmodule

# 同步代码的依赖项 该命令没有参数，除了一个打印信息的标志 -v 外
go mod tidy

# 要求使用本地目录中的模块代码
# 请使用 go mod edit 和 go get 命令 要确保文件描述的要求保持一致 使用 replace 指令时，Go 工具不会进行身份验证外部模块
go mod edit -replace example.com/greetings=../greetings
go mod edit -replace=example.com/theirmodule@v0.0.0-unpublished=../theirmodule
go get example.com/theirmodule@v0.0.0-unpublished

# 使用存储库标识符获取特定提交
go get example.com/theirmodule@4cf76c2

# 获取特定分支中的模块
go get example.com/theirmodule@bugfixes

# 删除特定依赖项 指定模块的模块路径并附加 ，如以下示例所示：@none
go get example.com/theirmodule@none

# 停止跟踪所有未使用的模块
go mod tidy

# 指定模块代理服务器
GOPROXY="https://proxy.golang.org,direct"

# 设置环境变量来配置命令 从私有源下载和构建模块
GOPRIVATE=*.corp.example.com,*.research.example.com
```

## 编程语言规范

很长的英文文档 https://go.dev/ref/spec#Return_statements

## 标准库

这是很重要的文档！！！

https://pkg.go.dev/std

## Effective Go

文档提供了编写清晰、惯用的 Go 代码的提示。 

https://go.dev/doc/effective_go
    

 