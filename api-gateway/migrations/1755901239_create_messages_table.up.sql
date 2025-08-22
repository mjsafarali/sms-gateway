CREATE TABLE messages
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    company_id BIGINT      NOT NULL,
    receiver   VARCHAR(15) NOT NULL,
    content    TEXT        NOT NULL,
    created_at DATETIME(6) NOT NULL,
    INDEX      idx_company_created (company_id, created_at),
    FOREIGN KEY (company_id) REFERENCES companies (id)
) ENGINE=InnoDB;