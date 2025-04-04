import { useQuery } from "@tanstack/react-query";
import { LoadDashboardData, GetCollectionsTree } from "@wails/go/services/DashboardService";

export const useLoadDashboardData = () => {
	return useQuery({
		queryKey: ["ALL", "STATE_INITIALIZER"],
		queryFn: async () => {
			const [dashboardData, collectionTree] = await Promise.all([
				LoadDashboardData(),
				GetCollectionsTree(),
			]);

			return { dashboardData, collectionTree };
		},
	});
};
