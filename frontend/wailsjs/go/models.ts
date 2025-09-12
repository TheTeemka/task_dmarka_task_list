export namespace models {
	
	export class TagDTO {
	    id: number;
	    name: string;
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new TagDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.color = source["color"];
	    }
	}
	export class TaskDTO {
	    ID: number;
	    Title: string;
	    Description: string;
	    Status: number;
	    Priority: number;
	    // Go type: time
	    DueDate: any;
	    // Go type: time
	    CompletedDate: any;
	    // Go type: time
	    CreatedDate: any;
	
	    static createFrom(source: any = {}) {
	        return new TaskDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Title = source["Title"];
	        this.Description = source["Description"];
	        this.Status = source["Status"];
	        this.Priority = source["Priority"];
	        this.DueDate = this.convertValues(source["DueDate"], null);
	        this.CompletedDate = this.convertValues(source["CompletedDate"], null);
	        this.CreatedDate = this.convertValues(source["CreatedDate"], null);
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
	export class TaskFilter {
	    Limit?: number;
	    Offset?: number;
	    TagIDs: number[];
	    TagsNot: boolean;
	    Status: number[];
	    StatusNot: boolean;
	    Priority: number[];
	    PriorityNot: boolean;
	    Search?: string;
	    SortBy?: string;
	    SortOrder?: string;
	
	    static createFrom(source: any = {}) {
	        return new TaskFilter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Limit = source["Limit"];
	        this.Offset = source["Offset"];
	        this.TagIDs = source["TagIDs"];
	        this.TagsNot = source["TagsNot"];
	        this.Status = source["Status"];
	        this.StatusNot = source["StatusNot"];
	        this.Priority = source["Priority"];
	        this.PriorityNot = source["PriorityNot"];
	        this.Search = source["Search"];
	        this.SortBy = source["SortBy"];
	        this.SortOrder = source["SortOrder"];
	    }
	}

}

