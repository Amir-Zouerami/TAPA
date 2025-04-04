import Battery from "@/components/reusable/Battery";

function InitialLoadingPage() {
	return (
		<div className="flex h-dvh w-dvw items-center justify-center">
			<div>
				<div>Damn, you have a slow computer...!</div>
				<Battery />
			</div>
		</div>
	);
}

export default InitialLoadingPage;
