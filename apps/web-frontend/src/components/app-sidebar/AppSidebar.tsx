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

import { auth0 } from "@/lib/auth0";
import AppSidebarAccounts from "./AppSidebarAccounts";
import AppSidebarFooter from "./AppSidebarFooter";

import { IAccount } from "@/types/Accounts";

interface IGetAccountsPayload {
	email?: string;
	provider: string;
	provider_id: string;
}

// Menu items.
const items = [
    {
      title: "Dashboard",
      url: "/dashboard",
      icon: LayoutDashboard,
    },
];
  
export async function AppSidebar() {
    const session = await auth0.getSession();

    const accounts: IAccount[] = [];

	if (session?.user) {
		try {
			const provider = session.user.sub.split("|");
			const payload: IGetAccountsPayload = {
				email: session.user.email,
				provider: provider[0],
				provider_id: provider[1],
			};

			// Build query parameters
			const queryParams = new URLSearchParams({
				email: payload.email || '',
				provider: payload.provider,
				provider_id: payload.provider_id,
			});

			const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/accounts?${queryParams.toString()}`, {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
				},
			});

			const data = await response.json();
			console.log(data);

            // Save accounts to context
            accounts.push(...data.accounts);
		} catch (error) {
			console.error("Error: ", error);
		}
	}

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
                <AppSidebarAccounts accountsData={accounts} />
            </SidebarContent>
            
            <AppSidebarFooter />
        </Sidebar>
    )
};
