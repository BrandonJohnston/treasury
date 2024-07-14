import React from "react";
import PageFrame from "./Components/PageFrame/PageFrame.tsx";
import PageMain from "./Components/PageMain/PageMain.tsx";

export default function App(): React.ReactElement {

	return (
		<PageFrame>
			<PageMain />
		</PageFrame>
	)
}
