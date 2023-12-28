function pageInit(){
	$("#pic_file_select").change(function (){
		uploadPic();
		console.log("Selected")
	});
}

function setUploaded(){
	$('#comment').text("Снимок отправлен на сервер");
}

function setError(er){
	$('#comment').text(er);
}

function uploadPic(){
	//setUploaded();
	//return;
	var formData = new FormData();
	formData.append('file', $('#pic_file_select')[0].files[0]);

	$.ajax({
		type: "POST",
		url: "api.php",
		data: formData,
		contentType: false,
		processData: false,
		error: function (error) {
		    // handle error
		    console.log(error.statusText);
		    setError(error.statusText);
		},
		success: function (data) {
			setUploaded();
		}		
	});		
}
