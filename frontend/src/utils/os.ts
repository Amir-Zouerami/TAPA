import { BrowserOpenURL } from "@wails/runtime";

export const openInUserBrowser = (url: string) => {
	BrowserOpenURL(url);
};
