import { Minus, Maximize2, X } from "lucide-react";
import { WindowToggleMaximise, WindowMinimise, Quit } from "@wails/runtime";

function WindowControls() {
	return (
		<div className="flex items-center justify-center gap-2">
			<div
				className="rounded-2xl p-1 hover:cursor-pointer hover:bg-neutral-600"
				onClick={WindowMinimise}
			>
				<Minus size={18} />
			</div>
			<div
				className="rounded-2xl p-1 hover:cursor-pointer hover:bg-neutral-600"
				onClick={WindowToggleMaximise}
			>
				<Maximize2 size={18} />
			</div>
			<div className="rounded-2xl p-1 hover:cursor-pointer hover:bg-neutral-600" onClick={Quit}>
				<X size={18} />
			</div>
		</div>
	);
}

export default WindowControls;
