-- This DSL automatically init the database

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
    report_text VARCHAR(10000) NOT NULL,
    INDEX (camera_id),
    INDEX (time),
    INDEX (report_type)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
