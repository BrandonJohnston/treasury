import {ITransaction} from "../../Types/transactions.tsx";
import transactionsData from "../../assets/transactions.json";

export default async function getTransactions(): Promise<ITransaction[]> {
	try {
		// Typically make a fetch call here
		const jsonData = transactionsData;

		return jsonData.data as ITransaction[];
	} catch (error) {
		console.error("getTransactions() error: ", error);
		return [];
	}
}
