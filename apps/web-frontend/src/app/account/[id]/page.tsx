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

	// Placeholder account data
	const accountData = {
		id: accountId,
		accountName: "My Savings Account",
		balance: 1250.75,
		currency: "USD",
		createdAt: "2024-01-15T10:30:00Z",
		lastUpdated: "2024-01-20T14:45:00Z"
	};

	return (
		<div className="w-full p-10">
			<div className="max-w-4xl mx-auto">
				<h1 className="text-3xl font-bold mb-6">Account Details</h1>
				
				<div className="bg-white p-6 rounded-lg shadow-lg mb-6">
					<h2 className="text-xl font-semibold mb-4">{accountData.accountName}</h2>
					<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
						<div>
							<p className="text-sm text-gray-600">Account ID</p>
							<p className="font-mono text-lg">{accountData.id}</p>
						</div>
						<div>
							<p className="text-sm text-gray-600">Balance</p>
							<p className="text-2xl font-bold text-green-600">
								${accountData.balance.toLocaleString()}
							</p>
						</div>
						<div>
							<p className="text-sm text-gray-600">Currency</p>
							<p className="text-lg">{accountData.currency}</p>
						</div>
						<div>
							<p className="text-sm text-gray-600">Created</p>
							<p className="text-lg">
								{new Date(accountData.createdAt).toLocaleDateString()}
							</p>
						</div>
					</div>
				</div>

				<div className="bg-white p-6 rounded-lg shadow-lg">
					<h3 className="text-lg font-semibold mb-4">Recent Transactions</h3>
					<div className="space-y-3">
						<div className="flex justify-between items-center p-3 bg-gray-50 rounded">
							<div>
								<p className="font-medium">Deposit</p>
								<p className="text-sm text-gray-600">2024-01-20</p>
							</div>
							<p className="text-green-600 font-semibold">+$500.00</p>
						</div>
						<div className="flex justify-between items-center p-3 bg-gray-50 rounded">
							<div>
								<p className="font-medium">Transfer Out</p>
								<p className="text-sm text-gray-600">2024-01-18</p>
							</div>
							<p className="text-red-600 font-semibold">-$250.00</p>
						</div>
						<div className="flex justify-between items-center p-3 bg-gray-50 rounded">
							<div>
								<p className="font-medium">Interest</p>
								<p className="text-sm text-gray-600">2024-01-15</p>
							</div>
							<p className="text-green-600 font-semibold">+$0.75</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}
