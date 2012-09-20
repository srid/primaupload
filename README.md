# Getting Started

PrimaUpload is a [Go](http://golang.org/) web application. To get it running, you need bazaar, mercurial (for `go get` to work) and Go itself:

```bash
## mac setup
brew install bzr mercurial go
## ubuntu setup
sudo apt-get install bzr mercurial golang
```

Set up an empty directory as your GOPATH. Dependencies downloaded by the go tool will go into this directory,

```bash
export GOPATH=~/go
mkdir -p $GOPATH
go get
```

To compile and run the server:

```bash
make run
```

Then visit http://0.0.0.0:8080/

If you want to access a publicly deployed version, please contact me.

# Tools used

* **Go** is used on the server as that is my favourite programming language at the moment.
* **[YUI Uploader](http://yuilibrary.com/yui/docs/uploader/)** is used as a cross-browser JavaScript file upload library.

# IE support

Internet Explorer versions upto 9.0 [do not support](http://caniuse.com/xhr2) XMLHttpRequest 2, which adds more functionality to AJAX requests like file uploads, transfer progress information and the ability to send form data. This is the primary reason we use the YUI uploader library (which uses Flash on IE browsers).

The site was tested on Chrome 22.0, Firefox 12.0 and IE 8.0.

# Concurrent requests

Go's http package serves each request in a [goroutine](http://golang.org/src/pkg/net/http/server.go?s=28722:28771#L1042), and since goroutines do not block each other, we automatically support concurrent uploads from multiple users. This was verified manually and also using ApacheBench ([~500 requests/second](https://gist.github.com/3753557)).

# Progressive upload

The upload begins immediately after the user selects a file. While the upload happens, the user may edit the metadata (description in our case) of the being-uploaded file. If he now clicks "save", the form will not be submitted until the upload is completed. This pattern is useful when the file being uploaded is big, and the wait time introduced by the upload can be exploited by the user to fill in the metadata. [SoundCloud](http://soundcloud.com/) has a similar mechanism for its audio upload page.

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

# Scaling

Using S3 for storing uploads will remove local storage as a single point of failure. Further, it allows us to horizontally scale the server instance as the data will not be stored locally. The load balancer should be smart enough to distribute requests based on the upload size (available in the Content-Length header) so as to ensure that no small subset of instances receive a huge of chunk of work even if that is a lesser number of requests with the biggest file uploads.

# Code reading

These are the primary files to begin with if you want to understand the code.

* `web.go` contain the server code.
* `static/upload.js` is the JavaScript code for upload and form management.
* `index.html, view.html` are the HTML templates.
