package model

type Reports struct {
	CameraID uint64   `json:"camera_id"`
	Report   []Report `json:"report"`
}

type Report struct {
	ReportUUID string     `json:"report_uuid"`
	TimeStamp  uint64     `json:"timestamp"`
	VideoID    string     `json:"video_id"`
	Latitude   float64    `json:"lat"`
	Longitude  float64    `json:"long"`
	Severity   uint8      `json:"severity"`
	ReportType ReportType `json:"report_type"`
	Reason     string     `json:"reason"`
}
