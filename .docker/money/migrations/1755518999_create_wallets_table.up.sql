CREATE TABLE wallets
(
    company_id BIGINT PRIMARY KEY,
    balance    BIGINT NOT NULL,
    version    BIGINT NOT NULL DEFAULT 0,
    updated_at DATETIME(6) NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies (id)
) ENGINE=InnoDB;