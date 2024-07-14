import {lazy} from "react";
import {RouteObject} from "react-router-dom";
import App from "../App";

const TransactionsView = lazy(() => import("../Views/Transactions/TransactionsView.tsx"));

const routes: RouteObject[] = [
	{
		path: "/",
		element: <App />,
		children: [
			{
				index: true,
				element: <TransactionsView />,
			},
			{
				path: "transactions",
				element: <TransactionsView />,
			},
		],
	},
];

export default routes;
