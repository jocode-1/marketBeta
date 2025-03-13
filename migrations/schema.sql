DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_id UUID DEFAULT gen_random_uuid() UNIQUE,
                       username VARCHAR(100) NOT NULL UNIQUE,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       hashed_password TEXT NOT NULL,
                       phone_number VARCHAR(20),
                       user_address TEXT,
                       profile_photo_url TEXT,
                       ip_address VARCHAR(50),
                       is_verified BOOLEAN DEFAULT FALSE,
                       is_admin BOOLEAN DEFAULT FALSE,
                       is_vendor BOOLEAN DEFAULT FALSE,
                       role VARCHAR(50) DEFAULT 'user',
                       status BOOLEAN DEFAULT TRUE,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
