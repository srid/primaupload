// TODO: setup crossdomain.xml policy file on the server
// TODO: setup header Access-Control-Allow-Origin

YUI().use('uploader', function (Y) {
	Y.log("Detected uploader type:" + Y.Uploader.TYPE);
	if (Y.Uploader.TYPE != "none") {
		var uploader = new Y.Uploader({
			width:  "120px",
			height: "40px",
			multipleFiles: false,
		}).render("#myUploader");

		// once a file is selected, begin uploading
		uploader.after("fileselect", function (event){
			uploader.uploadAll();
		})

		uploader.on("uploadstart", function(event) {
			Y.one("#overallProgress").show();
			Y.one("#serverdata").hide();
		})

		// upload progress monitoring
		uploader.on("totaluploadprogress", function (event) {
			Y.one("#overallProgress").setHTML(
				"Total uploaded: <strong>" + event.percentLoaded + "%</strong");
		});

		// when upload completes, server sends a uuid for the uploaded file
		// save this for later submit
		uploader.on("uploadcomplete", function (event){
			Y.one("#serverdata a").set('href', event.data);
			Y.one("#serverdata").show();
			Y.one("#overallProgress").hide();
		})
	}
})