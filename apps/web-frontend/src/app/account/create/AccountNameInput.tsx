"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { useAccounts } from "@/contexts/AccountContext";

interface AccountNameInputProps {
	email?: string;
	provider: string;
	providerId: string;
}

export default function AccountNameInput(props: AccountNameInputProps) {
	// Aliases
	const { email, provider, providerId } = props;
	const router = useRouter();
	const { addAccount } = useAccounts();

    const [editing, setEditing] = useState(true);
    const [accountName, setAccountName] = useState("");

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setAccountName(event.target.value);
    };

    const handleEnterKeyPress = (event: React.KeyboardEvent<HTMLInputElement>) => {
        if (event.key === "Enter") {
            console.log("Enter key pressed! Value:", event.currentTarget.value);
            setAccountName(event.currentTarget.value);
            handleBlur();
        }
    };

    const handleBlur = () => {
        if (accountName.length > 0) {
            setEditing(false);
            handleSave();
        }
    };

    const handleDoubleClick = () => {
        setEditing(true);
    };

    const handleSave = async () => {
        console.log("saving...");
        const payload = {
            accountName,
            email,
            provider,
            providerId,
        }

        const response = await fetch("http://localhost:8080/api/accounts/create", {
            method: "POST",
            body: JSON.stringify(payload),
        });
        const data = await response.json();
        console.log(data);

        if (data.status === "ok") {
            console.log("calling addAccount()");
            // Add the new account to the context
            addAccount(data.account);
            
            // redirect to account page
            router.push(`/account/${data.account.id}`);
        }
    };
    
	return (
        <div className="grid gap-3">
            <Label htmlFor="account-name">Account Name</Label>
            {editing ? (
                <Input id="account-name" 
                    type="text" 
                    placeholder="Account Name" 
                    autoFocus 
                    onKeyDown={handleEnterKeyPress}
                    onBlur={handleBlur}
                    value={accountName}
                    onChange={handleChange} />
            ) : (
                <div onDoubleClick={handleDoubleClick}>{accountName}</div>
            )}
        </div>
	);
}
