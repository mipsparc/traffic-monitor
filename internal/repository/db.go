package repository

import (
	"database/sql"
	"fmt"
	"time"
	"traffic-monitor/internal/config"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// There are better way to handle this. This is for demo purpose.
var DB *sql.DB

func NewDB(conf *config.Config) error {
	cfg := mysql.NewConfig()
	cfg.User = conf.DBUser
	cfg.Passwd = conf.DBPassword
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", conf.DBHost, conf.DBPort)
	cfg.DBName = conf.DBName

	conn, err := mysql.NewConnector(cfg)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	DB = sql.OpenDB(conn)
	DB.SetMaxOpenConns(30)
	DB.SetMaxIdleConns(30)
	DB.SetConnMaxLifetime(5 * time.Minute)

	if err := DB.Ping(); err != nil {
		// ignore error on close
		_ = DB.Close()
		return fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to database")

	return nil
}
