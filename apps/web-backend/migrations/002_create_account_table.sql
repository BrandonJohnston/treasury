-- Migration: 002_create_account_table.sql
-- Description: Creates the accounts table for the Treasury application
-- Created: 2025-09-28

-- Drop table if exists (for reset functionality)
DROP TABLE IF EXISTS accounts CASCADE;

-- Create accounts table
CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_name VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    total_balance DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- Create index on email for faster lookups
CREATE INDEX idx_accounts_user_id ON accounts(user_id);

-- Create trigger to automatically update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_accounts_updated_at 
    BEFORE UPDATE ON accounts 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Insert initial account data (for development/testing)
-- Get the user ID for brandon.johnston83@gmail.com and create an account
INSERT INTO accounts (account_name, user_id, total_balance) 
SELECT 'Business Checking - Card Shop', id, 0.00
FROM users 
WHERE email = 'brandon.johnston83@gmail.com';
