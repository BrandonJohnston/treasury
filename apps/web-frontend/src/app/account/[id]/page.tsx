import { auth0 } from "@/lib/auth0";
import { redirect } from "next/navigation";

interface IGetAccountDetailsPayload {
	email?: string;
	provider: string;
	provider_id: string;
}

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

    if (session?.user) {
		try {
			const provider = session.user.sub.split("|");
			const payload: IGetAccountDetailsPayload = {
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

			const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/accounts/${accountId}?${queryParams.toString()}`, {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
				},
			});
            console.log(response);

			const data = await response.json();
            console.log(data);

            // Save account data to context
            //
		} catch (error) {
			console.error("Error: ", error);
		}
	}

	return (
		<div className="w-full p-10">
			<h1>Account Details</h1>
            <p>Account ID: {accountId}</p>
		</div>
	);
}
