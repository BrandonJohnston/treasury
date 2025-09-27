import { auth0 } from "@/lib/auth0";
import { redirect } from "next/navigation";

export default async function Settings() {
	// Aliases
	//

	const session = await auth0.getSession();

	if (!session) {
		redirect("/");
	}

	return (
		<div className="items-center justify-items-center w-full p-10">
			Settings
		</div>
	);
}
