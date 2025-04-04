import type { MouseEvent } from "react";
import MenuBarItems from "./MenuBarItems";
import WindowControls from "./WindowControls";
import { Menubar } from "@/components/ui/menubar";
import EnvironmentPicker from "./EnvironmentPicker";
import TapaTextLogo from "@/components/branding/TapaTextLogo";
import { WindowToggleMaximise } from "@wails/runtime/runtime";

const handleWindowToggleMaximize = (e: MouseEvent<HTMLDivElement, globalThis.MouseEvent>) => {
	if (e.target === e.currentTarget) {
		WindowToggleMaximise();
	}
};

function ApplicationMenuBar() {
	return (
		<div className="draggable w-full">
			<Menubar className="flex items-center justify-between" onDoubleClick={handleWindowToggleMaximize}>
				<div className="flex items-center">
					<TapaTextLogo size="xxs" color="orange" className="my-1 mr-5" />
					<MenuBarItems />
				</div>
				<div className="flex">
					<EnvironmentPicker className="mr-10" />
					<WindowControls />
				</div>
			</Menubar>
		</div>
	);
}

export default ApplicationMenuBar;
