export namespace models {
	
	export class TaskDTO {
	    ID: number;
	    Title: string;
	    Description: string;
	    Status: string;
	    Priority: number;
	    // Go type: time
	    StartDate: any;
	    // Go type: time
	    DueDate: any;
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
	        this.StartDate = this.convertValues(source["StartDate"], null);
	        this.DueDate = this.convertValues(source["DueDate"], null);
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
	    Priority?: number;
	    SortBy?: string;
	    SortOrder?: number;
	
	    static createFrom(source: any = {}) {
	        return new TaskFilter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Limit = source["Limit"];
	        this.Offset = source["Offset"];
	        this.TagIDs = source["TagIDs"];
	        this.Priority = source["Priority"];
	        this.SortBy = source["SortBy"];
	        this.SortOrder = source["SortOrder"];
	    }
	}

}

