import { create } from "zustand";
import type { ConfigActions, ConfigState } from "@/types/stores";
import { DEFAULT_STATE } from "@/config/constants";

export const useConfigStore = create<ConfigState & ConfigActions>(set => ({
	// --- STATE
	keyboardShortcuts: {},
	selectedEnv: null,
	userSettings: DEFAULT_STATE.CONFIG,

	// --- ACTIONS
	addKeyboardShortcut: newShortcut =>
		set(state => ({
			keyboardShortcuts: {
				...state.keyboardShortcuts,
				...newShortcut,
			},
		})),

	removeKeyboardShortcut: shortcutToRemove =>
		set(state => {
			const updatedShortcuts = { ...state.keyboardShortcuts };
			for (const key in shortcutToRemove) {
				if (key in updatedShortcuts) {
					delete updatedShortcuts[key];
				}
			}
			return { keyboardShortcuts: updatedShortcuts };
		}),

	addUserSettings: newSetting =>
		set(state => ({
			userSettings: {
				...state.userSettings,
				...newSetting,
			},
		})),

	removeUserSettings: settingToRemove =>
		set(state => {
			const updatedSettings = { ...state.userSettings };

			for (const key in settingToRemove) {
				if (key in updatedSettings) {
					delete updatedSettings[key as keyof typeof updatedSettings];
				}
			}

			return { userSettings: updatedSettings };
		}),

	setSelectedEnv: env => set({ selectedEnv: env }),

	setInitialState: state => set(() => state),
}));
