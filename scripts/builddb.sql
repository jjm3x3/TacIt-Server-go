-- CREATE USER gorm;
-- CREATE DATABASE tacit_db;

-- asumes localhost access
CREATE USER 'gorm'@'localhost';
GRANT ALL PRIVILEGES ON *.* TO 'gorm'@'localhost';
CREATE DATABASE tacit_db;
-- ADD PASSWORD TO SCRIPT