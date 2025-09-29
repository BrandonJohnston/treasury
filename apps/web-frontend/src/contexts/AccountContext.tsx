"use client";

import React, { createContext, useContext, useState, useEffect } from 'react';

interface Account {
    id: string;
    accountName: string;
    // Add other account properties as needed
}

interface AccountContextType {
    accounts: Account[];
    refreshAccounts: () => Promise<void>;
    addAccount: (account: Account) => void;
}

const AccountContext = createContext<AccountContextType | undefined>(undefined);

export function AccountProvider({ children }: { children: React.ReactNode }) {
    const [accounts, setAccounts] = useState<Account[]>([]);

    const refreshAccounts = async () => {
        // TODO: Replace with actual API call when backend is ready
        console.log('Refreshing accounts...');
        // For now, just log - you can add actual API call here later
    };

    const addAccount = (account: Account) => {
        setAccounts(prev => [...prev, account]);
    };

    useEffect(() => {
        refreshAccounts();
    }, []);

    return (
        <AccountContext.Provider value={{ accounts, refreshAccounts, addAccount }}>
            {children}
        </AccountContext.Provider>
    );
}

export function useAccounts() {
    const context = useContext(AccountContext);
    if (context === undefined) {
        throw new Error('useAccounts must be used within an AccountProvider');
    }
    return context;
}
