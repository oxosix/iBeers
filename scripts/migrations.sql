-- migrations.sql

-- Migração para a tabela beer_type
CREATE TABLE IF NOT EXISTS beer_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Migração para a tabela beer_style
CREATE TABLE IF NOT EXISTS beer_style (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Migração para a tabela beer
CREATE TABLE IF NOT EXISTS beer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type_id INT REFERENCES beer_type(id),
    style_id INT REFERENCES beer_style(id)
);