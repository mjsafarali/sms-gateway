CREATE TABLE transactions
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    company_id BIGINT      NOT NULL,
    amount     BIGINT      NOT NULL,
    action     VARCHAR(64) NOT NULL,
    balance    BIGINT      NOT NULL,
    created_at DATETIME(6) NOT NULL,
    INDEX      idx_company_created (company_id, created_at),
    FOREIGN KEY (company_id) REFERENCES companies (id)
) ENGINE=InnoDB;