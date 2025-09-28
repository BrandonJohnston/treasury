import Link from "next/link";
import {
    Sidebar,
    SidebarContent,
    SidebarGroup,
    SidebarGroupContent,
    SidebarMenu,
    SidebarMenuItem,
    SidebarMenuButton,
    SidebarHeader,
    SidebarGroupLabel,
} from "@/components/ui/sidebar";
import { LayoutDashboard } from "lucide-react";

import AppSidebarAccountHeader from "./AppSidebarAccountHeader";
import AppSidebarFooter from "./AppSidebarFooter";

// Menu items.
const items = [
    {
      title: "Dashboard",
      url: "/dashboard",
      icon: LayoutDashboard,
    },
];
  
export function AppSidebar() {

    return (
        <Sidebar className="pt-[60px]" collapsible="icon">
            <SidebarHeader />
            <SidebarContent>
                <SidebarGroup>
                    <SidebarGroupLabel>General</SidebarGroupLabel>
                    <SidebarGroupContent>
                        <SidebarMenu>
                            {items.map((item) => (
                                <SidebarMenuItem key={item.title}>
                                    <SidebarMenuButton asChild>
                                        <Link href={item.url} className="w-full">
                                            <item.icon />
                                            <span>{item.title}</span>
                                        </Link>
                                    </SidebarMenuButton>
                                </SidebarMenuItem>
                            ))}
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
                <AppSidebarAccountHeader />
            </SidebarContent>
            
            <AppSidebarFooter />
        </Sidebar>
    )
};
