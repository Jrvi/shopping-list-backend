-- Käyttäjätaulu
CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- Kategoriataulu
CREATE TABLE IF NOT EXISTS category (
    id INTEGER PRIMARY KEY  AUTOINCREMENT,
    title VARCHAR(255) NOT NULL
);

-- Listataulu
CREATE TABLE IF NOT EXISTS list (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    is_shared INT NOT NULL DEFAULT 0,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

-- Jaettu listataulu
CREATE TABLE list_shares (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    list_id INT NOT NULL,
    shared_with_user_id INT NOT NULL,
    FOREIGN KEY (list_id) REFERENCES list(id)
);

-- Tuotetaulu
CREATE TABLE IF NOT EXISTS product (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    category_id INT,
    list_id INT,
    FOREIGN KEY (category_id) REFERENCES category(id),
    FOREIGN KEY (list_id) REFERENCES list(id)
);