# Task: Album Admin System

本教程的示例代码 album-admin-gin

接下来，我们将使用所学知识，来构建一个Album Admin System API，实现相册信息的管理操作API构建。（如果有 Java Web 开发经历，请独立完成。）

代码很简单，请自行看注释 https://github.com/lushi78778/go-tutorial/blob/main/code/album-admin-gin/main.go

## 调试运行

安装gcc https://jmeubank.github.io/tdm-gcc/download/

![gcc.png](gcc.png)

```Go
go mod tidy

# 安装个 gcc 这里 sqlite 驱动是使用 CGO 的
$env:CGO_ENABLED=1

go run .
```

## 接口在线调试文档

https://apifox.com/apidoc/shared-2ae3ecce-64a6-4ba7-b41a-1570dbd9a941/api-137769561
