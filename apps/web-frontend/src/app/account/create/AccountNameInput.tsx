"use client";

import React, { useState } from "react";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";

export default function AccountNameInput() {
	// Aliases
	//

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
        const response = await fetch("/api/account/create", {
            method: "POST",
            body: JSON.stringify({ accountName }),
        });
        const data = await response.json();
        console.log(data);
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
