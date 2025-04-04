import { useEffect } from "react";
import { APP_META_DATA } from "@/config/constants";

/**
 * Global handler to prevent the default context menu unless an element
 * or its ancestor has the `data-allow-contextmenu="true"` attribute.
 * @param event - The MouseEvent from the contextmenu event.
 */
const handleGlobalContextMenu = (event: MouseEvent): void => {
	const targetElement = event.target as Element;
	const allowContextMenu = targetElement.closest(`[${APP_META_DATA.ALLOW_CONTEXT_MENU_ATTR}="true"]`);

	if (!allowContextMenu) {
		event.preventDefault();
	}
};

/**
 * A custom React hook to disable the default browser context menu globally,
 * except for elements specifically marked with `data-allow-contextmenu="true"`.
 */
export function useDisableContextMenu(): void {
	useEffect(() => {
		document.addEventListener("contextmenu", handleGlobalContextMenu);

		return () => {
			document.removeEventListener("contextmenu", handleGlobalContextMenu);
		};
	}, []);
}
