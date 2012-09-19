// TODO: setup crossdomain.xml policy file on the server

YUI().use('uploader', function (Y) {
	Y.log("Detected uploader type:" + Y.Uploader.TYPE);
	if (Y.Uploader.TYPE != "none") {
		var uploader = new Y.Uploader({
			width:  "120px",
			height: "40px",
			multipleFiles: false,
		}).render("#myUploader");

		var form = Y.one('form');
		var submitBtn = form.one('input[type=submit]');
		var descArea = form.one('textarea');
		var uploading = false;   // is a file being uploaded?
		var autosubmit = false;  // did the user want to submit the form?

		// clear and initializing state of elements
		function clear(){
			uploading = true;
			autosubmit = false;
			submitBtn.set('value', 'Save');
			submitBtn.set('disabled', false);
			descArea.set('disabled', false);
			Y.one("#savedfile").set('value', '');
			Y.one("#overallProgress").hide();
			Y.one("#serverdata").hide();
		}

		// defer form submit and freeze it no more edits are possible.
		function deferFormSubmit(){
			autosubmit = true;
			submitBtn.set('disabled', true);
			descArea.set('disabled', true);
			submitBtn.set('value', 'Waiting for the upload to finish...')
		}

		// call when the upload is completed
		function uploadComplete(uploadedPath){
			Y.one("#serverdata a").set('href', uploadedPath);
			// add the uploaded file path to the form, that is about to be saved.
			// this allows the server to link the description with the file.
			Y.one("#savedfile").set('value', uploadedPath);
			Y.one("#serverdata").show();
			Y.one("#overallProgress").hide();
			uploading = false;
			if (autosubmit){
				submitBtn.set('value', 'Submitting...')
				form.submit()
			}			
		}

		// once a file is selected, begin uploading
		uploader.after("fileselect", function (event){
			clear();
			uploader.uploadAll();
		})

		uploader.on("uploadstart", function(event) {
			Y.one("#overallProgress").show();
		})

		// upload progress monitoring
		uploader.on("totaluploadprogress", function (event) {
			Y.one("#overallProgress").setHTML(
				"Total uploaded: <strong>" + event.percentLoaded + "%</strong");
		});

		// when upload completes, server sends a uuid for the uploaded file
		// save this for later submit
		uploader.on("uploadcomplete", function (event){
			uploadComplete(event.data);
		})

		// if the user submits already, wait for the upload to be complete
		form.on('submit', function (event){
			if (uploading){
				event.preventDefault();
				deferFormSubmit();
			}
		})
	}
})