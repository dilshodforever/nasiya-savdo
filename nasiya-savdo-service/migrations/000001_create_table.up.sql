-- Create ENUM types
CREATE TYPE contract_status AS ENUM ('pending', 'finished', 'canceled');
CREATE TYPE exchange_status AS ENUM ('buy', 'sell');

-- Create tables
CREATE TABLE users (
    id UUID PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    address TEXT,
    phone_number VARCHAR(20),
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE storage (
    id UUID PRIMARY KEY,
    user_id UUID,
    name VARCHAR(100) NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE products (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    color VARCHAR(50),
    model VARCHAR(50),
    image_url TEXT,
    made_in VARCHAR(100),
    date_of_creation TIMESTAMP,
    storage_id UUID REFERENCES storage(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE contract (
    id UUID PRIMARY KEY,
    storage_id UUID REFERENCES storage(id) ON DELETE SET NULL, 
    consumer_name VARCHAR(255) NOT NULL,
    consumer_passport_serial VARCHAR(50),
    consumer_address TEXT,
    consumer_phone_number VARCHAR(13),
    passport_image TEXT,
    status contract_status,
    created_at TIMESTAMP DEFAULT NOW(),
    duration INT,
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);


CREATE TABLE exchange (
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES products(id) ON DELETE SET NULL,
    amount INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    status exchange_status,
    contract_id UUID REFERENCES contract(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    contract_id UUID REFERENCES contract(id) ON DELETE CASCADE,
    price DECIMAL(10, 2) NOT NULL,
    duration INT,  -- Payment duration
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE notification (
    id UUID PRIMARY KEY,
    message TEXT NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW()
);





INSERT INTO products (
    id,
    name,
    color,
    model,
    image_url,
    made_in,
    date_of_creation,
    storage_id,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    '4bc42fb2-6e17-495a-b40f-7bceb5801940',      -- Generates a new UUID for the id
    'Sample Product',        -- Example product name
    'Red',                   -- Color
    'Model X',               -- Model
    'http://example.com/image.jpg', -- Image URL
    'USA',                   -- Made in location
    NOW(),                   -- Date of creation
    '2cb30b45-e877-4872-b8ff-d9aabc33cb69',                    -- Storage ID (set to NULL if no related storage exists)
    NOW(),                   -- Created at (defaults to current timestamp)
    NOW(),                   -- Updated at (defaults to current timestamp)
    0                        -- Deleted at (0 indicates not deleted)
);