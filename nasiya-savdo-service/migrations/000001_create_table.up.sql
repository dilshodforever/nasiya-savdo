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
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at BIGINT
);

CREATE TABLE storage (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE product (
  id UUID PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  color VARCHAR(50),
  model VARCHAR(50),
  image_url TEXT,
  made_in VARCHAR(100),
  date_of_creation TIMESTAMPTZ,
  storage_id UUID REFERENCES storage(id) ON DELETE SET NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at BIGINT
);

CREATE TABLE contract (
  id UUID PRIMARY KEY,
  consumer_name VARCHAR(255) NOT NULL,
  consumer_passport_serial VARCHAR(50),
  consumer_address TEXT,
  passport_image TEXT,
  status contract_status,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  duration INT,
  deleted_at TIMESTAMPTZ
);

CREATE TABLE exchange (
  id UUID PRIMARY KEY,
  product_id UUID REFERENCES product(id) ON DELETE SET NULL,
  amount INT NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  status exchange_status,
  contract_id UUID REFERENCES contract(id) ON DELETE SET NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at BIGINT
);

CREATE TABLE transaction (
  id UUID PRIMARY KEY,
  contract_id UUID REFERENCES contract(id) ON DELETE CASCADE,
  price DECIMAL(10, 2) NOT NULL,
  duration INT,  -- Payment duration
  created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE notification (
  id UUID PRIMARY KEY,
  message TEXT NOT NULL,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ DEFAULT NOW()
);
