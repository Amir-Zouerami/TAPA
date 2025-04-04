import {
	AlertDialog,
	AlertDialogAction,
	AlertDialogContent,
	AlertDialogDescription,
	AlertDialogFooter,
	AlertDialogHeader,
	AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { APP_META_DATA } from "@/config/constants";
import { openInUserBrowser } from "@/utils/os";
import { useRouter } from "@tanstack/react-router";
import { SquareArrowOutUpRight } from "lucide-react";

interface ErrorAlertProps {
	open: boolean;
	message: string | React.ReactNode;
	reportable: boolean;
}

function ErrorAlert({ open, message, reportable }: ErrorAlertProps) {
	const router = useRouter();

	return (
		<AlertDialog open={open}>
			<AlertDialogContent>
				<AlertDialogHeader>
					<AlertDialogTitle className="flicker-effect font-roboto-black text-red-400">
						An Error Occurred
					</AlertDialogTitle>
					<AlertDialogDescription className="mt-1 leading-6">
						{message}{" "}
						{reportable ? (
							<span>
								Please report this issue on the{" "}
								<a
									href="https://github.com/Amir-Zouerami/TAPA"
									target="_blank"
									rel="noopener noreferrer"
									className="font-roboto-black text-blue-400 transition-all hover:text-blue-300"
									onClick={() =>
										openInUserBrowser(APP_META_DATA.TAPA_GITHUB_URL + "/issues")
									}
								>
									<SquareArrowOutUpRight size={14} className="inline" /> Issues section of
									our GitHub page.
								</a>
								.
							</span>
						) : null}
					</AlertDialogDescription>
				</AlertDialogHeader>
				<AlertDialogFooter>
					<AlertDialogAction
						onClick={() => {
							router.invalidate();
							router.navigate({ to: "/", replace: true });
						}}
					>
						Reload
					</AlertDialogAction>
				</AlertDialogFooter>
			</AlertDialogContent>
		</AlertDialog>
	);
}

export default ErrorAlert;
