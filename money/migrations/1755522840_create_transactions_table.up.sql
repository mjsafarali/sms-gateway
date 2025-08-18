CREATE TABLE transactions
(
    id              BIGINT PRIMARY KEY AUTO_INCREMENT,
    company_id      BIGINT      NOT NULL,
    amount          BIGINT      NOT NULL,
    action          VARCHAR(64) NOT NULL,
    ref_type        VARCHAR(64) NOT NULL,
    ref_id          VARCHAR(64) NOT NULL,
    idempotency_key VARCHAR(64) NOT NULL,
    created_at      DATETIME(6) NOT NULL,
    UNIQUE KEY uniq_company_idem (company_id, idempotency_key),
    INDEX           idx_company_created (company_id, created_at),
    FOREIGN KEY (company_id) REFERENCES companies (id)
) ENGINE=InnoDB;