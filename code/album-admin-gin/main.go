package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Album 结构体表示数据库中的 album 表的行
type Album struct {
	ID     int     `json:"id" example:"1"`
	Title  string  `json:"title" example:"Album Title"`
	Artist string  `json:"artist" example:"Artist Name"`
	Price  float64 `json:"price" example:"19.99"`
}

var db *sql.DB // 全局数据库连接

func main() {
	// 初始化数据库
	initDB()

	// 创建 Gin 引擎
	r := gin.Default()

	// 注册业务路由
	registerRoutes(r)

	// 启动服务器
	r.Run("localhost:800")
}

// initDB 初始化 SQLite 数据库连接
func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}

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
}

// registerRoutes 注册业务路由
func registerRoutes(r *gin.Engine) {


	// Album 相关路由
	albumGroup := r.Group("/api/albums")
	{
		// @Summary 获取所有的 album
		// @Produce json
		// @Success 200 {array} Album
		// @Router /api/albums [get]
		albumGroup.GET("", GetAlbums)


		// @Summary 获取特定的 album
		// @Produce json
		// @Param id path int true "Album ID"
		// @Success 200 {object} Album
		// @Router /api/albums/{id} [get]
		albumGroup.GET("/:id", GetAlbum)

		// @Summary 添加新的 album
		// @Accept json
		// @Produce json
		// @Param album body Album true "Album object"
		// @Success 200 {object} Album
		// @Router /api/albums [post]
		albumGroup.POST("", AddAlbum)

		// @Summary 更新特定的 album
		// @Accept json
		// @Produce json
		// @Param id path int true "Album ID"
		// @Param album body Album true "Album object"
		// @Success 200 {object} Album
		// @Router /api/albums/{id} [put]
		albumGroup.PUT("/:id", UpdateAlbum)

		// @Summary 删除特定的 album
		// @Param id path int true "Album ID"
		// @Success 200 {object} gin.H
		// @Router /api/albums/{id} [delete]
		albumGroup.DELETE("/:id", DeleteAlbum)
	}
}


// GetAlbums 返回所有的 album
func GetAlbums(c *gin.Context) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var album Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		albums = append(albums, album)
	}

	c.JSON(200, albums)
}

// GetAlbum 根据ID返回特定的 album
func GetAlbum(c *gin.Context) {
	id := c.Param("id")
	var album Album
	err := db.QueryRow("SELECT * FROM album WHERE id=?", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, album)
}

// AddAlbum 添加一个新的 album
func AddAlbum(c *gin.Context) {
	var album Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	album.ID = int(id)

	c.JSON(200, album)
}

// UpdateAlbum 根据ID更新特定的 album
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE album SET title=?, artist=?, price=? WHERE id=?", album.Title, album.Artist, album.Price, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, album)
}

// DeleteAlbum 根据ID删除特定的 album
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM album WHERE id=?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Album deleted successfully"})
}

