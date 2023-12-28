# Connect Mysql

本教程的示例代码为 data-access

简单介绍一下数据库连接，建议用时二十分钟。

## 准备工作

执行以下sql语句

```SQL
create database recordings;

use recordings;

DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
```

```SQL
/*
 Navicat Premium Data Transfer

 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : recordings

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 28/12/2023 16:18:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for album
-- ----------------------------
DROP TABLE IF EXISTS `album`;
CREATE TABLE `album`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `artist` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` decimal(5, 2) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of album
-- ----------------------------
INSERT INTO `album` VALUES (1, 'Blue Train', 'John Coltrane', 56.99);
INSERT INTO `album` VALUES (2, 'Giant Steps', 'John Coltrane', 63.99);
INSERT INTO `album` VALUES (3, 'Jeru', 'Gerry Mulligan', 17.99);
INSERT INTO `album` VALUES (4, 'Sarah Vaughan', 'Sarah Vaughan', 34.98);
INSERT INTO `album` VALUES (5, 'The Modern Sound of Betty Carter', 'Betty Carter', 49.99);

SET FOREIGN_KEY_CHECKS = 1;

```

建好表后，保存好自己的连接配置

```Go
// 我的配置
cfg := mysql.Config{
    User:   "data-access",
    Passwd: "data-accessdata-access",
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "recordings",
}
```

## 访问数据库

### 前置惯常操作

```Bash
mkdir data-access
cd data-access

# 初始化项目
go mod init example/data-access

```

### 测试sql连接

```Go
// main.go 

package main

import "github.com/go-sql-driver/mysql"

// 声明一个 *sql 类型的变量。数据库。 这是您的数据库句柄
var db *sql.DB

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
}
```

### 查询多行

```Go
// main.go 紧挨着上面加上代码

// 定义结构
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

//============================================================================
// main func
//============================================================================
// 调用您添加的albumsByArtistalbums函数，将其返回值分配给 一个新变量。
albums, err := albumsByArtist("John Coltrane")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Albums found: %v\n", albums)
```

### 查询单行

```Go
// main.go 紧挨着上面加上代码
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

//============================================================================
// main func
//============================================================================
// 调用您添加的albumByID函数。 
alb, err := albumByID(2)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Album found: %v\n", alb)
```

### 添加数据

```Go
// main.go 紧挨着上面加上代码
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

