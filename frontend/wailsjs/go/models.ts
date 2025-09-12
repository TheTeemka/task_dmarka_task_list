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
	    id: number;
	    title: string;
	    description: string;
	    status: number;
	    priority: number;
	    due_date: string;
	    completed_date: string;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new TaskDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.status = source["status"];
	        this.priority = source["priority"];
	        this.due_date = source["due_date"];
	        this.completed_date = source["completed_date"];
	        this.created_at = source["created_at"];
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

