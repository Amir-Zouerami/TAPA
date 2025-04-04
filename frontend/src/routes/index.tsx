import "../App.css";
import { UI_MESSAGES } from "@/config/constants";
import ErrorAlert from "@/components/Alerts/ErrorAlert";
import { useLoadDashboardData } from "@/hooks/api/dashboard";
import { useInitializeState } from "@/hooks/useInitializeState";
import CommandPalette from "@/components/layout/CommandPalette";
import { createFileRoute, redirect } from "@tanstack/react-router";
import { IsFirstLaunch } from "@wails/go/services/DashboardService";
// import { useDisableContextMenu } from "@/hooks/useDisableContextMenu";
import InitialLoadingPage from "@/components/reusable/InitialLoadingPage";
import ApplicationMenuBar from "@/components/layout/menuBar/ApplicationMenuBar";

export const Route = createFileRoute("/")({
	loader: async () => {
		const result = await IsFirstLaunch();
		if (result) {
			throw redirect({ to: "/welcome", replace: true });
		}
	},
	staleTime: Infinity,
	component: App,
});

function App() {
	// TODO: enable later
	// useDisableContextMenu();
	const { isPending, error, data } = useLoadDashboardData();
	useInitializeState(data);

	if (isPending) return <InitialLoadingPage />;
	if (error) return <ErrorAlert open reportable message={UI_MESSAGES.ERR_LOADING_INITIAL_STATE} />;

	return (
		<>
			<CommandPalette />
			<div className="App h-dvh w-dvw overflow-hidden">
				<ApplicationMenuBar />
				<br />
				<code>{JSON.stringify(data)}</code>
			</div>
		</>
	);
}
