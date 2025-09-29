"use client";

import {
    SidebarGroup,
    SidebarGroupLabel,
    SidebarGroupAction,
} from "@/components/ui/sidebar";
import { Plus } from "lucide-react";
import { redirect } from "next/navigation";

export default function AppSidebarAccountHeader() {

    const handleAddAccount = () => {
        console.log("Add Account");
        redirect("/account/create");
    }

    return (
        <SidebarGroup>
            <SidebarGroupLabel>Accounts</SidebarGroupLabel>
            <SidebarGroupAction title="Add Project" onClick={handleAddAccount}>
                <Plus /> <span className="sr-only">Add Account</span>
            </SidebarGroupAction>
        </SidebarGroup>
    );
};
