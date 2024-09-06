-- Create ENUM types
CREATE TYPE contract_status AS ENUM ('pending', 'finished', 'canceled');
CREATE TYPE exchange_status AS ENUM ('buy', 'sell');

-- Create tables
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
    consumer_name VARCHAR(255) NOT NULL,
    consumer_passport_serial VARCHAR(50),
    consumer_address TEXT,
    consumer_phone_number VARCHAR(13),
    passport_image TEXT,
    status contract_status,
    created_at TIMESTAMP DEFAULT NOW(),
    duration INT,
    deleted_at TIMESTAMP
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
