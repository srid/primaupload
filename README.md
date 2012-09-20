# Getting Started

PrimaUpload is a [Go](http://golang.org/) web application. To get it running, you need bazzar, mercurial (got `go get`) and Go itself. On Mac, you would normally run:

```bash
brew install bzr mercurial golang
```

To compile and run the server:

```bash
make run
```

Then visit http://0.0.0.0:8080/

# Tools used

* **Go** is used for the server. The [http](http://golang.org/pkg/net/http/) package from Go serves each request in a separate goroutine.
* **[YUI Uploader](http://yuilibrary.com/yui/docs/uploader/)** provides a cross-browser JavaScript file upload library

# IE7 support

Internet Explorer versions upto 9.0 [do not support](http://caniuse.com/xhr2) XMLHttpRequest 2, which adds more functionality to AJAX requests like file uploads, transfer progress information and the ability to send form data. This is the primary reason we use the YUI uploader library (which uses Flash on IE browsers).

# Concurrent

GO's http package serves each request in a goroutine, and since goroutines do not block each other, we automatically support concurrent uploads from multiple users.

# Progressive upload

The upload begins immediately after the user selects a file. While the upload happens, the user may edit metadata of the uploaded (description in our case). If he now clicks "save", the form will not be submitted until the upload is completed. This pattern is useful when the upload may take a minute more, which idleness can be exploited by the user to enter the metadata. [Soundcloud](http://soundcloud.com/) has a similar mechanism for its audio upload page.

The YUI uploader library supports uploading multiple files as well, though that is not used on this project.

# Deployment

This app can be deployed to Heroku with the help of [Go buildpack](https://gist.github.com/299535bbf56bf3016cba). To do this, run:

```bash
heroku create --buildpack git://github.com/kr/heroku-buildpack-go.git
git push -f heroku master
```

# Upload conflicts

As two users may upload files with same name, each filename is prepended with an [UUID](http://en.wikipedia.org/wiki/Universally_unique_identifier). 

# Persistent uploads

Applications deployed to Heroku get an [ephemeral filesystem](https://devcenter.heroku.com/articles/dynos#ephemeral-filesystem). Therefore, all uploaded files will go away on an app restart or redeployment.

One solution to this problem is to have the sever upload the files to Amazon S3. This would however introduce a subsequent delay right after the browser has completed the upload. 