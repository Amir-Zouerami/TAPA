import { Outlet, createRootRoute } from "@tanstack/react-router";
import { ThemeProvider } from "@/components/providers/theme-provider";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

export const Route = createRootRoute({
	component: () => (
		<ThemeProvider defaultTheme="dark" storageKey="tapa-ui-theme">
			<Outlet />
			<TanStackRouterDevtools />
		</ThemeProvider>
	),
});
