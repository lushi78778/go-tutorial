package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

// 根据名字查询
func albumsByArtist(name string) ([]Album, error) {
    // 声明定义的类型的切片
    var albums []Album
	// 执行语句以查询带有指定的艺术家的名字的记录
	// Query的第一个参数是 SQL 语句。之后 参数，可以传递零个或多个任意类型的参数。减小xss注入措施
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // 循环访问返回的行，使用 Rows.Scan 将每行的列值分配给结构字段。Album
	for rows.Next() {
        var alb Album
		// Scan获取指向 Go 值的指针列表，其中列值将被写入。将指针传递到使用运算符创建的变量中的字段。 通过指针写入以更新 struct 字段
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
		// 在循环中，将新的附加到切片
        albums = append(albums, alb)
    }
	// 循环后，使用 检查整个查询中的错误。请注意，如果查询本身失败，请检查错误 这是发现结果不完整的唯一方法。rows.Err
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

// 查询带有指定ID的 
func albumByID(id int64) (Album, error) {
    var alb Album

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}

// 添加新数据
func addAlbum(alb Album) (int64, error) {
	// 执行INSERT语句
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}


func main() {
    // 配置参数
    cfg := mysql.Config{
        // User:   os.Getenv("DBUSER"),可以借此使用环境变量，构建docker镜像时常用
        User:   "data-access",
        Passwd: "data-accessdata-access",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
    // 获取数据库连接柄
    var err error
    // 调用 sql.Open 初始化变量，传递 的返回值。dbFormatDSN
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        // 为了简化代码，你要调用结束 执行并将错误打印到控制台。在生产代码中，以更优雅的方式处理错误。log.Fatal
        log.Fatal(err)
    }

    // 调用 DB.Ping 到 确认连接到数据库是否正常。
    pingErr := db.Ping()
    
    // 检查错误，以防连接失败。
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    
    // 如果连接成功，则打印一条消息
    fmt.Println("Connected!")


	// 调用您添加的albumsByArtistalbums函数，将其返回值分配给 一个新变量。
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// 调用您添加的albumByID函数。 
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	//使用addAlbum，并返回 ID 添加到变量
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
}