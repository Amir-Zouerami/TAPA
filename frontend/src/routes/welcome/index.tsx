import { useState } from "react";
import { openInUserBrowser } from "@/utils/os";
import { APP_META_DATA } from "@/config/constants";
import { useMutation } from "@tanstack/react-query";
import TapaTextLogo from "@/components/branding/TapaTextLogo";
import { LoaderCircle, SquareArrowOutUpRight } from "lucide-react";
import { createFileRoute, Link, useLoaderData, useRouter } from "@tanstack/react-router";
import { GetUserConfigDir, SetFirstLaunchDone } from "@wails/go/services/DashboardService";

export const Route = createFileRoute("/welcome/")({
	component: RouteComponent,
	loader: async () => {
		return await GetUserConfigDir();
	},
});

function RouteComponent() {
	const router = useRouter();
	const [isLoading, setIsLoading] = useState(false);
	const userConfigDir = useLoaderData({ from: "/welcome/" });
	const { mutateAsync } = useMutation({ mutationFn: SetFirstLaunchDone });

	return (
		<div className="relative min-h-screen">
			<div className="circuit-pattern-bg absolute h-dvh w-dvw" />
			<div className="relative z-10 flex min-h-screen items-center justify-center">
				<div className="transition-all duration-500">
					<TapaTextLogo size="m" color="pink-accent" />
				</div>

				<div className="mr-7 ml-3 h-[400px] w-[5px] rounded-2xl bg-[#d9d9d9]"></div>

				<div className="text-white">
					<h1 className="font-lobster my-5 text-5xl text-white">Welcome to TAPA!</h1>
					<p className="my-5">Thank you for choosing Tapa! Please Note that Tapa:</p>

					<ul className="list-inside list-disc pl-4 leading-8">
						<li>
							is <strong>minimalist</strong> by design 🧹
						</li>

						<li>
							is <strong>local-first</strong> and does not gather <strong>any</strong> data 📡
						</li>

						<li>
							is <strong>keyboard-centric</strong> (check the shortcuts) ⌨️
						</li>

						<li>
							stores all your data at <code>{userConfigDir}</code> 📂
						</li>
					</ul>

					<p className="my-4">
						For more information, checkout{" "}
						<a
							href="https://github.com/Amir-Zouerami/TAPA"
							target="_blank"
							rel="noopener noreferrer"
							className="font-roboto-black text-blue-400 transition-all hover:text-blue-300"
							onClick={() => openInUserBrowser(APP_META_DATA.TAPA_GITHUB_URL)}
						>
							<SquareArrowOutUpRight size={18} className="inline" /> Tapa's Github Page
						</a>
					</p>

					<Link
						to="/"
						onClick={async () => {
							setIsLoading(true);
							await mutateAsync();
							setTimeout(() => {
								setIsLoading(false);
								router.navigate({ to: "/" });
							}, 1000);
						}}
						className="bg-pink-accent hover:bg-pink-accent-light float-right mt-5 inline-block animate-bounce rounded-sm px-16 py-3 transition-all"
					>
						{isLoading ? (
							<span>
								Start Developing <LoaderCircle className="inline animate-spin" />
							</span>
						) : (
							"Start Developing 🚀"
						)}
					</Link>
				</div>
			</div>
		</div>
	);
}
