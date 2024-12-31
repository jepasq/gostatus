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
	    // $("#response").html(data);
	    alert("Success");
	},
    });
}
