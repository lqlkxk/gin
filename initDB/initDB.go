package initDB

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var Db *sqlx.DB

func init() {
	// 官方
	//sql.Open("mysql","root@root:tcp(127.0.0.1:3306/db1)")
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db1")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	database.SetMaxOpenConns(20)
	database.SetMaxIdleConns(20)
	database.SetConnMaxLifetime(time.Second * 15)
	Db = database
}
