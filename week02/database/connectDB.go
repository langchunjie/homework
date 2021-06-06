package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = "root"
	ip       = "localhost"
	port     = "3306"
	dbName   = "class" //数据库名称
)

func GetDataBase() *sql.DB {
	//mysql 数据库
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println(path)
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ := sql.Open("mysql", path)
	if DB == nil {
		log.Fatal("连接失败！")
		return nil
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(time.Minute * 3)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(10)

	//验证连接
	if err := DB.Ping(); err != nil {
		log.Fatal("open database fail")
		return nil
	}
	return DB
}
