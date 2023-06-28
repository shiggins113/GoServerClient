CREATE DATABASE IF NOT EXISTS certs;
USE certs; 
CREATE TABLE IF NOT EXISTS certs (
    certID int NOT NULL AUTO_INCREMENT,
    cert varchar(255),
    beforeDate varchar(255),
    afterDate varchar(255),
    PRIMARY KEY (certID)
);