export namespace models {
	
	export class Tab {
	    request_id: number;
	    is_saved: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Tab(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.request_id = source["request_id"];
	        this.is_saved = source["is_saved"];
	    }
	}
	export class AppState {
	    id: number;
	    selected_environment?: number;
	    selected_tab?: number;
	    first_launch: boolean;
	    open_tabs: Tab[];
	    // Go type: types
	    updated_at: any;
	
	    static createFrom(source: any = {}) {
	        return new AppState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.selected_environment = source["selected_environment"];
	        this.selected_tab = source["selected_tab"];
	        this.first_launch = source["first_launch"];
	        this.open_tabs = this.convertValues(source["open_tabs"], Tab);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Collection {
	    id: number;
	    name: string;
	    description?: string;
	    position: number;
	    // Go type: types
	    updated_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Collection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.position = source["position"];
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RequestSummary {
	    id: number;
	    collection_id?: number;
	    folder_id?: number;
	    name: string;
	    method: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestSummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.collection_id = source["collection_id"];
	        this.folder_id = source["folder_id"];
	        this.name = source["name"];
	        this.method = source["method"];
	    }
	}
	export class Folder {
	    id: number;
	    collection_id: number;
	    name: string;
	    position: number;
	    // Go type: types
	    updated_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Folder(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.collection_id = source["collection_id"];
	        this.name = source["name"];
	        this.position = source["position"];
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class CollectionsTree {
	    collections: Record<number, Collection>;
	    folders: Record<number, Folder>;
	    request_list: Record<number, RequestSummary>;
	
	    static createFrom(source: any = {}) {
	        return new CollectionsTree(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.collections = this.convertValues(source["collections"], Collection, true);
	        this.folders = this.convertValues(source["folders"], Folder, true);
	        this.request_list = this.convertValues(source["request_list"], RequestSummary, true);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class UserSettings {
	    id: number;
	    theme: string;
	    max_history: number;
	    font_size: number;
	
	    static createFrom(source: any = {}) {
	        return new UserSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.theme = source["theme"];
	        this.max_history = source["max_history"];
	        this.font_size = source["font_size"];
	    }
	}
	export class KeyboardShortcut {
	    id: number;
	    action: string;
	    shortcut: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyboardShortcut(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.action = source["action"];
	        this.shortcut = source["shortcut"];
	    }
	}
	export class Environment {
	    id: number;
	    name: string;
	    // Go type: types
	    updated_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Environment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DashboardData {
	    environments: Record<number, Environment>;
	    current_env_variables: Record<string, string>;
	    keyboard_shortcuts: Record<number, KeyboardShortcut>;
	    user_settings: UserSettings;
	    app_state: AppState;
	
	    static createFrom(source: any = {}) {
	        return new DashboardData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.environments = this.convertValues(source["environments"], Environment, true);
	        this.current_env_variables = source["current_env_variables"];
	        this.keyboard_shortcuts = this.convertValues(source["keyboard_shortcuts"], KeyboardShortcut, true);
	        this.user_settings = this.convertValues(source["user_settings"], UserSettings);
	        this.app_state = this.convertValues(source["app_state"], AppState);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	
	

}

export namespace options {
	
	export class SecondInstanceData {
	    Args: string[];
	    WorkingDirectory: string;
	
	    static createFrom(source: any = {}) {
	        return new SecondInstanceData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Args = source["Args"];
	        this.WorkingDirectory = source["WorkingDirectory"];
	    }
	}

}

