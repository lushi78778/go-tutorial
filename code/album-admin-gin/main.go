// main.go
package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Album 结构体
type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

// EditData 结构体
type EditData struct {
	Album Album
}

func main() {
	// 连接SQLite数据库
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	// 创建 album 表
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS album (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			artist TEXT,
			price REAL
		)
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Error creating table:", err)
		os.Exit(1)
	}

	// 创建 Gin 引擎
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		// 查询所有 album 记录
		rows, err := db.Query("SELECT id, title, artist, price FROM album")
		if err != nil {
			fmt.Println("Error querying database:", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer rows.Close()

		var albums []Album
		for rows.Next() {
			var album Album
			err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
			if err != nil {
				fmt.Println("Error scanning row:", err)
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			albums = append(albums, album)
		}

		data := gin.H{
			"Albums": albums,
		}

		c.HTML(http.StatusOK, "index.html", data)
	})

	r.GET("/add", func(c *gin.Context) {
		// 渲染添加专辑的表单页面
		data := EditData{
			Album: Album{},
		}
		c.HTML(http.StatusOK, "form.html", data)
	})

	r.POST("/add", func(c *gin.Context) {
		// 处理添加专辑的表单提交
		var album Album
		if err := c.ShouldBind(&album); err != nil {
			fmt.Println("Error binding form data:", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// 执行插入操作
		_, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
		if err != nil {
			fmt.Println("Error inserting data:", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 重定向到首页或显示成功消息
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/edit/:id", func(c *gin.Context) {
		// 渲染编辑专辑的表单页面
		id := c.Param("id")

		// 查询数据库获取专辑信息
		var album Album
		err := db.QueryRow("SELECT id, title, artist, price FROM album WHERE id = ?", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			fmt.Println("Error querying data:", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 渲染表单页面
		data := EditData{
			Album: album,
		}
		c.HTML(http.StatusOK, "form.html", data)
	})

	r.POST("/edit/:id", func(c *gin.Context) {
		// 处理编辑专辑的表单提交
		id := c.Param("id")

		var album Album
		if err := c.ShouldBind(&album); err != nil {
			fmt.Println("Error binding form data:", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// 执行更新操作
		_, err := db.Exec("UPDATE album SET title=?, artist=?, price=? WHERE id=?", album.Title, album.Artist, album.Price, id)
		if err != nil {
			fmt.Println("Error updating data:", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 重定向到首页或显示成功消息
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.POST("/delete/:id", func(c *gin.Context) {
		// 处理删除专辑的请求
		id := c.Param("id")

		// 执行删除操作
		_, err := db.Exec("DELETE FROM album WHERE id=?", id)
		if err != nil {
			fmt.Println("Error deleting data:", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 重定向到首页或显示成功消息
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":8080")
}
