/// Add the standard action button with icons. Added to each line.
function tableActions() {
    return $(`<td><button type="button" class="btn btn-secondary">
		  <i class="bi bi-stop-btn"></i>
		</button>
		<button type="button" class="btn btn-secondary">
		  <i class="bi bi-pause-btn"></i>
		</button>
		<button type="button" class="btn btn-secondary">
		  <i class="bi bi-trash"></i>
		</button>/Resume/Delete</td>`);
}

function tableData(str) {
    return $('<td></td>').text(str);
}

function elementToTableRow(elem) {
    return "<tr>" + tableData(elem.Name) + tableData(elem.Stype) + "</tr>";
}

function loadServicesTable() {
    $("#servicesTableBody tr").detach();
    $('#loadingRow').show();

    $.ajax({
	url: "http://localhost:3333/api/services.json",
	method: "GET",
	success: function(data) {
	    let jsonObj = JSON.parse(data);
	    console.log(typeof jsonObj)
	    jsonObj.forEach(function(element) {
		let row = $('<tr></tr>');
		row.append(tableData(element.Name));
		row.append(tableData(element.Stype));
		row.append(tableData("")); /* fake Params */
		row.append(tableData("")); /* fake Status */
		row.append(tableActions());
		$('#servicesTableBody').append(row);
	    });
	},
	error: function(jqXHR, textStatus, errorThrown) {
	    msg = 'Erreur :' + textStatus + errorThrown;
	    alert(msg);
	}
    });
    $('#loadingRow').hide();
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
