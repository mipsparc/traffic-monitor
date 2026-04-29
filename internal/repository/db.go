package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"
	"traffic-monitor/internal/config"
	"traffic-monitor/internal/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB There are some better ways to handle. This is for a demo purpose.
var DB *sqlx.DB

func NewDB(conf *config.Config) error {
	var err error
	DB, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName),
	)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	return nil
}

// ReportInsertEntity DTO for inserting a report
type ReportInsertEntity struct {
	CameraID   uint64    `db:"camera_id"`
	UUID       string    `db:"uuid"`
	Time       time.Time `db:"time"`
	VideoID    string    `db:"video_id"`
	Latitude   float64   `db:"latitude"`
	Longitude  float64   `db:"longitude"`
	Severity   uint8     `db:"severity"`
	ReportType string    `db:"report_type"`
	Text       string    `db:"report_text"`
}

func InsertReport(cameraID uint64, report model.Report) error {
	entity := ReportInsertEntity{
		CameraID:   cameraID,
		UUID:       report.UUID,
		Time:       report.Time,
		VideoID:    report.VideoID,
		Latitude:   report.Latitude,
		Longitude:  report.Longitude,
		Severity:   report.Severity,
		ReportType: string(report.ReportType),
		Text:       report.Text,
	}

	res, err := DB.NamedExec(`
		INSERT INTO report
		(camera_id, uuid, time, video_id, latitude, longitude, severity, report_type, report_text)
		VALUES
		(:camera_id, :uuid, :time, :video_id, :latitude, :longitude, :severity, :report_type, :report_text)
		ON DUPLICATE KEY UPDATE uuid = uuid
		`, entity)
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
	query := `SELECT * FROM report`
	var args []interface{}

	if filter != "" {
		query += fmt.Sprintf(" WHERE %s = ? ", filter)
		args = append(args, filterValue)
	}

	if sort == "time" {
		query += fmt.Sprintf(" ORDER BY time %s ", ascDesc)
	} else {
		query += fmt.Sprintf(" ORDER BY %s %s, time DESC ", sort, ascDesc)
	}

	query += ` LIMIT 2000 `

	var reports []SearchResult
	err := DB.Select(&reports, query, args...)
	if err != nil {
		slog.Error("failed to query report: ", err)
		return nil, err
	}

	return &reports, nil
}

func LatestReportID() (uint64, error) {
	var id uint64
	err := DB.QueryRow("SELECT report_id FROM report ORDER BY report_id DESC LIMIT 1").Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		slog.Error("failed to query latest report id: ", err)
		return 0, err
	}
	return id, nil
}
