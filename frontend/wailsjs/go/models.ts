export namespace config {
	
	export class Config {
	    base_url: string;
	    timeout: number;
	    theme: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.base_url = source["base_url"];
	        this.timeout = source["timeout"];
	        this.theme = source["theme"];
	    }
	}

}

