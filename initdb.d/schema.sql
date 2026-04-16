CREATE TABLE IF NOT EXISTS report (
    report_id BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    camera_id INT NOT NULL,
    uuid VARCHAR(100) CHARSET ascii NOT NULL UNIQUE,
    time DATETIME NOT NULL,
    video_id VARCHAR(100) CHARSET ascii NOT NULL,
    latitude DOUBLE NOT NULL,
    longitude DOUBLE NOT NULL,
    severity TINYINT NOT NULL,
    report_type VARCHAR(100) CHARSET ascii NOT NULL,
    report_text TEXT NOT NULL,
    INDEX (camera_id),
    INDEX (time),
    INDEX (report_type)
);
