import React, {useState, useEffect} from "react";
import TransactionsViewStyles from "./TransactionsViewStyles.tsx";
import TransactionsTable from "../../Components/TransactionsTable/TransactionsTable.tsx";
import getTransactions from "./TransactionsUtils.tsx";
import {ITransaction} from "../../Types/transactions.tsx";

export default function TransactionsView(): React.ReactNode {
	// Aliases
	//

	// Local State
	const [transactionsList, setTransactionsList] = useState<ITransaction[]>([]);

	useEffect(() => {
		async function getData(): Promise<void> {
			const response = await getTransactions();
			setTransactionsList(response);
		}

		if (transactionsList.length === 0) {
			getData();
		}
	}, [transactionsList]);

	return (
		<TransactionsViewStyles>
			<h2>Transactions</h2>
			<TransactionsTable transactions={transactionsList} />
		</TransactionsViewStyles>
	);
}
