window.addEventListener("load", function load(event){

	var load_id = function(id){
		
		var fb = document.getElementById("feedback");
		fb.innerText = "fetching data for WOF ID " + id;
		
		mapzen.whosonfirst.inspector.lookup(id, function(rsp){
			
			var r = mapzen.whosonfirst.inspector.render(rsp);		
			
			var results = document.getElementById("results");
			results.innerHTML = "";
			
			results.appendChild(r);

			fb.innerText = "";			
		});		
	};
	
	var button = document.getElementById("lookup");
	
	button.onclick = function(){

	    try {
    		var id = document.getElementById("wof_id");

		id = id.value;
		id = id.trim();

		if (id == ""){
		    return false;
		}

		id = parseInt(id);

		if (id < 0){
		    return false;
		}

		    load_id(id);
	    }
	    
	    catch (e){
		console.log("WHAT", e);
	    }
	    
    	    return false;
	};

	var hash = location.hash;
	hash = hash.substring(1,);

	console.log("HASH", hash);
	
	if (hash){

		var id = parseInt(hash);
		console.log("ID", id);
		
		if (id > 0){
			load_id(id);
		}		
	}
});
