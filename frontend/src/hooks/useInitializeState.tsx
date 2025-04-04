import { useEffect, useRef } from "react";
import { useUIStore } from "@/stores/uiStore";
import type { models } from "@wails/go/models";
import { useDataStore } from "@/stores/dataStore";
import { useConfigStore } from "@/stores/configStore";
import type { ConfigState, DataState, UIState } from "@/types/stores";

export interface DashboardData {
	dashboardData: models.DashboardData;
	collectionTree: models.CollectionsTree;
}

export const useInitializeState = (data: DashboardData | undefined) => {
	const setInitialConfigState = useConfigStore(state => state.setInitialState);
	const setInitialDataState = useDataStore(state => state.setInitialState);
	const setInitialUIState = useUIStore(state => state.setInitialState);
	const initializedRef = useRef<boolean>(false);

	useEffect(() => {
		if (data && !initializedRef.current) {
			const initialUIState: UIState = {
				modalId: null,
				openTabId: data.dashboardData.app_state.selected_tab ?? null,
				searchQuery: "",
			};
			setInitialUIState(initialUIState);

			const initialConfigState: ConfigState = {
				keyboardShortcuts: data.dashboardData.keyboard_shortcuts,
				userSettings: data.dashboardData.user_settings,
				selectedEnv: data.dashboardData.app_state.selected_environment ?? null,
			};
			setInitialConfigState(initialConfigState);

			const initialDataState: DataState = {
				collections: data.collectionTree.collections,
				folders: data.collectionTree.folders,
				requests: data.collectionTree.request_list,
				environments: data.dashboardData.environments,
				environmentVariables: data.dashboardData.current_env_variables,
				openTabs: data.dashboardData.app_state.open_tabs,
			};
			setInitialDataState(initialDataState);

			initializedRef.current = true;
		}
	}, [data]);
};
