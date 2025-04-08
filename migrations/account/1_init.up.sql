-- Создание таблицы приложений (apps)
CREATE TABLE apps
(
    id INT IDENTITY(1,1) PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    secret VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE AuthWebAdmin
(
    id INT IDENTITY(1,1) PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL UNIQUE
);
