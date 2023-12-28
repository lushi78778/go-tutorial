# RESTful API via Gin

本教程的示例代码为 web-service-gin

简单介绍一下使用 Gin 路由请求，检索请求详细信息，并封送 JSON 以获取响应。建议用时十分钟。

## 惯常操作

```Bash
mkdir web-service-gin
cd web-service-gin

# 初始化项目
go mod init example/web-service-gin
```

## 查

```Go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // 定义 get /albums 路由
    router.GET("/albums", getAlbums)
    // 定义服务 endpoint
    router.Run("localhost:8080")
}

// albums 转 json 返回请求
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// 数据示例（应当来自数据库）
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// 定义结构体
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
```

应当同步依赖 `go mod tidy` 后运行 `go run .`

## 条件查
```Go
// 根据 id 查询 
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

//============================================================================
// main func
//============================================================================
//添加到 router.Run("localhost:8080") 之前
router.GET("/albums/:id", getAlbumByID)
```

## 增
```Go
// post 增加数据
func postAlbums(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

//============================================================================
// main func
//============================================================================
//添加到 router.Run("localhost:8080") 之前
router.POST("/albums", postAlbums)
```

参考命令 (建议用 postman 测试)

```Bash
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```