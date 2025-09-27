import { auth0 } from "@/lib/auth0";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Menu } from "lucide-react";

export default async function Header() {
	const session = await auth0.getSession();

	console.log("session: ", session);

	return (
		<div className="absolute bg-sky-300 border-b-1 border-sky-800 flex items-center justify-between left-0 shadow-md top-0 w-full h-[60px] px-[60px]">
			<div className="absolute left-0 top-0 h-full w-[61px] flex items-center justify-center">
				<Button variant="ghostNoHover" className="cursor-pointer">
					<Menu className="size-6" />
				</Button>
			</div>
			<h1 className="mx-[16px]">
				<Link href="/">Treasury</Link>
			</h1>
			<div>
				<Button asChild className={"mx-[16px]"}>
					{session?.user ? (
						<Link href="/auth/logout">Logout</Link>
					) : (
						<Link href="/auth/login">Login</Link>
					)}
				</Button>
			</div>
		</div>
	);
}
