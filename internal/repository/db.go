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
	cfg.ParseTime = true

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

// Using ORM is also a good idea.

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

type SearchResult struct {
	ID         uint64     `db:"report_id"`
	CameraID   int        `db:"camera_id"`
	UUID       string     `db:"uuid"`
	Time       *time.Time `db:"time"`
	VideoID    string     `db:"video_id"`
	Latitude   float64    `db:"latitude"`
	Longitude  float64    `db:"longitude"`
	Severity   uint8      `db:"severity"`
	ReportType string     `db:"report_type"`
	Text       string     `db:"report_text"`
}

func SearchReport(filter string, filterValue string, sort string, ascDesc string) (*[]SearchResult, error) {
	var where string
	if filter != "" {
		where = fmt.Sprintf(" WHERE %s = ? ", filter)
	} else {
		where = " WHERE 1 = ? "
		filterValue = "1"
	}

	var orderBy string
	if sort != "" {
		orderBy = fmt.Sprintf(" ORDER BY %s %s ", sort, ascDesc)
	} else {
		orderBy = " ORDER BY report_id DESC "
	}

	rows, err := DB.Query(`SELECT * FROM report`+
		where+
		orderBy+
		`LIMIT 2000`, filterValue)
	if err != nil {
		slog.Error("failed to query report: ", err)
		return nil, fmt.Errorf("failed to query report: %v", err)
	}
	defer rows.Close()

	var reports []SearchResult

	for rows.Next() {
		var report SearchResult
		err = rows.Scan(&report.ID, &report.CameraID, &report.UUID, &report.Time, &report.VideoID, &report.Latitude, &report.Longitude, &report.Severity, &report.ReportType, &report.Text)
		if err != nil {
			slog.Error("failed to scan report: ", err)
			return nil, err
		}
		reports = append(reports, report)
	}

	return &reports, nil
}
