export enum TransactionDirections {
	deposit = "deposit",
	withdrawl = "withdrawl",
}

export interface ITransaction {
	effectiveDate: string;
	name: string;
	amount: number;
	type: string;
	direction: TransactionDirections;
	memo: string;
}
