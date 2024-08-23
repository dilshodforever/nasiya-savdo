CREATE TYPE role_user AS ENUM ('admin', 'user');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL,
    full_name VARCHAR(150),
    native_language VARCHAR(10),
    role role_user DEFAULT 'user',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);
