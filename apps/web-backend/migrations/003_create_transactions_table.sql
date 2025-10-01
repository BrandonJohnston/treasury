-- Migration: 003_create_transactions_table.sql
-- Description: Creates the transactions table for the Treasury application
-- Created: 2025-09-30

-- Drop table if exists (for reset functionality)
DROP TABLE IF EXISTS transactions CASCADE;

-- Create transactions table
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount DECIMAL(10, 2) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL CHECK (transaction_type IN ('income', 'expense')),
    transaction_date TIMESTAMP DEFAULT NOW() NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- Create indexes for better performance
CREATE INDEX idx_transactions_account_id ON transactions(account_id);
CREATE INDEX idx_transactions_date ON transactions(transaction_date);
CREATE INDEX idx_transactions_type ON transactions(transaction_type);

-- Create trigger to automatically update updated_at timestamp for transactions
CREATE OR REPLACE FUNCTION update_transactions_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_transactions_updated_at 
    BEFORE UPDATE ON transactions 
    FOR EACH ROW 
    EXECUTE FUNCTION update_transactions_updated_at();

-- Function to update account balance when transactions change
CREATE OR REPLACE FUNCTION update_account_balance()
RETURNS TRIGGER AS $$
DECLARE
    account_uuid UUID;
    balance_change DECIMAL(10, 2);
BEGIN
    -- Determine which account_id to use and the balance change
    IF TG_OP = 'DELETE' THEN
        account_uuid := OLD.account_id;
        -- For deletions, reverse the transaction effect
        IF OLD.transaction_type = 'credit' THEN
            balance_change := -OLD.amount; -- Remove credit
        ELSE
            balance_change := OLD.amount; -- Remove debit (add back)
        END IF;
    ELSE
        account_uuid := NEW.account_id;
        -- For insertions and updates, calculate the net effect
        IF TG_OP = 'INSERT' THEN
            IF NEW.transaction_type = 'credit' THEN
                balance_change := NEW.amount;
            ELSE
                balance_change := -NEW.amount;
            END IF;
        ELSE -- UPDATE
            -- Calculate the difference between old and new
            DECLARE
                old_effect DECIMAL(10, 2);
                new_effect DECIMAL(10, 2);
            BEGIN
                -- Calculate old effect
                IF OLD.transaction_type = 'credit' THEN
                    old_effect := OLD.amount;
                ELSE
                    old_effect := -OLD.amount;
                END IF;
                
                -- Calculate new effect
                IF NEW.transaction_type = 'credit' THEN
                    new_effect := NEW.amount;
                ELSE
                    new_effect := -NEW.amount;
                END IF;
                
                balance_change := new_effect - old_effect;
            END;
        END IF;
    END IF;
    
    -- Update the account balance
    UPDATE accounts 
    SET total_balance = total_balance + balance_change,
        updated_at = CURRENT_TIMESTAMP
    WHERE id = account_uuid;
    
    -- Return appropriate record
    IF TG_OP = 'DELETE' THEN
        RETURN OLD;
    ELSE
        RETURN NEW;
    END IF;
END;
$$ language 'plpgsql';

-- Create triggers for all transaction operations
CREATE TRIGGER trigger_update_balance_on_insert
    AFTER INSERT ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_account_balance();

CREATE TRIGGER trigger_update_balance_on_update
    AFTER UPDATE ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_account_balance();

CREATE TRIGGER trigger_update_balance_on_delete
    AFTER DELETE ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_account_balance();
