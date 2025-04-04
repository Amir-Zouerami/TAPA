import type { models } from "@wails/go/models";
import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

export function findEnvironmentByName(obj: Record<number, models.Environment>, targetName: string) {
	for (const key in obj) {
		if (Object.prototype.hasOwnProperty.call(obj, key) && obj[key].name === targetName) {
			return obj[key];
		}
	}
}
