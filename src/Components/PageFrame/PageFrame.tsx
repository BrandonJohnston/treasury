import React from "react";
import PageHeader from "../PageHeader/PageHeader.tsx";
import PageFrameStyles from "./PageFrameStyles.tsx";

interface IPageFrameProps {
	children: React.ReactNode;
}

export default function PageFrame(props: IPageFrameProps): React.ReactNode {
	// Aliases
	const {children} = props;

	return (
		<PageFrameStyles>
			<PageHeader />
			{children}
		</PageFrameStyles>
	);
}
