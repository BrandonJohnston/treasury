import { auth0 } from "@/lib/auth0";
import { redirect } from "next/navigation";

interface AccountPageProps {
	params: Promise<{
		id: string;
	}>;
}

export default async function AccountPage({ params }: AccountPageProps) {
	const session = await auth0.getSession();

	if (!session?.user) {
		redirect("/");
	}

	const { id: accountId } = await params;

	return (
		<div className="w-full p-10">
			<h1>Account Details</h1>
            <p>Account ID: {accountId}</p>
		</div>
	);
}
