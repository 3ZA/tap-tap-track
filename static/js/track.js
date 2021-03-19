// Set up our HTTP request
//
//

function bindStateChanges(){
	var checkboxes = document.querySelectorAll("input[type=checkbox]");
	for (var index = 0; index < checkboxes.length; index++) {
		checkboxes[index].addEventListener(
			"change", 
			function(evt){ 
				var checkbox = evt.target;
				console.log(checkbox.value + " changed to " + checkbox.checked);
				postDoneTask(checkbox.value, checkbox.checked)
			}
		);
	}
}

window.onload = function() {
	bindStateChanges();
}


function postDoneTask(habit, done) {
	var xhr = new XMLHttpRequest();


	// Setup our listener to process compeleted requests
	xhr.onreadystatechange = function () {

		// Only run if the request is complete
		if (xhr.readyState === XMLHttpRequest.DONE) return;

		// Process our return data
		if (xhr.status >= 200 && xhr.status < 300) return;
	};

	xhr.open('POST', 'http://localhost:8585/habits', true);
	xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
	xhr.send('habit='+habit+'&'+'done='+done);
}

// Create and send a GET request
// The first argument is the post type (GET, POST, PUT, DELETE, etc.)
// The second argument is the endpoint URL
