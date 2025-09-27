import { auth0 } from "@/lib/auth0";
import Link from "next/link";
import { SidebarTrigger } from "@/components/ui/sidebar";
import { Button } from "@/components/ui/button";

export default async function Header() {
	const session = await auth0.getSession();

	return (
		<div className="fixed bg-sky-300 border-b-1 border-sky-800 flex items-center justify-between left-0 shadow-md top-0 w-full h-[60px] px-[60px] z-20">
			{session && (
				<div className="absolute left-0 top-0 h-full w-[61px] flex items-center justify-center">
					<SidebarTrigger />
				</div>
			)}
			<h1>
				<Link href="/">Treasury</Link>
			</h1>
			<div>
				<Button asChild className={"mx-[16px]"}>
					{!session?.user && (
						<Link href="/auth/login">Login</Link>
					)}
				</Button>
			</div>
		</div>
	);
}
