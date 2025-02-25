function loadServicesTable() {
    $("#servicesTableBody tr").detach();
    // must contact server to get service list as JSON
    $('#servicesTableBody').append('<tr><td>my data</td><td>more data</td></tr>');
}

function addListenerCallback() {
    let nameTxt = $('#listenerAddId').val();
    let typeTxt = $('#listenerType').find(":selected").val();
    if (!nameTxt) {
        alert("ID can't be empty.");
        return
    }
    $.ajax({
	url: "http://localhost:3333/listener/add",
	data: { name: nameTxt, type: typeTxt },
	method: "GET",
	success: function(data) {
	    loadServicesTable();
	    alert("Success. reloading");
//	    location.reload(true); // Or partially reload the table
	},
    });
}
