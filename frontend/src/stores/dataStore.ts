import { create } from "zustand";
import type { DataActions, DataState } from "@/types/stores";

export const useDataStore = create<DataState & DataActions>(set => ({
	// --- STATE
	collections: {},
	folders: {},
	requests: {},
	environments: [],
	currentEnvironment: null,
	environmentVariables: {},
	openTabs: [],

	// --- ACTIONS
	addCollection: collection =>
		set(state => ({
			collections: { ...state.collections, [collection.id]: collection },
		})),

	deleteCollection: id =>
		set(state => {
			const newCollections = { ...state.collections };
			delete newCollections[id];
			return { collections: newCollections };
		}),

	addFolder: folder =>
		set(state => ({
			folders: { ...state.folders, [folder.id]: folder },
		})),

	deleteFolder: id =>
		set(state => {
			const newFolders = { ...state.folders };
			delete newFolders[id];
			return { folders: newFolders };
		}),

	addRequest: request =>
		set(state => ({
			requests: { ...state.requests, [request.id]: request },
		})),

	deleteRequest: id =>
		set(state => {
			const newRequests = { ...state.requests };
			delete newRequests[id];
			return { requests: newRequests };
		}),

	addEnvironment: environment =>
		set(state => ({
			environments: { ...state.environments, [environment.id]: environment },
		})),

	deleteEnvironment: id =>
		set(state => {
			const newEnvs = { ...state.environments };
			delete newEnvs[id];
			return { environments: newEnvs };
		}),

	addEnvironmentVariable: (key, value) =>
		set(state => ({
			environmentVariables: {
				...state.environmentVariables,
				[key]: value,
			},
		})),

	deleteEnvironmentVariable: key =>
		set(state => {
			const newVars = { ...state.environmentVariables };
			delete newVars[key];
			return { environmentVariables: newVars };
		}),

	updateEnvironmentVariable: (key: string, value: string) =>
		set(state => ({
			environmentVariables: {
				...state.environmentVariables,
				[key]: value,
			},
		})),

	addOpenTab: tab =>
		set(state => {
			const exists = state.openTabs.some(t => t.request_id === tab.request_id);
			return exists ? state : { openTabs: [...state.openTabs, tab] };
		}),

	updateTab: tab =>
		set(state => ({
			openTabs: state.openTabs.map(t => (t.request_id === tab.request_id ? { ...t, ...tab } : t)),
		})),

	closeTab: tab =>
		set(state => ({
			openTabs: state.openTabs.filter(t => t.request_id !== tab.request_id),
		})),

	setInitialState: state => set(() => state),
}));
