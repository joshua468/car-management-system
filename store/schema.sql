-- Enable the pgcrypto extension to generate UUIDs
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Drop existing tables if they exist to allow fresh setup
DROP TABLE IF EXISTS engine CASCADE;
DROP TABLE IF EXISTS car CASCADE;

-- Create table for engines
CREATE TABLE engine (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    displacement NUMERIC NOT NULL,
    no_of_cylinders INTEGER NOT NULL,
    car_range VARCHAR(255) NOT NULL
);

-- Create table for cars
CREATE TABLE car (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    brand VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    engine_id UUID,
    FOREIGN KEY (engine_id) REFERENCES engine (id) ON DELETE SET NULL
);

-- Seed initial data for testing (optional)
INSERT INTO engine (displacement, no_of_cylinders, car_range)
VALUES
    (2.0, 4, 'Sedan'),
    (3.5, 6, 'SUV');

INSERT INTO car (brand, model, year, engine_id)
VALUES
    ('Toyota', 'Camry', 2020, (SELECT id FROM engine WHERE car_range = 'Sedan' LIMIT 1)),
    ('Honda', 'Civic', 2019, (SELECT id FROM engine WHERE car_range = 'SUV' LIMIT 1));
