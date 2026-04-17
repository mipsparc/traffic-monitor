package repository

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
	"time"
	"traffic-monitor/internal/config"
	"traffic-monitor/internal/model"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// DB There are some better ways to handle. This is for a demo purpose.
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

	return nil
}

func InsertReport(cameraID uint64, report model.Report) error {
	res, err := DB.Exec(strings.ReplaceAll(`
		INSERT INTO report
		(camera_id, uuid, time, video_id, latitude, longitude, severity, report_type, report_text)
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE ""uuid"" = ""uuid""
		`, `""`, "`"), cameraID, report.UUID, report.Time, report.VideoID, report.Latitude, report.Longitude, report.Severity, report.ReportType, report.Text,
	)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		slog.Error("no rows affected after inserting report")
	}

	return nil
}
