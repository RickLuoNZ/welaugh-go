
// TODO: change it to a shell script
// MySQL script to create the DATABASE for dev/testing purpose

CREATE DATABASE test;

CREATE USER 'test'@'localhost' IDENTIFIED BY 'test';

GRANT ALL PRIVILEGES ON test.* TO 'test'@'localhost';

FLUSH PRIVILEGES;