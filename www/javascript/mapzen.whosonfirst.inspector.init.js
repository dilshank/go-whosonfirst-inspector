window.addEventListener("load", function load(event){

	var load_id = function(id){

		location.hash = "#" + id;
		
		var fb = document.getElementById("feedback");
		fb.innerText = "fetching data for WOF ID " + id;

		var t1 = new Date();
		
		mapzen.whosonfirst.inspector.lookup(id, function(rsp){

			var t2 = new Date();

			var ttl = t2 - t1;
			console.log("TTL", ttl);
			
			var r = mapzen.whosonfirst.inspector.render(rsp);		
			
			var results = document.getElementById("results");
			results.innerHTML = "";
			
			results.appendChild(r);

			fb.innerText = "";			
		});		
	};
	
	var button = document.getElementById("lookup");
	
	button.onclick = function(){

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
	};

	var hash = location.hash;
	hash = hash.substring(1,);
	hash = hash.trim();
	
	if (hash){

		var id = parseInt(hash);
		
		if (id > 0){
			load_id(id);
		}		
	}
});
