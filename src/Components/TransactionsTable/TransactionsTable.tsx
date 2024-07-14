import React from "react";
import {ITransaction} from "../../Types/transactions.tsx";
import TransactionTableRow from "./TransactionRow.tsx";
import * as TransactionStyles from "./TransactionsTableStyles.tsx";

interface ITransactionsTableProps {
	transactions: ITransaction[];
}

export default function TransactionsTable(props: ITransactionsTableProps): React.ReactNode {
	// Aliases
	const {transactions} = props;

	// Local State
	//

	return (
		<TransactionStyles.TransactionsTableStyles>
			{transactions?.map((transaction) => (
				<TransactionTableRow
					key={`${transaction.effectiveDate}-${transaction.name}`}
					transaction={transaction}
				/>
			))}
		</TransactionStyles.TransactionsTableStyles>
	);
}
