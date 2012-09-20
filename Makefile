primaupload:  main.go web.go
		go build -o primaupload

run:    primaupload
		foreman start weblocal

runfresh:	clean fmt run

all:	clean fmt run

# create using: heroku create --buildpack git://github.com/kr/heroku-buildpack-go.git
push:
		git push -f heroku master

all:    primaupload.go

fmt:
		gofmt -w *.go

clean:
		rm -f static/uploads/*

setup:
		brew install bzr golang
		go get

