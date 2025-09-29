import { auth0 } from "@/lib/auth0";
import { redirect } from "next/navigation";

import AccountNameInput from "./AccountNameInput";

export default async function AccountCreate() {
	// Aliases
	//

	const session = await auth0.getSession();

	if (!session) {
		redirect("/");
	}

	return (
		<div className="w-full p-10">
            <AccountNameInput />
		</div>
	);
}
