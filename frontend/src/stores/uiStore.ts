import { create } from "zustand";
import type { UIActions, UIState } from "@/types/stores";

export const useUIStore = create<UIState & UIActions>(set => ({
	// --- STATE
	modalId: null,
	openTabId: null,
	searchQuery: "",

	// --- ACTIONS
	setCurrentModal: newId => set(() => ({ modalId: newId })),
	setOpenTab: newTabId => set(() => ({ openTabId: newTabId })),
	setSearchQuery: query => set(() => ({ searchQuery: query })),
	setInitialState: state => set(() => state),
}));
