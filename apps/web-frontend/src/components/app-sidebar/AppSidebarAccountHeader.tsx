"use client";

import {
    SidebarGroup,
    SidebarGroupLabel,
    SidebarGroupAction,
} from "@/components/ui/sidebar";
import { Plus } from "lucide-react";
  
export default function AppSidebarAccountHeader() {

    const handleAddAccount = () => {
        console.log("Add Account");
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
