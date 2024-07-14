import styled from "@emotion/styled";
import {AccordionSummary} from "@mui/material";

export const TransactionsTableStyles = styled.div`
	border-bottom: 1px solid rgba(0, 0, 0, .07);

	.table-row {
		align-items: center;
		border-top: 1px solid rgba(0, 0, 0, .07);
		cursor: pointer;
		display: flex;
		justify-content: space-between;
	}

	.table-cell {
		padding: 0.5rem;

		p {
			margin: 0;
		}
	}

	.trans-dir {
		align-items: center;
		display: flex;
		width: 2.5rem;
	}

	.trans-name {
		flex-grow: 1;
	}

	.trans-details {
		border-top: 1px solid rgba(0, 0, 0, .07);
		margin: 0 .5rem
	}

	.trans-details-row {
		align-items: flex-start;
		display: flex;
		justify-content: flex-start;
		padding-top: 1rem;

		> * {
			font-size: 1rem;
			margin: 0 1rem 0 0;
		}
	}
`;

export const TransAccordionSummary = styled(AccordionSummary)({
	display: "flex",
});
