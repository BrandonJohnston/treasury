"use client";

import React, { createContext, useContext, useState, useEffect } from 'react';

import { IAccount } from "@/types/Accounts";

interface AccountContextType {
    accounts: IAccount[];
    setAccountsData: (accounts: IAccount[]) => Promise<void>;
    addAccount: (account: IAccount) => void;
}

const AccountContext = createContext<AccountContextType | undefined>(undefined);

export function AccountProvider({ children }: { children: React.ReactNode }) {
    const [accounts, setAccounts] = useState<IAccount[]>([]);

    const setAccountsData = async (accounts: IAccount[]) => {
        setAccounts(accounts);
    };

    const addAccount = (account: IAccount) => {
        setAccounts(prev => [...prev, account]);
    };

    return (
        <AccountContext.Provider value={{ accounts, setAccountsData, addAccount }}>
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
