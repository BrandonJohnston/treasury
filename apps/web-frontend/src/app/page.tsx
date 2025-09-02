import { auth } from "./lib/auth";

export default async function Home() {
	const session = await auth();

	console.log("session: ", session);

	if (session) {
		return <p>redirect to logged in page.</p>
	}

	return (
		<div className="flex font-sans">
			<p>main landing page</p>
		</div>
	);
}
