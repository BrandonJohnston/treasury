import { auth0 } from "@/lib/auth0";
import { redirect } from "next/navigation";

export default async function Dashboard() {
	// Aliases
	//

	const session = await auth0.getSession();

	if (!session) {
		redirect("/");
	}

	return (
		<div className="login-form items-center justify-items-center w-full md:w-1/2 md:mx-auto p-10">
			dashboard protected route
		</div>
	);
}
