import type {
	Collections,
	Environments,
	EnvironmentVariables,
	Folders,
	KeyboardShortcuts,
	ModalField,
	CurrentTab,
	OpenTabs,
	RequestsList,
	SearchQueryField,
	SelectedEnvId,
	UserSettings,
	OpenTab,
} from "./index";

// ============== UI STORE
export interface UIState {
	openTabId: CurrentTab;
	modalId: ModalField;
	searchQuery: SearchQueryField;
}

export interface UIActions {
	setCurrentModal: (newId: UIState["modalId"] | null) => void;
	setOpenTab: (newId: UIState["openTabId"] | null) => void;
	setSearchQuery: (query: UIState["searchQuery"]) => void;
	setInitialState: (state: UIState) => void;
}

// ============== CONFIG STORE
export interface ConfigState {
	keyboardShortcuts: KeyboardShortcuts;
	userSettings: UserSettings;
	selectedEnv: SelectedEnvId;
}

export interface ConfigActions {
	addKeyboardShortcut: (newShortcut: ConfigState["keyboardShortcuts"]) => void;
	removeKeyboardShortcut: (newShortcut: ConfigState["keyboardShortcuts"]) => void;
	addUserSettings: (newSetting: NonNullable<ConfigState["userSettings"]>) => void;
	removeUserSettings: (newSetting: NonNullable<ConfigState["userSettings"]>) => void;
	setSelectedEnv: (env: ConfigState["selectedEnv"]) => void;
	setInitialState: (state: ConfigState) => void;
}

// ============== DATA STORE
export interface DataState {
	collections: Collections;
	folders: Folders;
	requests: RequestsList;
	environments: Environments;
	environmentVariables: EnvironmentVariables;
	openTabs: OpenTabs;
}

export interface DataActions {
	addCollection: (newCollection: DataState["collections"][number]) => void;
	deleteCollection: (collectionId: DataState["collections"][number]["id"]) => void;
	addFolder: (newFolder: DataState["folders"][number]) => void;
	deleteFolder: (folderId: DataState["folders"][number]["id"]) => void;
	addRequest: (newRequest: DataState["requests"][number]) => void;
	deleteRequest: (requestId: DataState["requests"][number]["id"]) => void;
	addEnvironment: (newEnvironment: DataState["environments"][number]) => void;
	deleteEnvironment: (environmentId: DataState["environments"][number]["id"]) => void;
	addEnvironmentVariable: (key: string, value: string) => void;
	deleteEnvironmentVariable: (key: string) => void;
	updateEnvironmentVariable: (key: string, value: string) => void;
	addOpenTab: (tab: OpenTab) => void;
	updateTab: (tab: OpenTab) => void;
	closeTab: (tab: OpenTab) => void;
	setInitialState: (state: DataState) => void;
}
