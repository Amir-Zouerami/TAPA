import type { models } from "@wails/go/models";

// type Folders = models.Folder[];

export type CurrentTab = number | null;
export type SearchQueryField = string;
export type ModalField = string | null;

export type KeyboardShortcuts = Record<number, models.KeyboardShortcut>;
export type UserSettings = models.UserSettings;
export type SelectedEnvId = number | null;

export type Collections = models.CollectionsTree["collections"];
export type Folders = models.CollectionsTree["folders"];
export type RequestsList = models.CollectionsTree["request_list"];
export type CollectionVariables = Record<string, string>;
export type Environments = Record<number, models.Environment>;
export type EnvironmentVariables = models.DashboardData["current_env_variables"];
export type OpenTab = models.DashboardData["app_state"]["open_tabs"][0];
export type OpenTabs = OpenTab[];
