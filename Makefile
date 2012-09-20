primaupload:  web.go
		go build -o primaupload

run:	primaupload
	./primaupload

runf:    primaupload
		foreman start weblocal

runfresh:	clean fmt runf

all:	clean fmt run

# create using: heroku create --buildpack git://github.com/kr/heroku-buildpack-go.git
push:
		git push -f heroku master

all:    primaupload.go

fmt:
		gofmt -w *.go

clean:
		rm -f static/uploads/[0-9]*

setup:
		brew install bzr golang
		go get

