package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //mysql
	"github.com/jinzhu/gorm"
	"time"
)

// NewDBConnection 新規データベースコネクションを取得します.
func NewDBConnection() *gorm.DB {
	return getMysqlConn()
}

func getMysqlConn() *gorm.DB {
	// DBMS := "mysql"
	USER := Current.Database.User
	PASS := Current.Database.Password
	PROTOCOL := "tcp(mysql:" + Current.Database.Port + ")"
	DBNAME := Current.Database.Database
	connectionString := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	conn, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = conn.DB().Ping()
	if err != nil {
		panic(err)
	}

	conn.LogMode(true)
	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(20)

	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := Current.Database.User
	PASS := Current.Database.Password
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := Current.Database.Database

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 10 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				fmt.Println(err.Error())
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	return db
}
