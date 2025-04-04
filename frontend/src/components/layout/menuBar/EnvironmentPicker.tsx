import { useState } from "react";
import { cn, findEnvironmentByName } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Check, ChevronsUpDown } from "lucide-react";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import {
	Command,
	CommandEmpty,
	CommandGroup,
	CommandInput,
	CommandItem,
	CommandList,
} from "@/components/ui/command";
import { useConfigStore } from "@/stores/configStore";
import { useDataStore } from "@/stores/dataStore";

interface EnvironmentPickerProps extends React.HTMLAttributes<HTMLDivElement> {}

function EnvironmentPicker({ ...props }: EnvironmentPickerProps) {
	const [open, setOpen] = useState(false);
	const currEnv = useConfigStore(state => state.selectedEnv);
	const environmentsMap = useDataStore(state => state.environments);
	const setSelectedEnv = useConfigStore(state => state.setSelectedEnv);

	return (
		<div {...props}>
			<Popover open={open} onOpenChange={setOpen}>
				<PopoverTrigger asChild>
					<Button
						variant="ghost"
						role="combobox"
						aria-expanded={open}
						className="w-[200px] justify-between"
					>
						{currEnv ? environmentsMap[currEnv].name : "Select environment..."}
						<ChevronsUpDown className="opacity-50" />
					</Button>
				</PopoverTrigger>
				<PopoverContent className="w-[200px] p-0">
					<Command>
						<CommandInput placeholder="Search environments..." className="h-9" />
						<CommandList>
							<CommandEmpty>No environment found</CommandEmpty>
							<CommandGroup>
								{Object.values(environmentsMap)?.map(env => (
									<CommandItem
										key={env.id}
										value={env.name}
										onSelect={selectedEnvName => {
											const newEnv = findEnvironmentByName(
												environmentsMap,
												selectedEnvName,
											);

											setSelectedEnv(newEnv?.id ?? null);
											setOpen(false);
										}}
									>
										{env.name}
										<Check
											className={cn(
												"ml-auto",
												env.id === currEnv ? "opacity-100" : "opacity-0",
											)}
										/>
									</CommandItem>
								))}
							</CommandGroup>
						</CommandList>
					</Command>
				</PopoverContent>
			</Popover>
		</div>
	);
}

export default EnvironmentPicker;
