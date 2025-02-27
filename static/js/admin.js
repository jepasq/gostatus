function tableData(str) {
    return "<tr>" + str + "</tr>"
}

function elementToTableRow(elem) {
    return "<tr>" + tableData(elem.Name) + tableData(elem.Stype) + "</tr>";
}

function loadServicesTable() {
    $("#servicesTableBody tr").detach();

    $.ajax({
	url: "http://localhost:3333/api/services.json",
	method: "GET",
	success: function(data) {
	    let jsonObj = JSON.parse(data);
	    console.log(typeof jsonObj)
	    let row = $('<tr></tr>');
	    jsonObj.forEach(function(element) {
		console.log(element);
		row.append(tableData(element.Name));
		row.append(tableData(element.Stype));
	    });
	    $('#servicesTableBody').append(row);
	},
	error: function(jqXHR, textStatus, errorThrown) {
	    msg = 'Erreur :' + textStatus + errorThrown;
	    alert(msg);
	}
    });
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
	},
	error: function(jqXHR, textStatus, errorThrown) {
	    msg = 'Erreur :' + textStatus + errorThrown;
	    alert(msg);
	}
    });
}
