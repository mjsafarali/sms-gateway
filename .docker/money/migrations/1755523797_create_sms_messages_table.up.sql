CREATE TABLE sms_message
(
    id                   BIGINT PRIMARY KEY AUTO_INCREMENT,
    company_id           BIGINT        NOT NULL,
    request_id           CHAR(26)      NOT NULL, -- ULID/KSUID for traceability
    idem_key             VARCHAR(64)   NOT NULL, -- client-provided idempotency key
    to_msisdn            VARCHAR(20)   NOT NULL, -- E.164
    body                 VARCHAR(1024) NOT NULL, -- "all sms's are in one sms"
    encoding             ENUM('GSM7','UCS2') NOT NULL,
    status               ENUM('accepted','queued','sent','delivered','failed') NOT NULL,
    provider             VARCHAR(64)  DEFAULT NULL,
    provider_msg_id      VARCHAR(128) DEFAULT NULL,
    price_reserved_minor BIGINT        NOT NULL,
    price_charged_minor  BIGINT       DEFAULT 0,
    error_code           VARCHAR(64)  DEFAULT NULL,
    created_at           DATETIME(6) NOT NULL,
    queued_at            DATETIME(6) DEFAULT NULL,
    sent_at              DATETIME(6) DEFAULT NULL,
    final_at             DATETIME(6) DEFAULT NULL,
    UNIQUE KEY uniq_company_idem (company_id, idem_key),
    INDEX                idx_company_created (company_id, created_at),
    INDEX                idx_company_status (company_id, status, created_at),
    FOREIGN KEY (company_id) REFERENCES company (id)
) ENGINE=InnoDB;