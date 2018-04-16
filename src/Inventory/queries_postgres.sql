-- Database: test

-- DROP DATABASE test;

CREATE DATABASE test
  WITH OWNER = postgres
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       LC_COLLATE = 'en_US.UTF-8'
       LC_CTYPE = 'en_US.UTF-8'
       CONNECTION LIMIT = -1;

CREATE TABLE ProductType(
    Id INTEGER PRIMARY KEY NOT NULL,
    Description VARCHAR(100) NOT NULL
);

CREATE TABLE Product(
    Id INTEGER PRIMARY KEY NOT NULL,
    Description VARCHAR(150) NOT NULL,
    Price NUMERIC(10,2) NOT NULL,
    Type INTEGER REFERENCES ProductType(Id) NOT NULL,
    Quantity NUMERIC NOT NULL
);

INSERT INTO PRODUCTTYPE (id, description) VALUES (1, 'Producto'), (2, 'Servicio');