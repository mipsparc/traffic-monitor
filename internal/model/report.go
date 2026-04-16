package model

import "time"

type Reports struct {
	CameraID uint64   `json:"camera_id"`
	Report   []Report `json:"report"`
}

type Report struct {
	UUID       string     `json:"uuid"`
	Time       time.Time  `json:"time"`
	VideoID    string     `json:"video_id"`
	Latitude   float64    `json:"lat"`
	Longitude  float64    `json:"long"`
	Severity   uint8      `json:"severity"`
	ReportType ReportType `json:"report_type"`
	Describe   string     `json:"describe"`
}
