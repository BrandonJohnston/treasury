import React, {useMemo} from "react";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import {ITransaction} from "../../Types/transactions.tsx";

dayjs.extend(relativeTime);

interface ITransactionRowDetailsProps {
	transaction: ITransaction;
}

export default function TransactionRowDetails(props: ITransactionRowDetailsProps): React.ReactNode {
	// Aliases
	const {transaction} = props;

	// Local State
	const formattedDate = useMemo(() => {
		return dayjs(transaction.effectiveDate).format("MMMM D, YYYY");
	}, [transaction]);

	const relativeDate = useMemo(() => {
		return dayjs(transaction.effectiveDate).fromNow();
	}, [transaction]);

	return (
		<div className="trans-details">
			<div className="trans-details-row">
				<h5>Type:</h5>
				<p>{transaction.type}</p>
			</div>
			<div className="trans-details-row">
				<h5>Date:</h5>
				<p>{formattedDate} ({relativeDate})</p>
			</div>
			<div className="trans-details-row">
				<h5>Memo:</h5>
				<p>{transaction.memo}</p>
			</div>
		</div>
	);
}
