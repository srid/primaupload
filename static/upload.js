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

		// upload progress monitoring
		uploader.on("totaluploadprogress", reportProgress);
		function reportProgress (event) {
			Y.one("#overallProgress").setHTML(
				"Total uploaded: <strong>" + event.percentLoaded + "%</strong");
		}
	}
})