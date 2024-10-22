USE
password_vault;
CREATE TABLE accounts
(
    id         INT AUTO_INCREMENT PRIMARY KEY,                                 -- Campo id com autoincrement (AUTO_INCREMENT)
    name       VARCHAR(255) NOT NULL,                                          -- Campo name como string (VARCHAR)
    login      VARCHAR(255) NOT NULL,                                          -- Campo login como string (VARCHAR)
    password   VARCHAR(255) NOT NULL,                                          -- Campo password como string (VARCHAR)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                            -- Campo created_at como data
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Campo updated_at como data com atualização automática
);