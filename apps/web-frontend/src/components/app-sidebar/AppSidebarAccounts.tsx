"use client";

import { useEffect } from "react";
import {
    SidebarGroup,
    SidebarGroupLabel,
    SidebarGroupAction,
    SidebarGroupContent,
    SidebarMenu,
    SidebarMenuItem,
    SidebarMenuButton,
} from "@/components/ui/sidebar";
import { Plus, Landmark } from "lucide-react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { useAccounts } from "@/contexts/AccountContext";
import { IAccount } from "@/types/Accounts";

interface AppSidebarAccountsProps {
    accountsData: IAccount[];
}

export default function AppSidebarAccounts(props: AppSidebarAccountsProps) {
    const { accountsData } = props;
    const router = useRouter();
    const { accounts, setAccountsData } = useAccounts();

    const handleAddAccount = () => {
        router.push("/account/create");
    }

    useEffect(() => {
        setAccountsData(accountsData);
    }, [accountsData]);

    return (
        <SidebarGroup>
            <SidebarGroupLabel>Accounts</SidebarGroupLabel>
            <SidebarGroupAction title="Add Account" onClick={handleAddAccount}>
                <Plus /> <span className="sr-only">Add Account</span>
            </SidebarGroupAction>
            <SidebarGroupContent>
                <SidebarMenu>
                    {accounts.map((account) => (
                        <SidebarMenuItem key={account.id}>
                            <SidebarMenuButton asChild>
                                <Link href={`/account/${account.id}`} className="w-full">
                                    <Landmark />
                                    <span>{account.accountName}</span>
                                </Link>
                            </SidebarMenuButton>
                        </SidebarMenuItem>
                    ))}
                </SidebarMenu>
            </SidebarGroupContent>
        </SidebarGroup>
    );
};
