import { auth0 } from "@/lib/auth0";
import { redirect } from "next/navigation";

import AccountNameInput from "./AccountNameInput";

export default async function AccountCreate() {
	// Aliases
	//

	const session = await auth0.getSession();

	if (!session?.user) {
		redirect("/");
	}

	const { email, sub } = session.user;
	const provider = sub?.split("|")[0];
	const providerId = sub?.split("|")[1];

	return (
		<div className="w-full p-10">
            <AccountNameInput email={email} provider={provider} providerId={providerId} />
		</div>
	);
}
