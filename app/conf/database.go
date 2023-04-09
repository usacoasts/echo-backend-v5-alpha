package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDBConnection 新規データベースコネクションを取得します.
func NewDBConnection() *gorm.DB {
	return getMysqlConn()
}

func getMysqlConn() *gorm.DB {
	USER := Current.Database.User
	PASS := Current.Database.Password
	DBHOST := Current.Database.Host
	DBNAME := Current.Database.Database

	PROTOCOL := "tcp(" + DBHOST + ":" + Current.Database.Port + ")"

	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	print(DBHOST)
	print(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	mysqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = mysqlDB.Ping()
	if err != nil {
		panic(err)
	}

	// err = mysqlDB.SetMaxIdleConns(10)
	// err = mysqlDB.SetMaxOpenConns(20)

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	return db
}
