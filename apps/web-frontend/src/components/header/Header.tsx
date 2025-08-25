import Link from "next/link";
import { Button } from "@/components/ui/button";

export default function Header() {
	return (
		<div className="absolute bg-sky-300 border-b-1 border-sky-800 flex items-center justify-between left-0 shadow-md top-0 w-full h-[60px]">
			<h1 className="mx-[16px]">
				<Link href="/">Treasury</Link>
			</h1>
			<div>
				<Button asChild className={"mx-[16px]"}>
					<Link href="/login">Login</Link>
				</Button>
			</div>
		</div>
	);
}
