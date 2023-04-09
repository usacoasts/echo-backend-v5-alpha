package conf

import "os"

// Config conf struct
type Config struct {
	Server struct {
		Port int
		Mock bool
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

// Current runnnig configuration
var Current Config

// NewConfig プロジェクトのコンフィグ設定をロードします.
func NewConfig() {
	Current = Config{
		Server: struct {
			Port int
			Mock bool
		}{Port: 8080, Mock: false},
		Database: struct {
			Host     string
			Port     string
			User     string
			Password string
			Database string
		}{Host: "mysql", Port: "3306", User: os.Getenv("MYSQL_USER"), Password: os.Getenv("MYSQL_PASSWORD"), Database: os.Getenv("MYSQL_DATABASE")},
	}

	return
}
