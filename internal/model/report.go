package model

import "time"

type Reports struct {
	CameraID uint64   `json:"camera_id" validate:"required"`
	Report   []Report `json:"report" validate:"required,dive"`
}

type Report struct {
	UUID       string     `json:"uuid" validate:"required"`
	Time       time.Time  `json:"time" validate:"required"`
	VideoID    string     `json:"video_id" validate:"required"`
	Latitude   float64    `json:"lat" validate:"required"`
	Longitude  float64    `json:"long" validate:"required"`
	Severity   uint8      `json:"severity" validate:"required"`
	ReportType ReportType `json:"report_type" validate:"required"`
	Text       string     `json:"text" validate:"required"`
}
