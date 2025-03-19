-- Enable UUID generation extension (only needed once per database)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users Table
CREATE TABLE IF NOT EXISTS users (
                                     user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- Vendor Profiles Table
CREATE TABLE IF NOT EXISTS vendor_profiles (
                                               id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    business_id UUID UNIQUE DEFAULT uuid_generate_v4(),
    business_name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    website VARCHAR(255),
    cac_number VARCHAR(100),
    payment_made BOOLEAN DEFAULT FALSE,
    is_verified BOOLEAN DEFAULT FALSE,
    is_business_verified BOOLEAN DEFAULT FALSE,
    verification_status VARCHAR(50) DEFAULT 'pending', -- New: 'pending', 'approved', 'rejected'
    rejected_reasons TEXT DEFAULT NULL,                -- New: Reason for rejection
    total_sales DECIMAL(12, 2) DEFAULT 0.00,          -- New: Track vendor's total sales
    rating DECIMAL(3,2) DEFAULT 0.00,                 -- New: Vendor rating (1-5 stars)
    review_count INT DEFAULT 0,                      -- New: Number of reviews received
    status VARCHAR(50) DEFAULT 'active',              -- New: 'active', 'suspended', 'deactivated'
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
    );

-- Ensure updated_at column updates automatically on record changes
CREATE OR REPLACE FUNCTION update_timestamp_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger for users table
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'trigger_update_users') THEN
CREATE TRIGGER trigger_update_users
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp_column();
END IF;
END $$;

-- Trigger for vendor_profiles table
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'trigger_update_vendor_profiles') THEN
CREATE TRIGGER trigger_update_vendor_profiles
    BEFORE UPDATE ON vendor_profiles
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp_column();
END IF;
END $$;