//============================================================================
// main func
//============================================================================
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
```

### 输出样例

```Bash
PS D:\projectCode\book\go-tutorial\code\data-access> go run .
Connected!
Albums found: [{1 Blue Train John Coltrane 56.99} {2 Giant Steps John Coltrane 63.99}]
Album found: {2 Giant Steps John Coltrane 63.99}
ID of added album: 6
```

## 常用 Drivers 指引

Drivers for Go’s sql package include https://go.dev/wiki/SQLDrivers

- Amazon AWS Athena: https://github.com/uber/athenadriver
- AWS Athena: https://github.com/segmentio/go-athena
- AWS DynamoDB: https://github.com/btnguyen2k/godynamo
- Apache Avatica/Phoenix: https://github.com/apache/calcite-avatica-go
- Apache H2: https://github.com/jmrobles/h2go
- Apache Hive: https://github.com/sql-machine-learning/gohive
- Apache Ignite/GridGain: https://github.com/amsokol/ignite-go-client
- Apache Impala: https://github.com/bippio/go-impala
- Azure Cosmos DB: https://github.com/btnguyen2k/gocosmos
- ClickHouse (uses HTTP API): https://github.com/mailru/go-clickhouse
- ClickHouse (uses native TCP interface): https://github.com/ClickHouse/clickhouse-go
- CockroachDB: Use any PostgreSQL driver
- Couchbase N1QL: https://github.com/couchbase/go_n1ql
- DB2 LUW (uses cgo): https://github.com/asifjalil/cli
- DB2 LUW and DB2/Z with DB2-Connect: https://bitbucket.org/phiggins/db2cli (Last updated 2015-08)
- DB2 LUW, z/OS, iSeries and Informix: https://github.com/ibmdb/go_ibm_db
- Databricks: https://github.com/databricks/databricks-sql-go
- DuckDB: https://github.com/marcboeker/go-duckdb
- Exasol: (pure Go): https://github.com/exasol/exasol-driver-go
- Firebird SQL: https://github.com/nakagami/firebirdsql
- Genji (pure go): https://github.com/genjidb/genji
- Google Cloud BigQuery: https://github.com/solcates/go-sql-bigquery
- Google Cloud Spanner: https://github.com/googleapis/go-sql-spanner
- Google Cloud Spanner: https://github.com/rakyll/go-sql-driver-spanner
- MS ADODB: https://github.com/mattn/go-adodb
- MS SQL Server (pure go): https://github.com/microsoft/go-mssqldb
- MS SQL Server (uses cgo): https://github.com/minus5/gofreetds
- MaxCompute: https://github.com/sql-machine-learning/gomaxcompute
- MySQL: https://github.com/go-sql-driver/mysql/ [*]
- MySQL: https://github.com/siddontang/go-mysql/ [**] (also handles replication)
- MySQL: https://github.com/ziutek/mymysql [*]
- ODBC: https://bitbucket.org/miquella/mgodbc (Last updated 2016-02)
- ODBC: https://github.com/alexbrainman/odbc
- Oracle (pure go): https://github.com/sijms/go-ora
- Oracle (uses cgo): https://github.com/godror/godror
- Oracle (uses cgo): https://github.com/mattn/go-oci8
- Oracle (uses cgo): https://gopkg.in/rana/ora.v4
- Postgres (pure Go): https://github.com/jackc/pgx [*]
- Postgres (pure Go): https://github.com/lib/pq [*]
- Postgres (uses cgo): https://github.com/jbarham/gopgsqldriver
- Presto: https://github.com/prestodb/presto-go-client
- QL: https://pkg.go.dev/modernc.org/ql
- SAP ASE (pure go): https://github.com/SAP/go-ase
- SAP ASE (uses cgo): https://github.com/SAP/cgo-ase
- SAP HANA (pure go): https://github.com/SAP/go-hdb
- SAP HANA (uses cgo): https://help.sap.com/viewer/0eec0d68141541d1b07893a39944924e/2.0.03/en-US/0ffbe86c9d9f44338441829c6bee15e6.html
- SQL over REST: https://github.com/adaptant-labs/go-sql-rest-driver
- SQLite (uses cgo): https://github.com/gwenn/gosqlite - Supports SQLite dynamic data typing
- SQLite (uses cgo): https://github.com/mattn/go-sqlite3 [*]
- SQLite (uses cgo): https://github.com/mxk/go-sqlite
- SQLite: (pure go): https://modernc.org/sqlite
- SQLite: (uses cgo): https://github.com/rsc/sqlite
- SingleStore: Use any MySQL driver
- Snowflake (pure Go): https://github.com/snowflakedb/gosnowflake
- Sybase ASE (pure go): https://github.com/thda/tds
- Sybase SQL Anywhere: https://github.com/a-palchikov/sqlago
- TiDB: Use any MySQL driver
- Trino: https://github.com/trinodb/trino-go-client
- Vertica: https://github.com/vertica/vertica-sql-go
- Vitess: https://pkg.go.dev/vitess.io/vitess/go/vt/vitessdriver
- YDB (pure go): https://github.com/ydb-platform/ydb-go-sdk
- YQL (Yahoo! Query Language): https://github.com/mattn/go-yql

Drivers marked with [*] are both included in and pass the compatibility test suite at https://github.com/bradfitz/go-sql-test. Drivers marked with [**] pass the compatibility test suite but are not currently included in it.

