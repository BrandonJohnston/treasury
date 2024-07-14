import React, {useMemo} from "react";
import {Accordion, AccordionDetails} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import TrendingUpIcon from "@mui/icons-material/TrendingUp";
import TrendingDownIcon from "@mui/icons-material/TrendingDown";
import TransactionRowDetails from "./TransactionRowDetails.tsx";
import {ITransaction} from "../../Types/transactions.tsx";
import {TransactionDirections} from "../../Types/transactions.tsx";
import * as TransactionStyles from "./TransactionsTableStyles.tsx";

interface ITransactionRowProps {
	transaction: ITransaction;
}

export default function TransactionTableRow(props: ITransactionRowProps): React.ReactNode {
	// Aliases
	const {transaction} = props;

	// Local State
	//

	const directionIcon = useMemo(() => {
		return transaction.direction === TransactionDirections.deposit ? <TrendingUpIcon style={{color: "#11ba71"}} /> : <TrendingDownIcon style={{color: "#ba1133"}} />
	}, [transaction]);

	return (
		<Accordion disableGutters={true}>
			<TransactionStyles.TransAccordionSummary expandIcon={<ExpandMoreIcon />}>
				<div className="table-cell trans-dir">
					{directionIcon}
				</div>
				<div className="table-cell trans-name">
					<p>{transaction.name}</p>
				</div>
				<div className="table-cell trans-amount">
					<p>${transaction.amount}</p>
				</div>
			</TransactionStyles.TransAccordionSummary>
			<AccordionDetails>
				<TransactionRowDetails transaction={transaction} />
			</AccordionDetails>
		</Accordion>
	);
}
