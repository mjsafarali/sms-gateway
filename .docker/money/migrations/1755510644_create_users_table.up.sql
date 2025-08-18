CREATE TABLE companies
(
    id            BIGINT PRIMARY KEY AUTO_INCREMENT,
    name          VARCHAR(255) NOT NULL,
    price_per_sms BIGINT       NOT NULL,
    daily_quota   BIGINT       NOT NULL,
    rps_limit     INT          NOT NULL DEFAULT 0,
    is_active     TINYINT(1) NOT NULL DEFAULT 1,
    created_at    DATETIME(6) NOT NULL,
    updated_at    DATETIME(6) NOT NULL
) ENGINE=InnoDB;