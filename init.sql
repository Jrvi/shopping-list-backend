CREATE DATABASE IF NOT EXISTS shopping_list_db;

USE shopping_list_db;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS lists (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INT,
    list_id INT,
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (list_id) REFERENCES lists(id)
);

CREATE USER 'shopping'@'localhost' IDENTIFIED BY 'salainen';

GRANT ALL PRIVILEGES ON shopping_list_db.* TO 'shopping'@'localhost';

FLUSH PRIVILEGES;