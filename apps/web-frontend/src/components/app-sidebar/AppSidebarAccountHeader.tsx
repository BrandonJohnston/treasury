"use client";

import {
    SidebarGroup,
    SidebarGroupLabel,
    SidebarGroupAction,
    SidebarGroupContent,
    SidebarMenu,
    SidebarMenuItem,
    SidebarMenuButton,
} from "@/components/ui/sidebar";
import { Plus } from "lucide-react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { useAccounts } from "@/contexts/AccountContext";

export default function AppSidebarAccountHeader() {
    const router = useRouter();
    const { accounts } = useAccounts();

    console.log("accounts", accounts);

    const handleAddAccount = () => {
        router.push("/account/create");
    }

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
